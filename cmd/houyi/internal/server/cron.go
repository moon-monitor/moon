package server

import (
	"github.com/go-kratos/kratos/v2/log"
	
	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

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
