package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
)

func NewEventBusServer(
	eventBusService *service.EventBusService,
	logger log.Logger,
) *EventBusServer {
	return &EventBusServer{
		helper:          log.NewHelper(log.With(logger, "module", "server.event-bus")),
		stop:            make(chan struct{}),
		eventBusService: eventBusService,
	}
}

type EventBusServer struct {
	helper *log.Helper
	stop   chan struct{}

	eventBusService *service.EventBusService
}

func (e *EventBusServer) Start(ctx context.Context) error {
	defer e.helper.Info("[EventBus] server is started")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				e.helper.Errorw("method", "watchEventBus", "panic", err)
			}
		}()
	}()

	return nil
}

func (e *EventBusServer) Stop(_ context.Context) error {
	defer e.helper.Info("[EventBus] server is stopped")
	close(e.stop)
	return nil
}
