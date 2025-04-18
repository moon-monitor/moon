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
	userService *service.UserService,
	dashboardService *service.DashboardService,
	callbackService *service.CallbackService,
	teamDashboardService *service.TeamDashboardService,
	datasourceService *service.TeamDatasourceService,
	dictService *service.TeamDictService,
	noticeService *service.TeamNoticeService,
	strategyService *service.TeamStrategyService,
	teamService *service.TeamService,
	systemService *service.SystemService,
) server.Servers {
	commonv1.RegisterHealthServer(rpcSrv, healthService)
	commonv1.RegisterServerServer(rpcSrv, serverService)

	commonv1.RegisterHealthHTTPServer(httpSrv, healthService)
	commonv1.RegisterServerHTTPServer(httpSrv, serverService)
	palacev1.RegisterAuthHTTPServer(httpSrv, authService)
	palacev1.RegisterResourceHTTPServer(httpSrv, resourceService)
	palacev1.RegisterUserHTTPServer(httpSrv, userService)
	palacev1.RegisterTeamDashboardHTTPServer(httpSrv, dashboardService)
	palacev1.RegisterCallbackHTTPServer(httpSrv, callbackService)
	palacev1.RegisterTeamDashboardHTTPServer(httpSrv, teamDashboardService)
	palacev1.RegisterTeamDatasourceHTTPServer(httpSrv, datasourceService)
	palacev1.RegisterTeamDictHTTPServer(httpSrv, dictService)
	palacev1.RegisterTeamNoticeHTTPServer(httpSrv, noticeService)
	palacev1.RegisterTeamStrategyHTTPServer(httpSrv, strategyService)
	palacev1.RegisterTeamHTTPServer(httpSrv, teamService)
	palacev1.RegisterSystemHTTPServer(httpSrv, systemService)
	return server.Servers{rpcSrv, httpSrv}
}
