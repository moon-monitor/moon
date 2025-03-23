package service

import (
	"context"

	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SendService struct {
	apiv1.UnimplementedSendServer
}

func NewSendService() *SendService {
	return &SendService{}
}

func (s *SendService) Email(ctx context.Context, req *apiv1.SendEmailRequest) (*apiv1.SendEmailReply, error) {
	return &apiv1.SendEmailReply{}, nil
}

func (s *SendService) Sms(ctx context.Context, req *apiv1.SendSmsRequest) (*apiv1.SendSmsReply, error) {
	return &apiv1.SendSmsReply{}, nil
}

func (s *SendService) Hook(ctx context.Context, req *apiv1.SendHookRequest) (*apiv1.SendHookReply, error) {
	return &apiv1.SendHookReply{}, nil
}
