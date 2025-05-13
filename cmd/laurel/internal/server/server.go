package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/laurel/internal/conf"
	"github.com/moon-monitor/moon/cmd/laurel/internal/service"
	"github.com/moon-monitor/moon/pkg/api/common"
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
) server.Servers {
	common.RegisterHealthServer(rpcSrv, healthService)
	common.RegisterHealthHTTPServer(httpSrv, healthService)

	return server.Servers{rpcSrv, httpSrv, tickerSrv}
}
