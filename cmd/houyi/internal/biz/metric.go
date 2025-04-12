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
	eventBusRepo repository.EventBus,
	logger log.Logger,
) (*Metric, error) {
	evaluateConf := bc.GetEvaluate()
	syncConfig := bc.GetConfig()
	m := &Metric{
		logger:           logger,
		helper:           log.NewHelper(log.With(logger, "module", "biz.metric")),
		judgeRepo:        judgeRepo,
		alertRepo:        alertRepo,
		metricInitRepo:   metricInitRepo,
		configRepo:       configRepo,
		eventBusRepo:     eventBusRepo,
		evaluateInterval: evaluateConf.GetInterval().AsDuration(),
		evaluateTimeout:  evaluateConf.GetTimeout().AsDuration(),
		syncInterval:     syncConfig.GetSyncInterval().AsDuration(),
		syncTimeout:      syncConfig.GetSyncTimeout().AsDuration(),
	}
	return m, m.sync()
}

type Metric struct {
	logger log.Logger
	helper *log.Helper

	judgeRepo        repository.Judge
	alertRepo        repository.Alert
	metricInitRepo   repository.MetricInit
	configRepo       repository.Config
	eventBusRepo     repository.EventBus
	evaluateInterval time.Duration
	evaluateTimeout  time.Duration
	syncInterval     time.Duration
	syncTimeout      time.Duration
}

func (m *Metric) sync() error {
	task := &server.TickTask{
		Fn:       m.syncMetricRuleConfigs,
		Interval: m.syncInterval,
		Name:     "syncMetricRuleConfigs",
		Timeout:  m.syncTimeout,
	}
	ticker := server.NewTicker(m.syncInterval, task, server.WithTickerLogger(m.logger), server.WithTickerImmediate(true))
	return ticker.Start(context.Background())
}

func (m *Metric) syncMetricRuleConfigs(ctx context.Context, isStop bool) error {
	if isStop {
		return nil
	}
	metricRules, err := m.configRepo.GetMetricRules(ctx)
	if err != nil {
		m.helper.Errorw("method", "syncMetricRuleConfigs", "err", err)
		return err
	}

	return m.syncMetricJob(ctx, metricRules...)
}

func (m *Metric) syncMetricJob(ctx context.Context, rules ...bo.MetricRule) error {
	inStrategyJobEventBus := m.eventBusRepo.InStrategyJobEventBus()
	for _, rule := range rules {
		strategyJob, err := m.newStrategyJob(ctx, rule)
		if err != nil {
			m.helper.Warnw("msg", "new strategy job error", "err", err)
			continue
		}
		inStrategyJobEventBus <- strategyJob
	}

	m.helper.Debug("save metric rules success")
	return nil
}

func (m *Metric) SaveMetricRules(ctx context.Context, rules ...bo.MetricRule) error {
	if len(rules) == 0 {
		return nil
	}

	if err := m.configRepo.SetMetricRules(ctx, rules...); err != nil {
		m.helper.Errorw("msg", "save metric rules error", "err", err)
		return err
	}

	return m.syncMetricJob(ctx, rules...)
}

func (m *Metric) newStrategyJob(_ context.Context, metric bo.MetricRule) (bo.StrategyJob, error) {
	opts := []event.StrategyMetricJobOption{
		event.WithStrategyMetricJobHelper(m.logger),
		event.WithStrategyMetricJobMetric(metric.UniqueKey(), metric.GetEnable()),
		event.WithStrategyMetricJobConfigRepo(m.configRepo),
		event.WithStrategyMetricJobJudgeRepo(m.judgeRepo),
		event.WithStrategyMetricJobAlertRepo(m.alertRepo),
		event.WithStrategyMetricJobMetricInitRepo(m.metricInitRepo),
		event.WithStrategyMetricJobSpec(server.CronSpecEvery(m.evaluateInterval)),
		event.WithStrategyMetricJobTimeout(m.evaluateTimeout),
		event.WithStrategyMetricJobEventBusRepo(m.eventBusRepo),
	}
	return event.NewStrategyMetricJob(metric.UniqueKey(), opts...)
}
