package service

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
)

func NewEventBusService(
	eventBus *biz.EventBus,
	logger log.Logger,
) *EventBusService {
	return &EventBusService{
		helper:   log.NewHelper(log.With(logger, "module", "service.event-bus")),
		eventBus: eventBus,
	}
}

type EventBusService struct {
	helper   *log.Helper
	eventBus *biz.EventBus
}

func (s *EventBusService) OutStrategyJobEventBus() <-chan bo.StrategyJob {
	return s.eventBus.OutStrategyJobEventBus()
}
