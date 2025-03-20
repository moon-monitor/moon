package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	commonapi "github.com/moon-monitor/moon/pkg/api/common"
)

type HealthService struct {
	commonapi.UnimplementedHealthServer

	healthBiz *biz.HealthBiz
	helper    *log.Helper
}

func NewHealthService(healthBiz *biz.HealthBiz, logger log.Logger) *HealthService {
	return &HealthService{
		healthBiz: healthBiz,
		helper:    log.NewHelper(log.With(logger, "module", "service.health")),
	}
}

func (s *HealthService) Check(ctx context.Context, req *commonapi.CheckRequest) (*commonapi.CheckReply, error) {
	return &commonapi.CheckReply{}, nil
}
