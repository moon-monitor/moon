package victoria

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
)

type Config interface {
	GetEndpoint() string
	GetHeaders() map[string]string
	GetMethod() common.DatasourceQueryMethod
	GetBasicAuth() datasource.BasicAuth
	GetTLS() datasource.TLS
	GetCA() string
}

func New(c Config, logger log.Logger) *Victoria {
	return &Victoria{
		c:      c,
		helper: log.NewHelper(log.With(logger, "module", "plugin.datasource.victoria")),
	}
}

type Victoria struct {
	c      Config
	helper *log.Helper
}
