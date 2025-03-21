package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SyncService struct {
	pb.UnimplementedSyncServer
}

func NewSyncService() *SyncService {
	return &SyncService{}
}

func (s *SyncService) Templates(ctx context.Context, req *pb.SyncTemplatesRequest) (*pb.SyncTemplatesReply, error) {
	return &pb.SyncTemplatesReply{}, nil
}
func (s *SyncService) Hooks(ctx context.Context, req *pb.SyncHooksRequest) (*pb.SyncHooksReply, error) {
	return &pb.SyncHooksReply{}, nil
}
func (s *SyncService) Receivers(ctx context.Context, req *pb.SyncReceiversRequest) (*pb.SyncReceiversReply, error) {
	return &pb.SyncReceiversReply{}, nil
}
func (s *SyncService) Sms(ctx context.Context, req *pb.SyncSmsRequest) (*pb.SyncSmsReply, error) {
	return &pb.SyncSmsReply{}, nil
}
