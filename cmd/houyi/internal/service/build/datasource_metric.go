package build

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/merr"
)

func ToMetricDatasourceConfig(metricItem *common.MetricDatasourceItem) (*do.DatasourceMetricConfig, error) {
	switch metricItem.GetDriver() {
	case common.MetricDatasourceItem_Driver_PROMETHEUS:
		prometheusConfig := metricItem.GetPrometheus()
		return ToMetricDatasourceConfigWithPrometheus(prometheusConfig)
	case common.MetricDatasourceItem_Driver_VICTORIA_METRICS:
		victoriaMetricsConfig := metricItem.GetVictoriaMetrics()
		return ToMetricDatasourceConfigWithVictoriaMetrics(victoriaMetricsConfig)
	default:
		return nil, merr.ErrorParamsError("invalid metric datasource driver: %s", metricItem.GetDriver())
	}
}

func ToMetricDatasourceConfigWithPrometheus(prometheusConfig *common.MetricDatasourceItem_Prometheus) (*do.DatasourceMetricConfig, error) {
	if prometheusConfig == nil {
		return nil, merr.ErrorParamsError("prometheus config is nil")
	}
	return &do.DatasourceMetricConfig{
		Driver:    common.MetricDatasourceItem_Driver_PROMETHEUS,
		Endpoint:  prometheusConfig.GetEndpoint(),
		Headers:   prometheusConfig.GetHeaders(),
		Method:    prometheusConfig.GetMethod(),
		CA:        prometheusConfig.GetCa(),
		BasicAuth: ToBasicAuth(prometheusConfig.GetBasicAuth()),
		TLS:       ToTLS(prometheusConfig.GetTls()),
		Enable:    prometheusConfig.GetEnable(),
	}, nil
}

func ToMetricDatasourceConfigWithVictoriaMetrics(victoriaMetricsConfig *common.MetricDatasourceItem_VictoriaMetrics) (*do.DatasourceMetricConfig, error) {
	if victoriaMetricsConfig == nil {
		return nil, merr.ErrorParamsError("victoria metrics config is nil")
	}
	return &do.DatasourceMetricConfig{
		Driver:    common.MetricDatasourceItem_Driver_VICTORIA_METRICS,
		Endpoint:  victoriaMetricsConfig.GetEndpoint(),
		Headers:   victoriaMetricsConfig.GetHeaders(),
		Method:    victoriaMetricsConfig.GetMethod(),
		CA:        victoriaMetricsConfig.GetCa(),
		BasicAuth: ToBasicAuth(victoriaMetricsConfig.GetBasicAuth()),
		TLS:       ToTLS(victoriaMetricsConfig.GetTls()),
		Enable:    victoriaMetricsConfig.GetEnable(),
	}, nil
}
