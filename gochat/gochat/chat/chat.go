package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"trpc.app.GoChatService/global"
)

type ChatCompletionsRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionsResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     string `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint any `json:"system_fingerprint"`
}

type AudioSpeechRequest struct {
	Model      string `json:"model"`
	Input      string `json:"input"`
	Voice      string `json:"voice,omitempty"`
	OutputFile string `json:"output_file,omitempty"`
}

func Chat(chatCompletions *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {

	chatCompletionsBytes, err := json.Marshal(chatCompletions)
	if err != nil {
		fmt.Println("Error marshaling chat completions request:", err)
		return nil, err
	}

	proxyURL := global.GetproxyURL()
	chatCompletionsResp, err := sendRequest("https://api.openai.com/v1/chat/completions", chatCompletionsBytes, proxyURL)
	if err != nil {
		fmt.Println("Error sending chat completions request:", err)
		return nil, err
	}
	// fmt.Println("Chat completions response:", string(chatCompletionsResp))

	response := &ChatCompletionsResponse{}
	if err := json.Unmarshal(chatCompletionsResp, &response); err != nil {
		fmt.Println("Error unmarshaling chat completions response:", err)
		return nil, err
	}
	return response, nil

}

func Audio(audioSpeech *AudioSpeechRequest, filename string) (string, error) {
	// audioSpeech := AudioSpeechRequest{
	// 	Model:      model,
	// 	Input:      "你好，我是我",
	// 	Voice:      "onyx",
	// 	OutputFile: "example.mp3",
	// }
	audioSpeechBytes, err := json.Marshal(audioSpeech)
	if err != nil {
		fmt.Println("Error marshaling audio speech request:", err)
		return "", err
	}

	fmt.Println("audio test")
	proxyURL := global.GetproxyURL()
	audioSpeechResponse, err := sendRequest("https://api.openai.com/v1/audio/speech", audioSpeechBytes, proxyURL)
	if err != nil {
		fmt.Println("Error sending audio speech request:", err)
		return "", err
	}
	// 保存响应内容到文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", err
	}
	defer file.Close()

	_, err = file.Write(audioSpeechResponse)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return "", err
	}

	// fmt.Println("Audio speech response:", string(audioSpeechResponse))
	return "", nil
}

func sendRequest(urlmsg string, data []byte, proxyURL string) ([]byte, error) {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 如果代理地址不为空，则设置代理
	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		client.Transport = transport
	}

	// 创建请求
	req, err := http.NewRequest("POST", urlmsg, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	openaikey := global.GetOpenaikey()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaikey)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
