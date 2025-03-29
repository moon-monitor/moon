package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/houyi/v1"
)

type SyncService struct {
	pb.UnimplementedSyncServer
}

func NewSyncService() *SyncService {
	return &SyncService{}
}

func (s *SyncService) MetricStrategy(ctx context.Context, req *pb.MetricStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) CertificateStrategy(ctx context.Context, req *pb.CertificateStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) ServerPortStrategy(ctx context.Context, req *pb.ServerPortStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) HttpStrategy(ctx context.Context, req *pb.HttpStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) PingStrategy(ctx context.Context, req *pb.PingStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) EventStrategy(ctx context.Context, req *pb.EventStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) LogsStrategy(ctx context.Context, req *pb.LogsStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
func (s *SyncService) RemoveStrategy(ctx context.Context, req *pb.RemoveStrategyRequest) (*pb.SyncStrategyReply, error) {
	return &pb.SyncStrategyReply{}, nil
}
