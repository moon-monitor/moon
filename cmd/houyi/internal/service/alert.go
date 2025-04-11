package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	houyiv1 "github.com/moon-monitor/moon/pkg/api/houyi/v1"
)

type AlertService struct {
	houyiv1.UnimplementedAlertServer

	helper *log.Helper
}

func NewAlertService(logger log.Logger) *AlertService {
	return &AlertService{
		helper: log.NewHelper(log.With(logger, "module", "service.alert")),
	}
}

func (s *AlertService) Push(ctx context.Context, req *houyiv1.PushAlertRequest) (*houyiv1.PushAlertReply, error) {
	return &houyiv1.PushAlertReply{}, nil
}
