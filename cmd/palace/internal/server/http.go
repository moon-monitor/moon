package server

import (
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bc *conf.Bootstrap) *http.Server {
	serverConf := bc.GetServer()
	httpConf := serverConf.GetHttp()
	authConf := bc.GetAuth()
	opts := []http.ServerOption{
		http.Middleware(
			middleware.JwtServer(authConf.GetJwt().GetSignKey()),
			middleware.MustLogin(),
			middleware.BindHeaders(),
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

	registerDocs(bc, srv)

	return srv
}

func registerDocs(c *conf.Bootstrap, srv *http.Server) {
	if !c.IsDev() {
		return
	}
	srv.HandlePrefix("/doc/", nethttp.StripPrefix("/doc/", nethttp.FileServer(nethttp.Dir("./swagger"))))
}
