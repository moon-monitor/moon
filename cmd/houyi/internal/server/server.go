package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/houyi/internal/conf"
	"github.com/moon-monitor/moon/cmd/houyi/internal/service"
	commonv1 "github.com/moon-monitor/moon/pkg/api/common"
	houyiv1 "github.com/moon-monitor/moon/pkg/api/houyi/v1"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

// ProviderSetServer is server providers.
var ProviderSetServer = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	RegisterService,
	NewCronStrategyJobServer,
	NewCronAlertJobServer,
)

// RegisterService register service
func RegisterService(
	c *conf.Bootstrap,
	rpcSrv *grpc.Server,
	httpSrv *http.Server,
	cronStrategySrv *CronStrategyJobServer,
	cronAlertSrv *CronAlertJobServer,
	healthService *service.HealthService,
	syncService *service.SyncService,
	alertService *service.AlertService,
) server.Servers {
	commonv1.RegisterHealthServer(rpcSrv, healthService)
	commonv1.RegisterHealthHTTPServer(httpSrv, healthService)
	houyiv1.RegisterSyncServer(rpcSrv, syncService)
	houyiv1.RegisterSyncHTTPServer(httpSrv, syncService)
	houyiv1.RegisterAlertHTTPServer(httpSrv, alertService)
	return server.Servers{rpcSrv, httpSrv, cronStrategySrv, cronAlertSrv}
}
