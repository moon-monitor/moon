package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/laurel/internal/data"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewCacheRepo(d *data.Data, logger log.Logger) repository.Cache {
	return &cacheImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.cache")),
	}
}

type cacheImpl struct {
	*data.Data

	helper *log.Helper
}

// StorageMetric implements repository.Cache.
func (c *cacheImpl) StorageMetric(ctx context.Context, metrics ...bo.MetricVec) error {
	metricsByType := slices.GroupBy(metrics, func(metric bo.MetricVec) vobj.MetricType {
		return metric.Type()
	})
	counterMetrics := make([]bo.MetricVec, 0, len(metricsByType[vobj.MetricTypeCounter]))
	gaugeMetrics := make([]bo.MetricVec, 0, len(metricsByType[vobj.MetricTypeGauge]))
	histogramMetrics := make([]bo.MetricVec, 0, len(metricsByType[vobj.MetricTypeHistogram]))
	summaryMetrics := make([]bo.MetricVec, 0, len(metricsByType[vobj.MetricTypeSummary]))
	for metricType, metrics := range metricsByType {
		switch metricType {
		case vobj.MetricTypeCounter:
			counterMetrics = append(counterMetrics, metrics...)
		case vobj.MetricTypeGauge:
			gaugeMetrics = append(gaugeMetrics, metrics...)
		case vobj.MetricTypeHistogram:
			histogramMetrics = append(histogramMetrics, metrics...)
		case vobj.MetricTypeSummary:
			summaryMetrics = append(summaryMetrics, metrics...)
		}
	}
	eg := new(errgroup.Group)
	if len(counterMetrics) > 0 {
		eg.Go(func() error {
			key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeCounter)
			values := slices.ToMap(counterMetrics, func(metric bo.MetricVec) string {
				return metric.GetMetricName()
			})
			return c.Data.GetCache().Client().HSet(ctx, key, values).Err()
		})
	}
	if len(gaugeMetrics) > 0 {
		eg.Go(func() error {
			key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeGauge)
			values := slices.ToMap(gaugeMetrics, func(metric bo.MetricVec) string {
				return metric.GetMetricName()
			})
			return c.Data.GetCache().Client().HSet(ctx, key, values).Err()
		})
	}
	if len(histogramMetrics) > 0 {
		eg.Go(func() error {
			key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeHistogram)
			values := slices.ToMap(histogramMetrics, func(metric bo.MetricVec) string {
				return metric.GetMetricName()
			})
			return c.Data.GetCache().Client().HSet(ctx, key, values).Err()
		})
	}
	if len(summaryMetrics) > 0 {
		eg.Go(func() error {
			key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeSummary)
			values := slices.ToMap(summaryMetrics, func(metric bo.MetricVec) string {
				return metric.GetMetricName()
			})
			return c.Data.GetCache().Client().HSet(ctx, key, values).Err()
		})
	}

	return eg.Wait()
}

func (c *cacheImpl) GetCounterMetrics(ctx context.Context) ([]*bo.CounterMetricVec, error) {
	key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeCounter)
	values, err := c.Data.GetCache().Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	metrics := make([]*bo.CounterMetricVec, 0, len(values))
	for _, metricValue := range values {
		var metric bo.CounterMetricVec
		err := metric.UnmarshalBinary([]byte(metricValue))
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &metric)
	}
	return metrics, nil
}

func (c *cacheImpl) GetGaugeMetrics(ctx context.Context) ([]*bo.GaugeMetricVec, error) {
	key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeGauge)
	values, err := c.Data.GetCache().Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	metrics := make([]*bo.GaugeMetricVec, 0, len(values))
	for _, metricValue := range values {
		var metric bo.GaugeMetricVec
		err := metric.UnmarshalBinary([]byte(metricValue))
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &metric)
	}
	return metrics, nil
}

func (c *cacheImpl) GetHistogramMetrics(ctx context.Context) ([]*bo.HistogramMetricVec, error) {
	key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeHistogram)
	values, err := c.Data.GetCache().Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	metrics := make([]*bo.HistogramMetricVec, 0, len(values))
	for _, metricValue := range values {
		var metric bo.HistogramMetricVec
		err := metric.UnmarshalBinary([]byte(metricValue))
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &metric)
	}
	return metrics, nil
}

func (c *cacheImpl) GetSummaryMetrics(ctx context.Context) ([]*bo.SummaryMetricVec, error) {
	key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeSummary)
	values, err := c.Data.GetCache().Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	metrics := make([]*bo.SummaryMetricVec, 0, len(values))
	for _, metricValue := range values {
		var metric bo.SummaryMetricVec
		err := metric.UnmarshalBinary([]byte(metricValue))
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &metric)
	}
	return metrics, nil
}

func (c *cacheImpl) GetMetric(ctx context.Context, metricType vobj.MetricType, metricName string) (bo.MetricVec, error) {
	switch metricType {
	case vobj.MetricTypeCounter:
		key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeCounter)
		var metric bo.CounterMetricVec
		err := c.Data.GetCache().Client().HGet(ctx, key, metricName).Scan(&metric)
		if err != nil {
			return nil, err
		}
		return &metric, nil
	case vobj.MetricTypeGauge:
		key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeGauge)
		var metric bo.GaugeMetricVec
		err := c.Data.GetCache().Client().HGet(ctx, key, metricName).Scan(&metric)
		if err != nil {
			return nil, err
		}
		return &metric, nil
	case vobj.MetricTypeHistogram:
		key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeHistogram)
		var metric bo.HistogramMetricVec
		err := c.Data.GetCache().Client().HGet(ctx, key, metricName).Scan(&metric)
		if err != nil {
			return nil, err
		}
		return &metric, nil
	case vobj.MetricTypeSummary:
		key := vobj.MetricCacheKeyPrefix.Key(vobj.MetricTypeSummary)
		var metric bo.SummaryMetricVec
		err := c.Data.GetCache().Client().HGet(ctx, key, metricName).Scan(&metric)
		if err != nil {
			return nil, err
		}
		return &metric, nil
	default:
		return nil, merr.ErrorParamsError("invalid metric type: %s", metricType)
	}
}

func (c *cacheImpl) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.Data.GetCache().Client().SetNX(ctx, key, 1, expiration).Result()
}

func (c *cacheImpl) Unlock(ctx context.Context, key string) error {
	return c.Data.GetCache().Client().Del(ctx, key).Err()
}
