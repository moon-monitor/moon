package service

import (
	"context"

	pb "github.com/moon-monitor/moon/pkg/api/common"
)

type ServerService struct {
	pb.UnimplementedServerServer
}

func NewServerService() *ServerService {
	return &ServerService{}
}

func (s *ServerService) Register(ctx context.Context, req *pb.ServerRegisterRequest) (*pb.ServerRegisterReply, error) {
	return &pb.ServerRegisterReply{}, nil
}
