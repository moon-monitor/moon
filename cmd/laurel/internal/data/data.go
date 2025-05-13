package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/moon-monitor/moon/cmd/laurel/internal/conf"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

func New(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	var err error

	data := &Data{
		counterMetrics:   safety.NewMap[string, *prometheus.CounterVec](),
		gaugeMetrics:     safety.NewMap[string, *prometheus.GaugeVec](),
		histogramMetrics: safety.NewMap[string, *prometheus.HistogramVec](),
		summaryMetrics:   safety.NewMap[string, *prometheus.SummaryVec](),
		helper:           log.NewHelper(log.With(logger, "module", "data")),
	}
	data.cache, err = cache.NewCache(c.GetCache())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		if err = data.cache.Close(); err != nil {
			log.NewHelper(logger).Errorw("method", "close cache", "err", err)
		}
	}
	return data, cleanup, nil
}

type Data struct {
	cache            cache.Cache
	counterMetrics   *safety.Map[string, *prometheus.CounterVec]
	gaugeMetrics     *safety.Map[string, *prometheus.GaugeVec]
	histogramMetrics *safety.Map[string, *prometheus.HistogramVec]
	summaryMetrics   *safety.Map[string, *prometheus.SummaryVec]
	helper           *log.Helper
}

func (d *Data) GetCache() cache.Cache {
	return d.cache
}

func (d *Data) GetCounterMetrics() *safety.Map[string, *prometheus.CounterVec] {
	return d.counterMetrics
}

func (d *Data) GetGaugeMetrics() *safety.Map[string, *prometheus.GaugeVec] {
	return d.gaugeMetrics
}

func (d *Data) GetHistogramMetrics() *safety.Map[string, *prometheus.HistogramVec] {
	return d.histogramMetrics
}

func (d *Data) GetSummaryMetrics() *safety.Map[string, *prometheus.SummaryVec] {
	return d.summaryMetrics
}

func (d *Data) SetCounterMetric(name string, metrics *prometheus.CounterVec) {
	d.counterMetrics.Set(name, metrics)
}

func (d *Data) SetGaugeMetric(name string, metrics *prometheus.GaugeVec) {
	d.gaugeMetrics.Set(name, metrics)
}

func (d *Data) SetHistogramMetric(name string, metrics *prometheus.HistogramVec) {
	d.histogramMetrics.Set(name, metrics)
}

func (d *Data) SetSummaryMetric(name string, metrics *prometheus.SummaryVec) {
	d.summaryMetrics.Set(name, metrics)
}

func (d *Data) GetCounterMetric(name string) (*prometheus.CounterVec, bool) {
	return d.counterMetrics.Get(name)
}

func (d *Data) GetGaugeMetric(name string) (*prometheus.GaugeVec, bool) {
	return d.gaugeMetrics.Get(name)
}

func (d *Data) GetHistogramMetric(name string) (*prometheus.HistogramVec, bool) {
	return d.histogramMetrics.Get(name)
}

func (d *Data) GetSummaryMetric(name string) (*prometheus.SummaryVec, bool) {
	return d.summaryMetrics.Get(name)
}
