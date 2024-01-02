package main

import (
	"context"
	"reflect"
	"testing"

	pb "gochat/proto"

	"github.com/golang/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/gochat/proto/gochat_mock.go -package=proto -self_package=gochat/proto --source=stub/gochat/proto/gochat.trpc.go

func Test_goChatServiceImpl_GoChat(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	goChatServiceService := pb.NewMockGoChatServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := goChatServiceService.EXPECT().GoChat(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.GoChatRequest) (*pb.GoChatResponse, error) {
		s := &goChatServiceImpl{}
		return s.GoChat(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.GoChatRequest
		rsp *pb.GoChatResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.GoChatResponse
			var err error
			if rsp, err = goChatServiceService.GoChat(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("goChatServiceImpl.GoChat() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("goChatServiceImpl.GoChat() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_goChatServiceImpl_GoChatAudio(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	goChatServiceService := pb.NewMockGoChatServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := goChatServiceService.EXPECT().GoChatAudio(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.GoChataudioRequest) (*pb.GoChataudioResponse, error) {
		s := &goChatServiceImpl{}
		return s.GoChatAudio(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.GoChataudioRequest
		rsp *pb.GoChataudioResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.GoChataudioResponse
			var err error
			if rsp, err = goChatServiceService.GoChatAudio(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("goChatServiceImpl.GoChatAudio() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("goChatServiceImpl.GoChatAudio() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
