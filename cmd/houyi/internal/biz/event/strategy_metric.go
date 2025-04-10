package event

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/robfig/cron/v3"

	"github.com/moon-monitor/moon/pkg/plugin/server"
)

var _ StrategyJob = (*strategyMetricJob)(nil)

func NewStrategyMetricJob(key string, opts ...StrategyMetricJobOption) (StrategyJob, error) {
	s := &strategyMetricJob{
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
	return func(s *strategyMetricJob) error {
		s.helper = log.NewHelper(log.With(logger, "module", "event.strategy.metric", "jobKey", s.key))
		return nil
	}
}

type strategyMetricJob struct {
	helper *log.Helper
	key    string
	id     cron.EntryID
	spec   server.CronSpec

	metric bo.MetricRule
}

type StrategyMetricJobOption func(*strategyMetricJob) error

func (s *strategyMetricJob) Run() {
	//TODO implement me
	panic("implement me")
}

func (s *strategyMetricJob) ID() cron.EntryID {
	return s.id
}

func (s *strategyMetricJob) Index() string {
	return s.key
}

func (s *strategyMetricJob) Spec() server.CronSpec {
	return s.spec
}

func (s *strategyMetricJob) WithID(id cron.EntryID) server.CronJob {
	s.id = id
	return s
}

func (s *strategyMetricJob) GetEnable() bool {
	return s.metric.GetEnable()
}
