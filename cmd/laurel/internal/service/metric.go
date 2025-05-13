package service

import (
	"context"

	apiv1 "github.com/moon-monitor/moon/pkg/api/laurel/v1"
)

func NewMetricService() *MetricService {
	return &MetricService{}
}

type MetricService struct {
	apiv1.UnimplementedMetricServer
}

func (s *MetricService) Push(ctx context.Context, req *apiv1.MetricPushRequest) (*apiv1.MetricPushReply, error) {
	return &apiv1.MetricPushReply{}, nil
}
