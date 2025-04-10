package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

var _ transport.Server = (*CronServer)(nil)

func NewCronServer(evaluateService *service.EventBusService, logger log.Logger) *CronServer {
	return &CronServer{
		evaluateService: evaluateService,
		logger:          logger,
		helper:          log.NewHelper(log.With(logger, "module", "server.cron")),
		CronJobServer:   server.NewCronJobServer(logger),
	}
}

type CronServer struct {
	evaluateService *service.EventBusService
	logger          log.Logger
	helper          *log.Helper
	*server.CronJobServer
}

func (c *CronServer) Start(ctx context.Context) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				c.helper.Errorw("method", "watchEventBus", "panic", err)
			}
		}()
		for strategyJob := range c.evaluateService.OutStrategyJobEventBus() {
			if strategyJob.GetEnable() {
				c.AddJob(strategyJob)
			} else {
				c.RemoveJob(strategyJob)
			}
		}
	}()
	return c.CronJobServer.Start(ctx)
}

func (c *CronServer) Stop(ctx context.Context) error {
	return c.CronJobServer.Stop(ctx)
}
