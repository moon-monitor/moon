package impl

import (
	"context"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/laurel/internal/data"
	"github.com/prometheus/client_golang/prometheus"
)

func NewMetricRegister(data *data.Data) repository.MetricRegister {
	return &metricRegisterImpl{
		Data: data,
	}
}

type metricRegisterImpl struct {
	*data.Data
}

// RegisterCounterMetric implements repository.MetricRegister.
// Subtle: this method shadows the method (*Data).RegisterCounterMetric of metricRegisterImpl.Data.
func (m *metricRegisterImpl) RegisterCounterMetric(ctx context.Context, name string, metric *prometheus.CounterVec) {
	m.SetCounterMetric(name, metric)
	prometheus.MustRegister(metric)
}

// RegisterGaugeMetric implements repository.MetricRegister.
// Subtle: this method shadows the method (*Data).RegisterGaugeMetric of metricRegisterImpl.Data.
func (m *metricRegisterImpl) RegisterGaugeMetric(ctx context.Context, name string, metric *prometheus.GaugeVec) {
	m.SetGaugeMetric(name, metric)
	prometheus.MustRegister(metric)
}

// RegisterHistogramMetric implements repository.MetricRegister.
// Subtle: this method shadows the method (*Data).RegisterHistogramMetric of metricRegisterImpl.Data.
func (m *metricRegisterImpl) RegisterHistogramMetric(ctx context.Context, name string, metric *prometheus.HistogramVec) {
	m.SetHistogramMetric(name, metric)
	prometheus.MustRegister(metric)
}

// RegisterSummaryMetric implements repository.MetricRegister.
// Subtle: this method shadows the method (*Data).RegisterSummaryMetric of metricRegisterImpl.Data.
func (m *metricRegisterImpl) RegisterSummaryMetric(ctx context.Context, name string, metric *prometheus.SummaryVec) {
	m.SetSummaryMetric(name, metric)
	prometheus.MustRegister(metric)
}
