package biz

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
)

func NewEventBus(eventBus repository.EventBus, logger log.Logger) *EventBus {
	return &EventBus{
		helper:   log.NewHelper(log.With(logger, "module", "biz.event-bus")),
		eventBus: eventBus,
	}
}

type EventBus struct {
	helper *log.Helper

	eventBus repository.EventBus
}

func (e *EventBus) InMetricIDEventBus() chan<- string {
	return e.eventBus.InMetricIDEventBus()
}

func (e *EventBus) OutMetricIDEventBus() <-chan string {
	return e.eventBus.OutMetricIDEventBus()
}
