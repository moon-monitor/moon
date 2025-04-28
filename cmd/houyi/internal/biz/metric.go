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
	cacheRepo repository.Cache,
	logger log.Logger,
) *Metric {
	evaluateConf := bc.GetEvaluate()
	syncConfig := bc.GetConfig()
	return &Metric{
		helper:           log.NewHelper(log.With(logger, "module", "biz.metric")),
		judgeRepo:        judgeRepo,
		alertRepo:        alertRepo,
		metricInitRepo:   metricInitRepo,
		configRepo:       configRepo,
		eventBusRepo:     eventBusRepo,
		cacheRepo:        cacheRepo,
		evaluateInterval: evaluateConf.GetInterval().AsDuration(),
		evaluateTimeout:  evaluateConf.GetTimeout().AsDuration(),
		syncInterval:     syncConfig.GetSyncInterval().AsDuration(),
		syncTimeout:      syncConfig.GetSyncTimeout().AsDuration(),
	}
}

type Metric struct {
	helper *log.Helper

	judgeRepo        repository.Judge
	alertRepo        repository.Alert
	metricInitRepo   repository.MetricInit
	configRepo       repository.Config
	eventBusRepo     repository.EventBus
	cacheRepo        repository.Cache
	evaluateInterval time.Duration
	evaluateTimeout  time.Duration
	syncInterval     time.Duration
	syncTimeout      time.Duration
}

func (m *Metric) Loads() []*server.TickTask {
	return []*server.TickTask{
		{
			Fn:        m.syncMetricRuleConfigs,
			Name:      "syncMetricRuleConfigs",
			Timeout:   m.syncTimeout,
			Interval:  m.syncInterval,
			Immediate: true,
		},
	}
}

func (m *Metric) syncMetricRuleConfigs(ctx context.Context, isStop bool) error {
	if isStop {
		return nil
	}
	metricRules, err := m.configRepo.GetMetricRules(ctx)
	if err != nil {
		m.helper.WithContext(ctx).Errorw("method", "syncMetricRuleConfigs", "err", err)
		return err
	}

	return m.syncMetricJob(ctx, metricRules...)
}

func (m *Metric) syncMetricJob(ctx context.Context, rules ...bo.MetricRule) error {
	inStrategyJobEventBus := m.eventBusRepo.InStrategyJobEventBus()
	for _, rule := range rules {
		rule.Renovate()
		strategyJob, err := m.newStrategyJob(ctx, rule)
		if err != nil {
			m.helper.WithContext(ctx).Warnw("msg", "new strategy job error", "err", err)
			continue
		}
		inStrategyJobEventBus <- strategyJob
	}

	m.helper.WithContext(ctx).Debug("save metric rules success")
	return nil
}

func (m *Metric) SaveMetricRules(ctx context.Context, rules ...bo.MetricRule) error {
	if len(rules) == 0 {
		return nil
	}

	if err := m.configRepo.SetMetricRules(ctx, rules...); err != nil {
		m.helper.WithContext(ctx).Errorw("msg", "save metric rules error", "err", err)
		return err
	}

	return m.syncMetricJob(ctx, rules...)
}

func (m *Metric) newStrategyJob(_ context.Context, metric bo.MetricRule) (bo.StrategyJob, error) {
	opts := []event.StrategyMetricJobOption{
		event.WithStrategyMetricJobHelper(m.helper.Logger()),
		event.WithStrategyMetricJobMetric(metric.UniqueKey(), metric.GetEnable()),
		event.WithStrategyMetricJobConfigRepo(m.configRepo),
		event.WithStrategyMetricJobJudgeRepo(m.judgeRepo),
		event.WithStrategyMetricJobAlertRepo(m.alertRepo),
		event.WithStrategyMetricJobMetricInitRepo(m.metricInitRepo),
		event.WithStrategyMetricJobSpec(m.evaluateInterval),
		event.WithStrategyMetricJobTimeout(m.evaluateTimeout),
		event.WithStrategyMetricJobEventBusRepo(m.eventBusRepo),
		event.WithStrategyMetricJobCacheRepo(m.cacheRepo),
	}
	return event.NewStrategyMetricJob(metric.UniqueKey(), opts...)
}
