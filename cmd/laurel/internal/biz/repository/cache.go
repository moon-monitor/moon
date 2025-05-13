package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/vobj"
)

type Cache interface {
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) error

	StorageMetric(ctx context.Context, metrics ...bo.MetricVec) error
	GetCounterMetrics(ctx context.Context) ([]*bo.CounterMetricVec, error)
	GetGaugeMetrics(ctx context.Context) ([]*bo.GaugeMetricVec, error)
	GetHistogramMetrics(ctx context.Context) ([]*bo.HistogramMetricVec, error)
	GetSummaryMetrics(ctx context.Context) ([]*bo.SummaryMetricVec, error)
	GetMetric(ctx context.Context, metricType vobj.MetricType, metricName string) (bo.MetricVec, error)
}
