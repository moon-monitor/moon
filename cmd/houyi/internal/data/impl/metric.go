package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
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
	helper *log.Helper
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
	case common.MetricDatasourceDriver_prometheus:
		if !ok {
			metricDatasource = prometheus.New(config, m.logger)
		}
	case common.MetricDatasourceDriver_victoriametrics:
		if !ok {
			metricDatasource = prometheus.New(config, m.logger)
		}
	default:
		return nil, merr.ErrorParamsError("invalid metric datasource driver: %s", config.GetDriver())
	}
	return &metricInstance{
		metric: metricDatasource,
		helper: log.NewHelper(log.With(m.logger, "module", "data.repo.metric.instance")),
	}, nil
}

func (m *metricInstance) Query(ctx context.Context, expr string, t time.Time) ([]*do.MetricQueryReply, error) {
	queryParams := &datasource.MetricQueryRequest{
		Expr: expr,
		Time: t.Unix(),
	}
	metricQueryResponse, err := m.metric.Query(ctx, queryParams)
	if err != nil {
		m.helper.Warnw("msg", "query metric failed", "err", err)
		return nil, err
	}
	list := make([]*do.MetricQueryReply, 0, len(metricQueryResponse.Data.Result))
	for _, result := range metricQueryResponse.Data.Result {
		queryValue := result.GetMetricQueryValue()
		item := &do.MetricQueryReply{
			Labels: result.Metric,
			Value: &do.MetricQueryValue{
				Value:     queryValue.Value,
				Timestamp: int64(queryValue.Timestamp),
			},
			ResultType: string(metricQueryResponse.Data.ResultType),
		}
		list = append(list, item)
	}
	return list, nil
}

func (m *metricInstance) QueryRange(ctx context.Context, expr string, start, end time.Time) ([]*do.MetricQueryRangeReply, error) {
	// 分辨率计算
	step := m.getOptimalStep(start, end)
	queryParams := &datasource.MetricQueryRequest{
		Expr:      expr,
		Time:      0,
		StartTime: start.Unix(),
		EndTime:   end.Unix(),
		Step:      uint32(step.Seconds()),
	}
	metricQueryResponse, err := m.metric.Query(ctx, queryParams)
	if err != nil {
		m.helper.Warnw("msg", "query metric range failed", "err", err)
		return nil, err
	}
	list := make([]*do.MetricQueryRangeReply, 0, len(metricQueryResponse.Data.Result))
	for _, result := range metricQueryResponse.Data.Result {
		queryValues := result.GetMetricQueryValues()
		for _, queryValue := range queryValues {
			item := &do.MetricQueryRangeReply{
				Labels: result.Metric,
				Values: []*do.MetricQueryValue{
					{
						Value:     queryValue.Value,
						Timestamp: int64(queryValue.Timestamp),
					},
				},
				ResultType: string(metricQueryResponse.Data.ResultType),
			}
			list = append(list, item)
		}
	}
	return list, nil
}

func (m *metricInstance) Metadata(ctx context.Context) (<-chan []*do.MetricItem, error) {
	metricMetadata, err := m.metric.Metadata(ctx)
	if err != nil {
		m.helper.Warnw("msg", "get metric metadata failed", "err", err)
		return nil, err
	}
	ch := make(chan []*do.MetricItem)
	go func() {
		defer func() {
			close(ch)
			if r := recover(); r != nil {
				m.helper.Errorw("msg", "panic occurred", "err", r)
			}
		}()
		for metadata := range metricMetadata {
			syncList := make([]*do.MetricItem, 0, len(metadata.Metric))
			for _, metricMetadataItem := range metadata.Metric {
				item := &do.MetricItem{
					Name:   metricMetadataItem.Name,
					Help:   metricMetadataItem.Help,
					Type:   metricMetadataItem.Type,
					Labels: metricMetadataItem.Labels,
					Unit:   metricMetadataItem.Unit,
				}
				syncList = append(syncList, item)
			}
			ch <- syncList
		}
	}()
	return ch, nil
}

func (m *metricInstance) getOptimalStep(start, end time.Time) time.Duration {
	duration := end.Sub(start)

	// Prometheus 通常会对较旧的数据进行降采样
	if duration > 15*24*time.Hour {
		// 对于超过15天的数据，使用较大的step
		return 2 * time.Hour
	} else if duration > 3*24*time.Hour {
		return 1 * time.Hour
	}

	// 对于近期数据，尝试匹配采集间隔
	scrapeInterval := m.metric.GetScrapeInterval()

	// 确保step至少是scrape_interval的倍数
	minStep := scrapeInterval

	// 计算一个合理的step，使返回点数在500-1000之间
	desiredPoints := 800
	calculatedStep := duration / time.Duration(desiredPoints)

	// 确保step不小于最小step，且是scrapeInterval的倍数
	if calculatedStep < minStep {
		return minStep
	}

	// 向上取整到scrapeInterval的倍数
	return ((calculatedStep + scrapeInterval - 1) / scrapeInterval) * scrapeInterval
}
