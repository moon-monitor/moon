package service

import (
	"context"

	common "github.com/moon-monitor/moon/pkg/api/common"
	apicommon "github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type AlertService struct {
	apiv1.UnimplementedAlertServer
}

func NewAlertService() *AlertService {
	return &AlertService{}
}

func (s *AlertService) SendAlert(ctx context.Context, req *common.AlertsItem) (*apicommon.EmptyReply, error) {
	return &apicommon.EmptyReply{}, nil
}
