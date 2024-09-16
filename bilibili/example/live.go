/*
提供了与直播平台通信的功能，例如启动直播应用、发送心跳、结束直播应用等。它还可能提供了与直播消息交互的能力，例如处理不同类型的消息，解析消息的命令和数据等。
rCfg变量初始化了一个live.Config对象，其中包含了与直播平台通信所需的配置信息，例如应用ID、身份码等。
然后，通过live.NewClient(rCfg)创建了一个live.Client对象，用于与直播平台建立连接并进行通信
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vtb-link/bianka/live"
	"github.com/vtb-link/bianka/proto"
	"github.com/vtb-link/bianka/savetofile"
)

var rCfg = live.NewConfig(
	"",
	"",
	0, // 应用id
)

var code = "" // 身份码 也叫 idCode

func main() {
	// 创建sdk实例
	liveClient := live.NewClient(rCfg)

	startResp, err := liveClient.AppStart(code)
	if err != nil {
		panic(err)
	}

	// 启用项目心跳 20s一次
	tk := time.NewTicker(time.Second * 20)
	go func() {
		for {
			select {
			case <-tk.C:
				// 心跳
				if err := liveClient.AppHeartbeat(startResp.GameInfo.GameID); err != nil {
					log.Println("Heartbeat fail", err)
				}
			}
		}
	}()

	// app end
	defer func() {
		tk.Stop()
		liveClient.AppEnd(startResp.GameInfo.GameID)
	}()

	// 一键开启websocket
	wcs, err := liveClient.StartWebsocket(startResp, map[uint32]live.DispatcherHandle{
		proto.OperationMessage: messageHandle,
	}, func(startResp *live.AppStartResponse) {
		// 注册关闭回调
		log.Println("WebsocketClient onClose", startResp)
	})

	if err != nil {
		panic(err)
	}

	defer wcs.Close()

	// 退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("WebsocketClient exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func messageHandle(msg *proto.Message) error {
	fmt.Println("messageHandlestart")
	// 单条消息raw
	log.Println(string(msg.Payload()))

	// 自动解析 sdk提供了自动解析消息的方法，可以快速解析为对应的cmd和data
	// 具体的cmd 可以参考 live/cmd.go
	cmd, data, err := live.AutomaticParsingMessageCommand(msg.Payload())
	if err != nil {
		return err
	}

	// Switch cmd 可以使用cmd进行switch
	switch cmd {
	case live.CmdLiveOpenPlatformDanmu:
		danmuData, ok := data.(*live.CmdLiveOpenPlatformDanmuData)
		if !ok {
			log.Println("Invalid danmu data")
			return nil
		}
		// 执行弹幕消息处理逻辑，使用danmuData中的数据
		log.Printf("cmd=%v, danmuData=%+v\n", cmd, danmuData)
		savetofile.SaveDanmuData(danmuData)

	case live.CmdLiveOpenPlatformSendGift:
		giftData, ok := data.(*live.CmdLiveOpenPlatformSendGiftData)
		if !ok {
			log.Println("Invalid gift data")
			return nil
		}
		// 执行礼物消息处理逻辑，使用giftData中的数据
		log.Printf("cmd=%v, giftData=%+v\n", cmd, giftData)
		savetofile.SaveGiftData(giftData)

	case live.CmdLiveOpenPlatformSuperChat:
		superChatData, ok := data.(*live.CmdLiveOpenPlatformSuperChatData)
		if !ok {
			log.Println("Invalid super chat data")
			return nil
		}
		// 执行上下线处理逻辑，使用superChatData中的数据
		log.Printf("cmd=%v, superChatData=%+v\n", cmd, superChatData)
		savetofile.SaveSuperChatData(superChatData)

	case live.CmdLiveOpenPlatformSuperChatDel:
		superChatDelData, ok := data.(*live.CmdLiveOpenPlatformSuperChatDelData)
		if !ok {
			log.Println("Invalid super chat del data")
			return nil
		}
		// 执行删除上下线消息处理逻辑，使用superChatDelData中的数据
		log.Printf("cmd=%v, superChatDelData=%+v\n", cmd, superChatDelData)
		savetofile.SaveSuperChatDelData(superChatDelData)

	case live.CmdLiveOpenPlatformGuard:
		guardData, ok := data.(*live.CmdLiveOpenPlatformGuardData)
		if !ok {
			log.Println("Invalid guard data")
			return nil
		}
		// 执行付费大航海消息处理逻辑，使用guardData中的数据
		log.Printf("cmd=%v, guardData=%+v\n", cmd, guardData)
		savetofile.SaveGuardData(guardData)

	case live.CmdLiveOpenPlatformLike:
		likeData, ok := data.(*live.CmdLiveOpenPlatformLikeData)
		if !ok {
			log.Println("Invalid like data")
			return nil
		}
		// 执行点赞消息处理逻辑，使用likeData中的数据
		log.Printf("cmd=%v, likeData=%+v\n", cmd, likeData)
		savetofile.SaveLikeData(likeData)

	default:
		log.Printf("Unknown command: %s\n", cmd)
		return nil
	}

	// Switch data type 可以使用data进行switch
	switch v := data.(type) {
	case *live.CmdLiveOpenPlatformGuardData:
		log.Println(cmd, v)
	}

	return nil
}
