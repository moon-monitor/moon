package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/service"
	commonv1 "github.com/moon-monitor/moon/pkg/api/common"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
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
	authService *service.AuthService,
	serverService *service.ServerService,
	resourceService *service.ResourceService,
) server.Servers {
	commonv1.RegisterHealthServer(rpcSrv, healthService)
	commonv1.RegisterServerServer(rpcSrv, serverService)

	commonv1.RegisterHealthHTTPServer(httpSrv, healthService)
	commonv1.RegisterServerHTTPServer(httpSrv, serverService)
	palacev1.RegisterAuthHTTPServer(httpSrv, authService)
	palacev1.RegisterResourceHTTPServer(httpSrv, resourceService)
	return server.Servers{rpcSrv, httpSrv}
}
