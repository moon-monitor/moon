package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/palace"
)

type CallbackService struct {
	pb.UnimplementedCallbackServer
}

func NewCallbackService() *CallbackService {
	return &CallbackService{}
}

func (s *CallbackService) SendMsgCallback(ctx context.Context, req *pb.SendMsgCallbackRequest) (*pb.SendMsgCallbackReply, error) {
	return &pb.SendMsgCallbackReply{}, nil
}
