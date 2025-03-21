package server

import (
	"embed"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/cmd/palace/internal/service"
	"github.com/moon-monitor/moon/pkg/middler"
	"github.com/moon-monitor/moon/pkg/util/docs"
)

//go:embed swagger
var docFS embed.FS

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, authService *service.AuthService, logger log.Logger) *http.Server {
	serverConf := bc.GetServer()
	httpConf := serverConf.GetHttp()
	jwtConf := bc.GetAuth().GetJwt()

	authMiddleware := selector.Server(
		middleware.JwtServer(jwtConf.GetSignKey()),
		middleware.MustLogin(authService.VerifyToken),
		middleware.BindHeaders(),
	).Match(middler.AllowListMatcher(jwtConf.GetAllowOperations()...)).Build()
	opts := []http.ServerOption{
		http.Filter(middler.Cors(httpConf)),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			authMiddleware,
			middler.Validate(),
		),
	}
	if httpConf.GetNetwork() != "" {
		opts = append(opts, http.Network(httpConf.GetNetwork()))
	}
	if httpConf.GetAddr() != "" {
		opts = append(opts, http.Address(httpConf.GetAddr()))
	}
	if httpConf.GetTimeout() != nil {
		opts = append(opts, http.Timeout(httpConf.GetTimeout().AsDuration()))
	}
	srv := http.NewServer(opts...)

	docs.RegisterDocs(srv, docFS, bc.IsDev())
	registerOAuth2(bc.GetAuth().GetOauth2(), srv, authService)

	return srv
}

func registerOAuth2(c *conf.Auth_OAuth2, httpSrv *http.Server, authService *service.AuthService) {
	if !c.GetEnable() {
		return
	}
	auth := httpSrv.Route("/auth")
	list := c.GetConfigs()
	for _, config := range list {
		app := vobj.OAuthAPP(config.GetApp())
		appRoute := auth.Group(strings.ToLower(app.String()))
		appRoute.GET("/", authService.OAuthLogin(app))
		appRoute.GET("/callback", authService.OAuthLoginCallback(app))
	}
}
