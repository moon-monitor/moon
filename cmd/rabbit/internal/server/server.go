package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/conf"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/service"
	commonv1 "github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

// ProviderSetServer is server providers.
var ProviderSetServer = wire.NewSet(NewGRPCServer, NewHTTPServer, RegisterService)

// RegisterService register service
func RegisterService(
	c *conf.Bootstrap,
	rpcSrv *grpc.Server,
	httpSrv *http.Server,
	healthService *service.HealthService,
) *server.Server {
	commonv1.RegisterHealthServer(rpcSrv, healthService)
	commonv1.RegisterHealthHTTPServer(httpSrv, healthService)
	return &server.Server{RpcSrv: rpcSrv, HttpSrv: httpSrv}
}
