package event

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

func NewStrategyMetricJob(key string, opts ...StrategyMetricJobOption) (*StrategyMetricJob, error) {
	s := &StrategyMetricJob{
		key: key,
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	if s.helper == nil {
		WithStrategyMetricJobHelper(log.GetLogger())
	}
	return s, nil
}

func WithStrategyMetricJobHelper(logger log.Logger) StrategyMetricJobOption {
	return func(s *StrategyMetricJob) error {
		if logger == nil {
			return merr.ErrorInternalServerError("logger is nil")
		}
		s.helper = log.NewHelper(log.With(logger, "module", "event.strategy.metric", "jobKey", s.key))
		return nil
	}
}

func WithStrategyMetricJobMetric(metric bo.MetricRule) StrategyMetricJobOption {
	return func(s *StrategyMetricJob) error {
		if metric == nil {
			return merr.ErrorInternalServerError("metric is nil")
		}
		s.metric = metric
		return nil
	}
}

func WithStrategyConfigRepo(configRepo repository.Config) StrategyMetricJobOption {
	return func(s *StrategyMetricJob) error {
		if configRepo == nil {
			return merr.ErrorInternalServerError("configRepo is nil")
		}
		s.configRepo = configRepo
		return nil
	}
}

type StrategyMetricJob struct {
	helper *log.Helper
	key    string
	id     cron.EntryID
	spec   server.CronSpec

	metric bo.MetricRule

	configRepo     repository.Config
	metricInitRepo repository.MetricInit
	judgeRepo      repository.Judge
	alertRepo      repository.Alert
}

type StrategyMetricJobOption func(*StrategyMetricJob) error

func (s *StrategyMetricJob) Run() {
	//TODO implement me
	panic("implement me")
}

func (s *StrategyMetricJob) ID() cron.EntryID {
	return s.id
}

func (s *StrategyMetricJob) Index() string {
	return s.key
}

func (s *StrategyMetricJob) Spec() server.CronSpec {
	return s.spec
}

func (s *StrategyMetricJob) WithID(id cron.EntryID) server.CronJob {
	s.id = id
	return s
}

func (s *StrategyMetricJob) GetEnable() bool {
	return s.metric.GetEnable()
}
