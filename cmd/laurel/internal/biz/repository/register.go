package repository

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricRegister interface {
	RegisterCounterMetric(ctx context.Context, name string, metric *prometheus.CounterVec)
	RegisterGaugeMetric(ctx context.Context, name string, metric *prometheus.GaugeVec)
	RegisterHistogramMetric(ctx context.Context, name string, metric *prometheus.HistogramVec)
	RegisterSummaryMetric(ctx context.Context, name string, metric *prometheus.SummaryVec)
}
