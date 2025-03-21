package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SendService struct {
	pb.UnimplementedSendServer
}

func NewSendService() *SendService {
	return &SendService{}
}

func (s *SendService) Email(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	return &pb.SendEmailReply{}, nil
}
func (s *SendService) Sms(ctx context.Context, req *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	return &pb.SendSmsReply{}, nil
}
func (s *SendService) Hook(ctx context.Context, req *pb.SendHookRequest) (*pb.SendHookReply, error) {
	return &pb.SendHookReply{}, nil
}
func (s *SendService) SendAll(ctx context.Context, req *pb.SendAllRequest) (*pb.SendAllReply, error) {
	return &pb.SendAllReply{}, nil
}
