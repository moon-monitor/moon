package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz"
)

func NewEvaluateService(
	metricBiz *biz.Metric,
	configBiz *biz.Config,
	eventBus *biz.EventBus,
	logger log.Logger,
) *EvaluateService {
	return &EvaluateService{
		helper:    log.NewHelper(log.With(logger, "module", "service.evaluate")),
		metricBiz: metricBiz,
		configBiz: configBiz,
		eventBus:  eventBus,
	}
}

type EvaluateService struct {
	helper    *log.Helper
	metricBiz *biz.Metric
	configBiz *biz.Config
	eventBus  *biz.EventBus
}

func (s *EvaluateService) EvaluateMetric(ctx context.Context, metricID string) error {
	metricRule, ok := s.configBiz.GetMetricRule(ctx, metricID)
	if !ok {
		return nil
	}
	return s.metricBiz.Evaluate(ctx, metricRule)
}

func (s *EvaluateService) EventBus() <-chan string {
	return s.eventBus.OutMetricIDEventBus()
}
