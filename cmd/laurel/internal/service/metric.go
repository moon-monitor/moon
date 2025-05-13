package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz"
	"github.com/moon-monitor/moon/cmd/laurel/internal/service/build"
	apiv1 "github.com/moon-monitor/moon/pkg/api/laurel/v1"
)

func NewMetricService(metricManager *biz.MetricManager) *MetricService {
	return &MetricService{
		metricManager: metricManager,
	}
}

type MetricService struct {
	apiv1.UnimplementedMetricServer
	metricManager *biz.MetricManager
}

func (s *MetricService) PushMetricData(ctx context.Context, req *apiv1.PushMetricDataRequest) (*apiv1.EmptyReply, error) {
	return &apiv1.EmptyReply{}, nil
}

func (s *MetricService) RegisterMetric(ctx context.Context, req *apiv1.RegisterMetricRequest) (*apiv1.EmptyReply, error) {
	counterVecs := build.ToCounterMetricVecs(req.CounterVecs)
	gaugeVecs := build.ToGaugeMetricVecs(req.GaugeVecs)
	histogramVecs := build.ToHistogramMetricVecs(req.HistogramVecs)
	summaryVecs := build.ToSummaryMetricVecs(req.SummaryVecs)

	s.metricManager.RegisterCounterMetric(ctx, counterVecs...)
	s.metricManager.RegisterGaugeMetric(ctx, gaugeVecs...)
	s.metricManager.RegisterHistogramMetric(ctx, histogramVecs...)
	s.metricManager.RegisterSummaryMetric(ctx, summaryVecs...)

	return &apiv1.EmptyReply{}, nil
}
