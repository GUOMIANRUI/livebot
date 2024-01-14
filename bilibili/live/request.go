/*
 * MIT License
 *
 * Copyright (c) 2023 VTB-LINK and runstp.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE, AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
 * OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF,
 * OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package live

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	HostProdLiveOpen = "https://live-open.biliapi.com" // 开放平台 (线上环境)
)

type Config struct {
	AccessKey            string // access_key
	AccessKeySecret      string // access_key_secret
	OpenPlatformHttpHost string // 开放平台 (线上环境)
	AppID                int64  // 应用id
}

func NewConfig(accessKey, accessKeySecret string, appID int64) *Config {
	return &Config{
		AccessKey:            accessKey,
		AccessKeySecret:      accessKeySecret,
		OpenPlatformHttpHost: HostProdLiveOpen,
		AppID:                appID,
	}
}

type Client struct {
	rCfg *Config
}

func NewClient(rCfg *Config) *Client {
	return &Client{
		rCfg: rCfg,
	}
}

// AppStart 启动app
func (c *Client) AppStart(code string) (*AppStartResponse, error) {
	startAppReq := AppStartRequest{
		Code:  code,
		AppID: c.rCfg.AppID,
	}

	reqJson, err := json.Marshal(startAppReq)
	if err != nil {
		err = errors.Wrap(err, "json marshal fail")
	}

	resp, err := c.doRequest(string(reqJson), "/v2/app/start")
	if err != nil {
		return nil, errors.WithMessage(err, "start app fail")
	}

	startAppRespData := &AppStartResponse{}
	if err = json.Unmarshal(resp.Data, &startAppRespData); err != nil {
		return nil, errors.Wrapf(err, "json unmarshal fail, data:%s", resp.Data)
	}

	return startAppRespData, nil
}

// AppEnd 关闭app
func (c *Client) AppEnd(gameID string) error {
	endAppReq := AppEndRequest{
		GameID: gameID,
		AppID:  c.rCfg.AppID,
	}

	reqJson, err := json.Marshal(endAppReq)
	if err != nil {
		err = errors.Wrap(err, "json marshal fail")
	}

	_, err = c.doRequest(string(reqJson), "/v2/app/end")
	if err != nil {
		return errors.WithMessage(err, "end app fail")
	}

	return nil
}

// AppHeartbeat 心跳
func (c *Client) AppHeartbeat(gameID string) error {
	heartbeatReq := AppHeartbeatRequest{
		GameID: gameID,
	}

	reqJson, err := json.Marshal(heartbeatReq)
	if err != nil {
		err = errors.Wrap(err, "json marshal fail")
	}

	_, err = c.doRequest(string(reqJson), "/v2/app/heartbeat")
	if err != nil {
		return errors.WithMessage(err, "heartbeat fail")
	}

	return nil
}

// AppBatchHeartbeat 批量心跳
func (c *Client) AppBatchHeartbeat(gameIDs []string) (*AppBatchHeartbeatResponse, error) {
	heartbeatReq := AppBatchHeartbeatRequest{
		GameIDs: gameIDs,
	}

	reqJson, err := json.Marshal(heartbeatReq)
	if err != nil {
		err = errors.Wrap(err, "json marshal fail")
	}

	resp, err := c.doRequest(string(reqJson), "/v2/app/batchHeartbeat")
	if err != nil {
		return nil, errors.WithMessage(err, "heartbeat fail")
	}

	heartbeatResp := &AppBatchHeartbeatResponse{}
	if err = json.Unmarshal(resp.Data, &heartbeatResp); err != nil {
		return nil, errors.Wrapf(err, "json unmarshal fail, data:%s", resp.Data)
	}

	return heartbeatResp, nil
}

// StartWebsocket 启动websocket
// 此方法会一键完成鉴权，心跳，消息分发
func (c *Client) StartWebsocket(startResp *AppStartResponse, dispatcherHandleMap map[uint32]DispatcherHandle, onCloseFunc func(startResp *AppStartResponse)) (*WsClient, error) {
	wc := NewWsClient(
		startResp,
		dispatcherHandleMap,
		nil).
		WithOnClose(onCloseFunc)

	if err := wc.Dial(startResp.WebsocketInfo.WssLink); err != nil {
		return nil, err
	}

	if err := wc.SendAuth(); err != nil {
		return nil, err
	}

	wc.Run()
	return wc, nil
}

func (c *Client) doRequest(reqJson, reqPath string) (*BaseResp, error) {
	return c.DoRequest(reqJson, reqPath, strconv.FormatInt(time.Now().UnixNano(), 10))
}

// DoRequest 发起请求
// 用于用户自定义请求
func (c *Client) DoRequest(reqJson, reqPath, nonce string) (*BaseResp, error) {
	header := &CommonHeader{
		ContentType:       JsonType,
		ContentAcceptType: JsonType,
		Timestamp:         strconv.FormatInt(time.Now().Unix(), 10),
		SignatureMethod:   HmacSha256,
		SignatureVersion:  BiliVersion,
		Nonce:             nonce, // 用于幂等
		AccessKeyID:       c.rCfg.AccessKey,
		ContentMD5:        Md5(reqJson),
	}
	header.Authorization = header.CreateSignature(c.rCfg.AccessKeySecret)

	result := BaseResp{}
	resp, err := resty.New().R().
		SetHeaders(header.ToMap()).
		SetBody(reqJson).
		SetResult(&result).
		Post(c.rCfg.OpenPlatformHttpHost + reqPath)

	if err != nil {
		return nil, errors.Wrapf(err, "request fail, url:%s body: %s", reqPath, reqJson)
	}

	if resp.StatusCode() >= http.StatusBadRequest {
		return nil, errors.Wrapf(BilibiliRequestFailed, "request response not ok, url:%s req: %v code:%d", reqPath, reqJson, resp.StatusCode())
	}

	if !result.Success() {
		return &result, errors.Wrapf(BilibiliResponseNotSuccess, "bilbil response code not ok, url:%s  body: %s result: %v", reqPath, reqJson, result)
	}

	return &result, nil
}

// VerifyH5RequestSignature 验证h5请求签名
func (c *Client) VerifyH5RequestSignature(req *http.Request) bool {
	h5sp := ParseH5SignatureParamsWithRequest(req)

	return c.VerifyH5RequestSignatureWithParams(h5sp)
}

// VerifyH5RequestSignatureWithParams 验证h5请求签名
func (c *Client) VerifyH5RequestSignatureWithParams(h5sp *H5SignatureParams) bool {
	return h5sp.ValidateSignature(c.rCfg.AccessKeySecret)
}
