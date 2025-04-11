package event

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/server"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewStrategyMetricJob(key string, opts ...StrategyMetricJobOption) (bo.StrategyJob, error) {
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
	checkOpts := []*checkItem{
		{"configRepo", s.configRepo},
		{"metricInitRepo", s.metricInitRepo},
		{"judgeRepo", s.judgeRepo},
		{"alertRepo", s.alertRepo},
		{"helper", s.helper},
		{"spec", s.spec},
		{"eventBusRepo", s.eventBusRepo},
	}
	return s, checkList(checkOpts...)
}

func WithStrategyMetricJobHelper(logger log.Logger) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if logger == nil {
			return merr.ErrorInternalServerError("logger is nil")
		}
		s.helper = log.NewHelper(log.With(logger, "module", "event.strategy.metric", "jobKey", s.key))
		s.logger = logger
		return nil
	}
}

func WithStrategyMetricJobMetric(metricStrategyUniqueKey string, metricStrategyEnable bool) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if metricStrategyUniqueKey == "" {
			return merr.ErrorInternalServerError("metric strategy unique key is null")
		}
		s.metricStrategyUniqueKey = metricStrategyUniqueKey
		s.metricStrategyEnable = metricStrategyEnable
		return nil
	}
}

func WithStrategyMetricJobConfigRepo(configRepo repository.Config) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if configRepo == nil {
			return merr.ErrorInternalServerError("configRepo is nil")
		}
		s.configRepo = configRepo
		return nil
	}
}

func WithStrategyMetricJobMetricInitRepo(metricInitRepo repository.MetricInit) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if metricInitRepo == nil {
			return merr.ErrorInternalServerError("metricInitRepo is nil")
		}
		s.metricInitRepo = metricInitRepo
		return nil
	}
}

func WithStrategyMetricJobJudgeRepo(judgeRepo repository.Judge) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if judgeRepo == nil {
			return merr.ErrorInternalServerError("judgeRepo is nil")
		}
		s.judgeRepo = judgeRepo
		return nil
	}
}

func WithStrategyMetricJobAlertRepo(alertRepo repository.Alert) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if alertRepo == nil {
			return merr.ErrorInternalServerError("alertRepo is nil")
		}
		s.alertRepo = alertRepo
		return nil
	}
}

func WithStrategyMetricJobSpec(spec server.CronSpec) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if spec == "" {
			return merr.ErrorInternalServerError("spec is empty")
		}
		s.spec = &spec
		return nil
	}
}

func WithStrategyMetricJobTimeout(timeout time.Duration) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if timeout == 0 {
			return merr.ErrorInternalServerError("timeout is 0")
		}
		s.timeout = timeout
		return nil
	}
}

func WithStrategyMetricJobEventBusRepo(eventBusRepo repository.EventBus) StrategyMetricJobOption {
	return func(s *strategyMetricJob) error {
		if eventBusRepo == nil {
			return merr.ErrorInternalServerError("eventBusRepo is nil")
		}
		s.eventBusRepo = eventBusRepo
		return nil
	}
}

type strategyMetricJob struct {
	logger log.Logger
	helper *log.Helper
	key    string
	id     cron.EntryID
	spec   *server.CronSpec

	metricStrategyUniqueKey string
	metricStrategyEnable    bool
	timeout                 time.Duration

	configRepo     repository.Config
	metricInitRepo repository.MetricInit
	judgeRepo      repository.Judge
	alertRepo      repository.Alert
	eventBusRepo   repository.EventBus
}

type StrategyMetricJobOption func(*strategyMetricJob) error

type checkItem struct {
	name  string
	value interface{}
}

func checkList(list ...*checkItem) error {
	for _, listItem := range list {
		if listItem.value == nil {
			return merr.ErrorInternalServerError("%s is nil", listItem.name)
		}
	}
	return nil
}

func (s *strategyMetricJob) Timeout() time.Duration {
	if s.timeout == 0 {
		s.timeout = time.Second * 5
	}
	return s.timeout
}

func (s *strategyMetricJob) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout())
	defer cancel()
	metricStrategy, ok := s.configRepo.GetMetricRule(ctx, s.metricStrategyUniqueKey)
	if !ok {
		s.helper.Warnw("metric strategy not found")
		return
	}
	datasourceConfig, ok := s.configRepo.GetMetricDatasourceConfig(ctx, metricStrategy.GetDatasource())
	if !ok {
		s.helper.Warnw("msg", "datasource config not found")
		return
	}
	query, err := s.metricInitRepo.Init(datasourceConfig)
	if err != nil {
		s.helper.Warnw("msg", "init metric fail", "err", err)
		return
	}

	end := time.Now()
	start := end.Add(-metricStrategy.GetDuration())
	queryRange, err := query.QueryRange(ctx, metricStrategy.GetExpr(), start, end)
	if err != nil {
		s.helper.Warnw("msg", "query fail", "err", err)
		return
	}
	metricJudgeData := slices.Map(queryRange, func(dataItem *do.MetricQueryRangeReply) bo.MetricJudgeData {
		return dataItem
	})

	alerts, err := s.judgeRepo.Metric(ctx, metricJudgeData, metricStrategy)
	if err != nil {
		s.helper.Warnw("msg", "judge fail", "err", err)
		return
	}
	if len(alerts) > 0 {
		s.helper.Debugw("msg", "judge success", "alerts", len(alerts))
	}
	if err := s.alertRepo.Save(ctx, alerts...); err != nil {
		s.helper.Warnw("msg", "alert fail", "err", err)
		return
	}
	alertEventBus := s.eventBusRepo.InAlertEventBus()
	alertJobOpts := []AlertJobOption{
		WithAlertJobHelper(s.logger),
		WithAlertJobEventBusRepo(s.eventBusRepo),
		WithAlertJobAlertRepo(s.alertRepo),
	}
	for _, alert := range alerts {
		alertJobItem, err := NewAlertJob(alert, alertJobOpts...)
		if err != nil {
			s.helper.Warnw("msg", "create alert job fail", "err", err)
			continue
		}
		alertEventBus <- alertJobItem
	}
}

func (s *strategyMetricJob) ID() cron.EntryID {
	if s == nil {
		return 0
	}
	return s.id
}

func (s *strategyMetricJob) Index() string {
	if s == nil {
		return ""
	}
	return s.key
}

func (s *strategyMetricJob) Spec() server.CronSpec {
	if s == nil || s.spec == nil {
		return server.CronSpecEvery(1 * time.Minute)
	}
	return *s.spec
}

func (s *strategyMetricJob) WithID(id cron.EntryID) server.CronJob {
	s.id = id
	return s
}

func (s *strategyMetricJob) GetEnable() bool {
	if s == nil {
		return false
	}
	return s.metricStrategyEnable
}
