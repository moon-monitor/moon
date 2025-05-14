package repository

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
)

type MetricRegister interface {
	RegisterCounterMetric(ctx context.Context, name string, metric *prometheus.CounterVec)
	RegisterGaugeMetric(ctx context.Context, name string, metric *prometheus.GaugeVec)
	RegisterHistogramMetric(ctx context.Context, name string, metric *prometheus.HistogramVec)
	RegisterSummaryMetric(ctx context.Context, name string, metric *prometheus.SummaryVec)

	WithCounterMetricValue(ctx context.Context, metrics ...*bo.MetricData)
	WithGaugeMetricValue(ctx context.Context, metrics ...*bo.MetricData)
	WithHistogramMetricValue(ctx context.Context, metrics ...*bo.MetricData)
	WithSummaryMetricValue(ctx context.Context, metrics ...*bo.MetricData)
}
