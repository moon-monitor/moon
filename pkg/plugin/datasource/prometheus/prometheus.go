package prometheus

import (
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
)

type Config interface {
	GetEndpoint() string
	GetBasicAuth() datasource.BasicAuth
}
