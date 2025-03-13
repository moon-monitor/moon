package server

import (
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap) *http.Server {
	httpConf := c.GetHttp()
	opts := []http.ServerOption{
		http.Middleware(),
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

	registerDocs(c, srv)

	return srv
}

func registerDocs(c *conf.Bootstrap, srv *http.Server) {
	if !c.IsDev() {
		return
	}
	srv.HandlePrefix("/doc/", nethttp.StripPrefix("/doc/", nethttp.FileServer(nethttp.Dir("./swagger"))))
}
