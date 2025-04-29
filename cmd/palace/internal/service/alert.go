package service

import (
	"context"

	apicommon "github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func NewAlertService() *AlertService {
	return &AlertService{}
}

type AlertService struct {
	palace.UnimplementedAlertServer
}

func (s *AlertService) PushAlert(ctx context.Context, req *apicommon.AlertItem) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
