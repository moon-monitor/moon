package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz"
)

func NewEvaluateService(
	metricBiz *biz.Metric,
	configBiz *biz.Config,
	logger log.Logger,
) *EvaluateService {
	return &EvaluateService{
		helper:    log.NewHelper(log.With(logger, "module", "service.evaluate")),
		metricBiz: metricBiz,
		configBiz: configBiz,
	}
}

type EvaluateService struct {
	helper    *log.Helper
	metricBiz *biz.Metric
	configBiz *biz.Config
}

func (s *EvaluateService) EvaluateMetric(metricID string) error {
	return nil
}
