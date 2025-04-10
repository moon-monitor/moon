package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

var _ transport.Server = (*CronServer)(nil)

func NewCronServer(evaluateService *service.EvaluateService, logger log.Logger) *CronServer {
	return &CronServer{
		evaluateService: evaluateService,
		logger:          logger,
		helper:          log.NewHelper(log.With(logger, "module", "server.cron")),
		CronJobServer:   server.NewCronJobServer(logger),
	}
}

type CronServer struct {
	evaluateService *service.EvaluateService
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
		for metricID := range c.evaluateService.EventBus() {
			c.call(metricID)
		}
	}()
	return c.CronJobServer.Start(ctx)
}

func (c *CronServer) Stop(ctx context.Context) error {
	return c.CronJobServer.Stop(ctx)
}

func (c *CronServer) call(metricID string) {
	c.helper.Info("cron job start", "metricID", metricID)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.evaluateService.EvaluateMetric(ctx, metricID); err != nil {
		c.helper.Warnw("msg", "cron job error", "metricID", metricID, "err", err)
	}
}
