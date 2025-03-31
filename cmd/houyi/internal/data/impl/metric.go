package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/plugin/datasource/prometheus"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

func NewMetricRepo(logger log.Logger) repository.MetricInit {
	return &metricImpl{
		instances: safety.NewMap[string, *metricInstance](),
		logger:    logger,
		help:      log.NewHelper(log.With(logger, "module", "data.repo.metric")),
	}
}

type metricImpl struct {
	help      *log.Helper
	logger    log.Logger
	instances *safety.Map[string, *metricInstance]
}

type metricInstance struct {
	datasource *bo.MetricDatasourceItem
	metric     datasource.Metric
}

func (m *metricImpl) Init(config *bo.MetricDatasourceItem) (instance repository.Metric, err error) {
	if config == nil {
		return nil, merr.ErrorInvalidArgument("metric datasource config is nil")
	}
	var ok bool

	switch config.Driver {
	case vobj.MetricDatasourceDriverPrometheus:
		instance, ok = m.instances.Get(config.Prometheus.GetEndpoint())
		if ok {
			return instance, nil
		}
		instance = &metricInstance{
			datasource: config,
			metric:     prometheus.New(config.Prometheus, m.logger),
		}
		m.instances.Set(config.Prometheus.GetEndpoint(), instance.(*metricInstance))
		return instance, nil
	case vobj.MetricDatasourceDriverVictoriaMetrics:
		return instance, nil
	default:
		return nil, merr.ErrorParamsError("invalid metric datasource driver: %s", config.Driver)
	}
}

func (m *metricInstance) Query(ctx context.Context, expr string, duration time.Duration) ([]*bo.MetricQueryReply, error) {
	//TODO implement me
	panic("implement me")
}

func (m *metricInstance) QueryRange(ctx context.Context, expr string, start, end int64, step uint32) ([]*bo.MetricQueryRangeReply, error) {
	//TODO implement me
	panic("implement me")
}

func (m *metricInstance) Metadata(ctx context.Context) (<-chan []*bo.MetricItem, error) {
	//TODO implement me
	panic("implement me")
}
