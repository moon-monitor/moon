package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	"github.com/moon-monitor/moon/pkg/api/palace"
)

func NewCallbackService(logsBiz *biz.Logs) *CallbackService {
	return &CallbackService{
		logsBiz: logsBiz,
	}
}

type CallbackService struct {
	palace.UnimplementedCallbackServer
	logsBiz *biz.Logs
}

func (s *CallbackService) SendMsgCallback(ctx context.Context, req *palace.SendMsgCallbackRequest) (*palace.SendMsgCallbackReply, error) {
	params := build.ToUpdateSendMessageLogStatusParams(req)
	if err := s.logsBiz.UpdateSendMessageLogStatus(ctx, params); err != nil {
		return nil, err
	}
	return &palace.SendMsgCallbackReply{
		Code: 0,
		Msg:  "success",
	}, nil
}

func (s *CallbackService) SyncMetadata(ctx context.Context, req *palace.SyncMetadataRequest) (*palace.SyncMetadataReply, error) {
	return &palace.SyncMetadataReply{
		Code: 0,
		Msg:  "success",
	}, nil
}
