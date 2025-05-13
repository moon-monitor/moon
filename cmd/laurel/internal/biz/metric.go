package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewMetricManager(metricRegisterRepo repository.MetricRegister, cacheRepo repository.Cache) *MetricManager {
	return &MetricManager{
		metricRegisterRepo: metricRegisterRepo,
		cacheRepo:          cacheRepo,
	}
}

type MetricManager struct {
	metricRegisterRepo repository.MetricRegister
	cacheRepo          repository.Cache
}

func (m *MetricManager) RegisterCounterMetric(ctx context.Context, metrics ...*bo.CounterMetricVec) error {
	if len(metrics) == 0 {
		return nil
	}
	cacheMetrics := slices.Map(metrics, func(metric *bo.CounterMetricVec) bo.MetricVec {
		return metric
	})
	if err := m.cacheRepo.StorageMetric(ctx, cacheMetrics...); err != nil {
		return err
	}
	for _, metric := range metrics {
		metricValue := metric.New()
		m.metricRegisterRepo.RegisterCounterMetric(ctx, metric.GetMetricName(), metricValue)
	}
	return nil
}

func (m *MetricManager) RegisterGaugeMetric(ctx context.Context, metrics ...*bo.GaugeMetricVec) error {
	if len(metrics) == 0 {
		return nil
	}
	cacheMetrics := slices.Map(metrics, func(metric *bo.GaugeMetricVec) bo.MetricVec {
		return metric
	})
	if err := m.cacheRepo.StorageMetric(ctx, cacheMetrics...); err != nil {
		return err
	}
	for _, metric := range metrics {
		metricValue := metric.New()
		m.metricRegisterRepo.RegisterGaugeMetric(ctx, metric.GetMetricName(), metricValue)
	}
	return nil
}

func (m *MetricManager) RegisterHistogramMetric(ctx context.Context, metrics ...*bo.HistogramMetricVec) error {
	if len(metrics) == 0 {
		return nil
	}
	cacheMetrics := slices.Map(metrics, func(metric *bo.HistogramMetricVec) bo.MetricVec {
		return metric
	})
	if err := m.cacheRepo.StorageMetric(ctx, cacheMetrics...); err != nil {
		return err
	}
	for _, metric := range metrics {
		metricValue := metric.New()
		m.metricRegisterRepo.RegisterHistogramMetric(ctx, metric.GetMetricName(), metricValue)
	}
	return nil
}

func (m *MetricManager) RegisterSummaryMetric(ctx context.Context, metrics ...*bo.SummaryMetricVec) error {
	if len(metrics) == 0 {
		return nil
	}
	cacheMetrics := slices.Map(metrics, func(metric *bo.SummaryMetricVec) bo.MetricVec {
		return metric
	})
	if err := m.cacheRepo.StorageMetric(ctx, cacheMetrics...); err != nil {
		return err
	}
	for _, metric := range metrics {
		metricValue := metric.New()
		m.metricRegisterRepo.RegisterSummaryMetric(ctx, metric.GetMetricName(), metricValue)
	}
	return nil
}
