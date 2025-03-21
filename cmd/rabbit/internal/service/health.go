package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	"github.com/moon-monitor/moon/pkg/api/common"
)

type HealthService struct {
	common.UnimplementedHealthServer

	healthBiz   *biz.HealthBiz
	registerBiz *biz.RegisterBiz
	helper      *log.Helper
}

func NewHealthService(healthBiz *biz.HealthBiz, registerBiz *biz.RegisterBiz, logger log.Logger) *HealthService {
	return &HealthService{
		healthBiz:   healthBiz,
		registerBiz: registerBiz,
		helper:      log.NewHelper(log.With(logger, "module", "service.health")),
	}
}

func (s *HealthService) Check(ctx context.Context, req *common.CheckRequest) (*common.CheckReply, error) {
	return &common.CheckReply{}, nil
}

func (s *HealthService) Online(ctx context.Context) error {
	return s.registerBiz.Online(ctx)
}

func (s *HealthService) Offline(ctx context.Context) error {
	return s.registerBiz.Offline(ctx)
}
