package impl

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/laurel/internal/data"
)

func NewMetricRegister(data *data.Data) repository.MetricRegister {
	return &metricRegisterImpl{
		Data: data,
	}
}

type metricRegisterImpl struct {
	*data.Data
}

// WithCounterMetricValue implements repository.MetricRegister.
func (m *metricRegisterImpl) WithCounterMetricValue(ctx context.Context, metrics ...*bo.MetricData) {
	for _, metric := range metrics {
		counterVec, ok := m.GetCounterMetric(metric.Name)
		if !ok {
			continue
		}
		counterVec.With(metric.Labels).Add(metric.Value)
	}
}

// WithGaugeMetricValue implements repository.MetricRegister.
func (m *metricRegisterImpl) WithGaugeMetricValue(ctx context.Context, metrics ...*bo.MetricData) {
	for _, metric := range metrics {
		gaugeVec, ok := m.GetGaugeMetric(metric.Name)
		if !ok {
			continue
		}
		gaugeVec.With(metric.Labels).Set(metric.Value)
	}
}

// WithHistogramMetricValue implements repository.MetricRegister.
func (m *metricRegisterImpl) WithHistogramMetricValue(ctx context.Context, metrics ...*bo.MetricData) {
	for _, metric := range metrics {
		histogramVec, ok := m.GetHistogramMetric(metric.Name)
		if !ok {
			continue
		}
		histogramVec.With(metric.Labels).Observe(metric.Value)
	}
}

// WithSummaryMetricValue implements repository.MetricRegister.
func (m *metricRegisterImpl) WithSummaryMetricValue(ctx context.Context, metrics ...*bo.MetricData) {
	for _, metric := range metrics {
		summaryVec, ok := m.GetSummaryMetric(metric.Name)
		if !ok {
			continue
		}
		summaryVec.With(metric.Labels).Observe(metric.Value)
	}
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
