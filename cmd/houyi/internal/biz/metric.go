package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/conf"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

func NewMetric(
	bc *conf.Bootstrap,
	judgeRepo repository.Judge,
	alertRepo repository.Alert,
	metricInitRepo repository.MetricInit,
	configRepo repository.Config,
	logger log.Logger,
) *Metric {
	evaluateConf := bc.GetEvaluate()
	return &Metric{
		logger:           logger,
		helper:           log.NewHelper(log.With(logger, "module", "biz.metric")),
		judgeRepo:        judgeRepo,
		alertRepo:        alertRepo,
		metricInitRepo:   metricInitRepo,
		configRepo:       configRepo,
		evaluateInterval: evaluateConf.GetInterval().AsDuration(),
		evaluateTimeout:  evaluateConf.GetTimeout().AsDuration(),
	}
}

type Metric struct {
	logger           log.Logger
	helper           *log.Helper
	judgeRepo        repository.Judge
	alertRepo        repository.Alert
	metricInitRepo   repository.MetricInit
	configRepo       repository.Config
	evaluateInterval time.Duration
	evaluateTimeout  time.Duration
}

func (m *Metric) NewStrategyJob(_ context.Context, metric bo.MetricRule) (bo.StrategyJob, error) {
	opts := []event.StrategyMetricJobOption{
		event.WithStrategyMetricJobHelper(m.logger),
		event.WithStrategyMetricJobMetric(metric.UniqueKey(), metric.GetEnable()),
		event.WithStrategyMetricJobConfigRepo(m.configRepo),
		event.WithStrategyMetricJobJudgeRepo(m.judgeRepo),
		event.WithStrategyMetricJobAlertRepo(m.alertRepo),
		event.WithStrategyMetricJobMetricInitRepo(m.metricInitRepo),
		event.WithStrategyMetricJobSpec(server.CronSpecEvery(m.evaluateInterval)),
		event.WithStrategyMetricJobTimeout(m.evaluateTimeout),
	}
	return event.NewStrategyMetricJob(metric.UniqueKey(), opts...)
}
