package service

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
)

func NewEvaluateService(
	eventBus *biz.EventBus,
	logger log.Logger,
) *EvaluateService {
	return &EvaluateService{
		helper:   log.NewHelper(log.With(logger, "module", "service.evaluate")),
		eventBus: eventBus,
	}
}

type EvaluateService struct {
	helper   *log.Helper
	eventBus *biz.EventBus
}

func (s *EvaluateService) EventBus() <-chan event.StrategyJob {
	return s.eventBus.OutStrategyJobEventBus()
}
