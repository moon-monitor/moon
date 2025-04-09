package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/plugin/datasource/prometheus"
)

func NewMetricRepo(d *data.Data, logger log.Logger) repository.MetricInit {
	return &metricImpl{
		Data:   d,
		logger: logger,
		help:   log.NewHelper(log.With(logger, "module", "data.repo.metric")),
	}
}

type metricImpl struct {
	*data.Data
	help   *log.Helper
	logger log.Logger
}

type metricInstance struct {
	metric datasource.Metric
}

func (m *metricImpl) Init(config bo.MetricDatasourceConfig) (repository.Metric, error) {
	if config == nil {
		return nil, merr.ErrorInvalidArgument("metric datasource config is nil")
	}

	var (
		metricDatasource datasource.Metric
		ok               bool
	)

	metricDatasource, ok = m.GetMetricDatasource(config.UniqueKey())
	switch config.GetDriver() {
	case common.MetricDatasourceItem_Driver_PROMETHEUS:
		if !ok {
			metricDatasource = prometheus.New(config, m.logger)
		}
	case common.MetricDatasourceItem_Driver_VICTORIA_METRICS:
		if !ok {
			metricDatasource = prometheus.New(config, m.logger)
		}
	default:
		return nil, merr.ErrorParamsError("invalid metric datasource driver: %s", config.GetDriver())
	}
	return &metricInstance{metric: metricDatasource}, nil
}

func (m *metricInstance) Query(ctx context.Context, expr string, duration time.Duration) ([]*do.MetricQueryReply, error) {
	//TODO implement me
	panic("implement me")
}

func (m *metricInstance) QueryRange(ctx context.Context, expr string, start, end int64) ([]*do.MetricQueryRangeReply, error) {
	// 分辨率计算
	//TODO implement me
	panic("implement me")
}

func (m *metricInstance) Metadata(ctx context.Context) (<-chan []*do.MetricItem, error) {
	//TODO implement me
	panic("implement me")
}
