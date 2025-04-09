package service

import (
	"context"

	houyiv1 "github.com/moon-monitor/moon/pkg/api/houyi/v1"
)

type SyncService struct {
	houyiv1.UnimplementedSyncServer
}

func NewSyncService() *SyncService {
	return &SyncService{}
}

func (s *SyncService) MetricStrategy(ctx context.Context, req *houyiv1.MetricStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) CertificateStrategy(ctx context.Context, req *houyiv1.CertificateStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) ServerPortStrategy(ctx context.Context, req *houyiv1.ServerPortStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) HttpStrategy(ctx context.Context, req *houyiv1.HttpStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) PingStrategy(ctx context.Context, req *houyiv1.PingStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) EventStrategy(ctx context.Context, req *houyiv1.EventStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) LogsStrategy(ctx context.Context, req *houyiv1.LogsStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}

func (s *SyncService) RemoveStrategy(ctx context.Context, req *houyiv1.RemoveStrategyRequest) (*houyiv1.SyncStrategyReply, error) {
	return &houyiv1.SyncStrategyReply{}, nil
}
