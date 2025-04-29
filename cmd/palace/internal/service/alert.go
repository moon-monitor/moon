package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	apicommon "github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func NewAlertService(realtimeBiz *biz.Realtime) *AlertService {
	return &AlertService{
		realtimeBiz: realtimeBiz,
	}
}

type AlertService struct {
	palace.UnimplementedAlertServer
	realtimeBiz *biz.Realtime
}

func (s *AlertService) PushAlert(ctx context.Context, req *apicommon.AlertItem) (*common.EmptyReply, error) {
	params := build.ToAlertParams(req)
	if err := s.realtimeBiz.SaveAlert(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}
