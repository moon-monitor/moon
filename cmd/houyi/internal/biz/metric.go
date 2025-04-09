package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewMetric(
	judgeRepo repository.Judge,
	alertRepo repository.Alert,
	metricInitRepo repository.MetricInit,
	configRepo repository.Config,
	logger log.Logger,
) *Metric {
	return &Metric{
		helper:         log.NewHelper(log.With(logger, "module", "biz.metric")),
		judgeRepo:      judgeRepo,
		alertRepo:      alertRepo,
		metricInitRepo: metricInitRepo,
		configRepo:     configRepo,
	}
}

type Metric struct {
	helper         *log.Helper
	judgeRepo      repository.Judge
	alertRepo      repository.Alert
	metricInitRepo repository.MetricInit
	configRepo     repository.Config
}

func (m *Metric) GetAlerts(ctx context.Context) ([]bo.Alert, error) {
	return nil, nil
}

func (m *Metric) Evaluate(ctx context.Context, metric bo.MetricRule) error {
	datasourceConfig, ok := m.configRepo.GetMetricDatasourceConfig(ctx, metric.GetDatasource())
	if !ok {
		m.helper.Warnw("method", "Evaluate", "msg", "datasource config not found")
		return nil
	}
	query, err := m.metricInitRepo.Init(datasourceConfig)
	if err != nil {
		return err
	}

	end := time.Now()
	start := end.Add(-metric.GetDuration())
	queryRange, err := query.QueryRange(ctx, metric.GetExpr(), start.Unix(), end.Unix())
	if err != nil {
		return err
	}
	metricJudgeData := slices.Map(queryRange, func(dataItem *do.MetricQueryRangeReply) bo.MetricJudgeData {
		return dataItem
	})

	alerts, err := m.judgeRepo.Metric(ctx, metricJudgeData, metric)
	if err != nil {
		return err
	}
	if err := m.alertRepo.Save(ctx, alerts...); err != nil {
		return err
	}
	return nil
}
