package service

import (
	"context"

	commonapi "github.com/moon-monitor/moon/pkg/api/common"
)

type HealthService struct {
	commonapi.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check(ctx context.Context, req *commonapi.CheckRequest) (*commonapi.CheckReply, error) {
	return &commonapi.CheckReply{}, nil
}
