package bo

import (
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
)

type MetricDatasourceConfig interface {
	cache.Object
	GetDriver() common.MetricDatasourceItem_Driver
	GetEndpoint() string
	GetHeaders() map[string]string
	GetMethod() common.DatasourceQueryMethod
	GetBasicAuth() datasource.BasicAuth
	GetTLS() datasource.TLS
	GetCA() string
	GetEnable() bool
}
