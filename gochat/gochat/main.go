package main

import (
	pb "gochat/proto"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterGoChatServiceService(s.Service("gochat.GoChatService"), &goChatServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
