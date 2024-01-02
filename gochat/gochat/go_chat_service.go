package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	pb "gochat/proto"

	"github.com/hajimehoshi/go-mp3"
	"trpc.app.GoChatService/chat"
	"trpc.group/trpc-go/tnet/log"
)

type goChatServiceImpl struct {
	pb.UnimplementedGoChatService
}

// GoChat GoChat says hello.
func (s *goChatServiceImpl) GoChat(
	ctx context.Context,
	req *pb.GoChatRequest,
) (*pb.GoChatResponse, error) {
	log.Infof("got Chat request: %s", req.Msg)
	// 获取chatgpt返回的消息
	chattext := &chat.ChatCompletionsRequest{}
	chattext.Model = "gpt-3.5-turbo"
	chattext.Temperature = 0.7
	chattext.Messages = append(chattext.Messages, chat.Message{
		Content: req.Msg,
		Role:    "user",
	})
	rspmsg, err := chat.Chat(chattext)
	if err != nil {
		panic(err)
	}

	rsp := &pb.GoChatResponse{}
	rsp.Msg = rspmsg.Choices[0].Message.Content
	return rsp, nil
}

func (s *goChatServiceImpl) GoChatAudio(
	ctx context.Context,
	req *pb.GoChataudioRequest,
) (*pb.GoChataudioResponse, error) {
	log.Infof("got Chataudio request: %s", req.Msg)
	// 获取chatgpt返回的消息
	chataudio := &chat.AudioSpeechRequest{}
	chataudio.Model = "tts-1-hd"
	chataudio.Input = req.Msg
	chataudio.Voice = "onyx"
	chataudio.OutputFile = "output.mp3"

	filename := "data/voice/" + req.Filename
	if filename == "" {
		fmt.Println("filename is nil!!!")
		return &pb.GoChataudioResponse{Msg: "filename is nil!!!"}, nil
	}

	rspmsg := ""
	p := 0
	p2 := 0
retry:
	if _, err := os.Stat(filename); os.IsNotExist(err) || strings.Contains(filename, "answer") == true {
		log.Info("文件不存在，开始生成...")
		rspmsg, err = chat.Audio(chataudio, filename)
		if err != nil {
			if p <= 5 {
				log.Infof("文生语音失败，开始第%v次重试", p)
				p++
				goto retry
			}
			panic(err)
		}
	}

	// 循环等待文件生成完毕，每隔一段时间检查文件是否已经存在
	maxWait := 300 // 最大等待时间，单位：秒
	waitTime := 1  // 每次等待的时间间隔，单位：秒
	waited := 0
	for {
		_, err := os.Stat(filename)
		if err == nil {
			// 文件已存在，说明生成完成，可以进行后续操作
			break
		}
		if waited >= maxWait {
			log.Info("等待超时，文件未生成完毕")
			return &pb.GoChataudioResponse{Msg: "等待超时，文件未生成完毕"}, nil
		}
		// 继续等待一段时间
		time.Sleep(time.Duration(waitTime) * time.Second)
		waited += waitTime
	}

	if p2 > 2 {
		time.Sleep(30 * time.Second)
	}
	// 计算生成的 MP3 文件的时长
	duration, err := calculateMP3Duration(filename)
	if err != nil {
		if p2 <= 50 {
			log.Infof("计算生成的 MP3 文件的时长，开始第%v次重试，删除文件重新生成", p2)
			err = os.Remove(filename)
			if err != nil {
				log.Infof("删除文件%v失败%v", filename, err)
			}
			p2++
			time.Sleep(30 * time.Second)
			goto retry
		}
		panic(err)
	}
	// 转换下
	duration = duration / 105.12 * 26

	// go func() {
	// 	err := playMP3InBackground(filename)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
	log.Infof("开始播放文件%v", filename)
	err = playMP3InBackground(filename)
	if err != nil {
		panic(err)
	}

	rsp := &pb.GoChataudioResponse{}
	rsp.Msg = rspmsg
	rsp.Filelongth = int32(duration)
	log.Infof("Filename=%vFIlelongth=%v", filename, duration)
	return rsp, nil
}

func playMP3InBackground(filename string) error {
	cmd := exec.Command("cmd", "/C", "start", filename)
	return cmd.Run()
}

// func calculateMP3Duration(filename string) (float64, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer file.Close()

// 	decoder, err := mp3.NewDecoder(file)
// 	if err != nil {
// 		return 0, err
// 	}

// 	totalSamples := int64(decoder.Length())
// 	sampleRate := int64(decoder.SampleRate())

//		duration := time.Duration(float64(totalSamples) / float64(sampleRate) * float64(time.Second))
//		return duration.Seconds(), nil
//	}
func calculateMP3Duration(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return 0, err
	}

	// 获取采样率和采样总数
	sampleRate := float64(decoder.SampleRate())
	totalSamples := float64(decoder.Length())

	// 计算持续时间
	durationInSeconds := totalSamples / sampleRate

	// 使用文件大小和比特率估算持续时间
	fileSize := float64(fileInfo.Size())
	bitRate := (fileSize * 8) / durationInSeconds

	// 根据比特率修正持续时间
	duration := fileSize / (bitRate / 8)

	return duration, nil
}
