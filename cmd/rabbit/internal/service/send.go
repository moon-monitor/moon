package service

import (
	"context"

	common "github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SendService struct {
	apiv1.UnimplementedSendServer
}

func NewSendService() *SendService {
	return &SendService{}
}

func (s *SendService) Email(ctx context.Context, req *apiv1.SendEmailRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SendService) Sms(ctx context.Context, req *apiv1.SendSmsRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SendService) Hook(ctx context.Context, req *apiv1.SendHookRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
