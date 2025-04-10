package bo

import (
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
)

type MetricDatasourceConfig interface {
	cache.Object
	GetId() uint32
	GetName() string
	GetDriver() common.MetricDatasourceItem_Driver
	GetEndpoint() string
	GetHeaders() map[string]string
	GetMethod() common.DatasourceQueryMethod
	GetBasicAuth() datasource.BasicAuth
	GetTLS() datasource.TLS
	GetCA() string
	GetEnable() bool
	GetScrapeInterval() time.Duration
}
