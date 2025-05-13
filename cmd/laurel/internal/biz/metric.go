package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewMetricManager(
	metricRegisterRepo repository.MetricRegister,
	cacheRepo repository.Cache,
	logger log.Logger,
) *MetricManager {
	metricManager := &MetricManager{
		metricRegisterRepo: metricRegisterRepo,
		cacheRepo:          cacheRepo,
		helper:             log.NewHelper(log.With(logger, "module", "biz.metric")),
	}
	defer metricManager.loadMetrics()
	return metricManager
}

type MetricManager struct {
	metricRegisterRepo repository.MetricRegister
	cacheRepo          repository.Cache

	helper *log.Helper
}

func (m *MetricManager) loadMetrics() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	eg := new(errgroup.Group)
	eg.Go(func() error {
		return m.loadCounterMetrics(ctx)
	})
	eg.Go(func() error {
		return m.loadGaugeMetrics(ctx)
	})
	eg.Go(func() error {
		return m.loadHistogramMetrics(ctx)
	})
	eg.Go(func() error {
		return m.loadSummaryMetrics(ctx)
	})
	if err := eg.Wait(); err != nil {
		m.helper.Errorw("msg", "load metrics error", "error", err)
	}
}

func (m *MetricManager) loadCounterMetrics(ctx context.Context) error {
	counterMetrics, err := m.cacheRepo.GetCounterMetrics(ctx)
	if err != nil {
		return err
	}
	for _, metric := range counterMetrics {
		m.metricRegisterRepo.RegisterCounterMetric(ctx, metric.GetMetricName(), metric.New())
	}
	return nil
}

func (m *MetricManager) loadGaugeMetrics(ctx context.Context) error {
	gaugeMetrics, err := m.cacheRepo.GetGaugeMetrics(ctx)
	if err != nil {
		return err
	}
	for _, metric := range gaugeMetrics {
		m.metricRegisterRepo.RegisterGaugeMetric(ctx, metric.GetMetricName(), metric.New())
	}
	return nil
}

func (m *MetricManager) loadHistogramMetrics(ctx context.Context) error {
	histogramMetrics, err := m.cacheRepo.GetHistogramMetrics(ctx)
	if err != nil {
		return err
	}
	for _, metric := range histogramMetrics {
		m.metricRegisterRepo.RegisterHistogramMetric(ctx, metric.GetMetricName(), metric.New())
	}
	return nil
}

func (m *MetricManager) loadSummaryMetrics(ctx context.Context) error {
	summaryMetrics, err := m.cacheRepo.GetSummaryMetrics(ctx)
	if err != nil {
		return err
	}
	for _, metric := range summaryMetrics {
		m.metricRegisterRepo.RegisterSummaryMetric(ctx, metric.GetMetricName(), metric.New())
	}
	return nil
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
