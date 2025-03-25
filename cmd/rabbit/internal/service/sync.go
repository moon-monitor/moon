package service

import (
	"context"

	common "github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SyncService struct {
	apiv1.UnimplementedSyncServer
}

func NewSyncService() *SyncService {
	return &SyncService{}
}
func (s *SyncService) Sms(ctx context.Context, req *apiv1.SyncSmsRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SyncService) Email(ctx context.Context, req *apiv1.SyncEmailRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
