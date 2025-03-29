package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/common"
)

type HealthService struct {
	pb.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckReply, error) {
	return &pb.CheckReply{}, nil
}
