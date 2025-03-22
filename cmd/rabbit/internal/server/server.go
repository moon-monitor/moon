package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/conf"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/service"
	commonv1 "github.com/moon-monitor/moon/pkg/api/common"
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

// ProviderSetServer is server providers.
var ProviderSetServer = wire.NewSet(NewGRPCServer, NewHTTPServer, NewTicker, RegisterService)

// RegisterService register service
func RegisterService(
	c *conf.Bootstrap,
	rpcSrv *grpc.Server,
	httpSrv *http.Server,
	tickerSrv *server.Ticker,
	healthService *service.HealthService,
	sendService *service.SendService,
	syncService *service.SyncService,
) server.Servers {
	commonv1.RegisterHealthServer(rpcSrv, healthService)
	commonv1.RegisterHealthHTTPServer(httpSrv, healthService)
	rabbitv1.RegisterSendServer(rpcSrv, sendService)
	rabbitv1.RegisterSyncServer(rpcSrv, syncService)
	rabbitv1.RegisterSendHTTPServer(httpSrv, sendService)
	rabbitv1.RegisterSyncHTTPServer(httpSrv, syncService)
	return server.Servers{rpcSrv, httpSrv, tickerSrv}
}
