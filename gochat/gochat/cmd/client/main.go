// Package main is originally generated by trpc-cmdline v1.0.5.
// It is located at `project/cmd/client`.
// Run this file by executing `go run cmd/client/main.go` in the project directory.
package main

import (
	"fmt"
	pb "gochat/proto"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func callGoChatServiceGoChat(content string) string {
	proxy := pb.NewGoChatServiceClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()

	req := &pb.GoChatRequest{}
	req.Msg = content
	reply, err := proxy.GoChat(ctx, req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	// log.Debugf("simple  rpc   receive: %+v", reply)
	return reply.Msg
}

func callGoChatServiceGoChatAudio(content string) {
	proxy := pb.NewGoChatServiceClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()

	req := &pb.GoChataudioRequest{}
	req.Msg = content
	reply, err := proxy.GoChatAudio(ctx, req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func main() {
	// Load configuration following the logic in trpc.NewServer.
	cfg, err := trpc.LoadConfig(trpc.ServerConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	trpc.SetGlobalConfig(cfg)
	if err := trpc.Setup(cfg); err != nil {
		panic("setup plugin fail: " + err.Error())
	}
	content := "你好，请简要回复,100字以内：我想准备好了之后面试研发岗位,怎么刷算法题，感觉好难"
	reqcontent := callGoChatServiceGoChat(content)
	fmt.Printf("reqcontent:%s\n", reqcontent)
	callGoChatServiceGoChatAudio(reqcontent)
}
