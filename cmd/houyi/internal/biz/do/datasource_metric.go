package do

import (
	"encoding/json"
	"fmt"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/util/password"
)

var _ cache.Object = (*DatasourceMetricConfig)(nil)

type DatasourceMetricConfig struct {
	Driver    common.MetricDatasourceItem_Driver `json:"driver"`
	Endpoint  string                             `json:"endpoint"`
	Headers   map[string]string                  `json:"headers"`
	Method    common.DatasourceQueryMethod       `json:"method"`
	CA        string                             `json:"ca"`
	BasicAuth *BasicAuth                         `json:"basicAuth"`
	TLS       *TLS                               `json:"tls"`
	Enable    bool                               `json:"enable"`
}

func (d *DatasourceMetricConfig) GetEnable() bool {
	if d == nil {
		return false
	}
	return d.Enable
}

func (d *DatasourceMetricConfig) GetDriver() common.MetricDatasourceItem_Driver {
	if d == nil {
		return common.MetricDatasourceItem_Driver_UNKNOWN
	}
	return d.Driver
}

func (d *DatasourceMetricConfig) GetEndpoint() string {
	if d == nil {
		return ""
	}
	return d.Endpoint
}

func (d *DatasourceMetricConfig) GetHeaders() map[string]string {
	if d == nil {
		return nil
	}
	return d.Headers
}

func (d *DatasourceMetricConfig) GetMethod() common.DatasourceQueryMethod {
	if d == nil {
		return common.DatasourceQueryMethod_QueryMethod_HTTP_POST
	}
	return d.Method
}

func (d *DatasourceMetricConfig) GetBasicAuth() datasource.BasicAuth {
	if d == nil {
		return nil
	}
	return d.BasicAuth
}

func (d *DatasourceMetricConfig) GetTLS() datasource.TLS {
	if d == nil {
		return nil
	}
	return d.TLS
}

func (d *DatasourceMetricConfig) GetCA() string {
	if d == nil {
		return ""
	}
	return d.CA
}

func (d *DatasourceMetricConfig) MarshalBinary() (data []byte, err error) {
	return json.Marshal(d)
}

func (d *DatasourceMetricConfig) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *DatasourceMetricConfig) UniqueKey() string {
	return fmt.Sprintf("%d:%s", d.Driver, password.MD5(d.Endpoint))
}
