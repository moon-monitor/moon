package build

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToMetricQueryResults(results []*bo.MetricQueryResult) []*common.MetricQueryResult {
	return slices.Map(results, ToMetricQueryResult)
}

func ToMetricQueryResult(result *bo.MetricQueryResult) *common.MetricQueryResult {
	return &common.MetricQueryResult{
		Metric: result.Metric,
		Values: ToMetricQueryResultValues(result.Values),
		Value:  ToMetricQueryResultValue(result.Value),
	}
}

func ToMetricQueryResultValue(value *bo.MetricQueryValue) *common.MetricQueryResultValue {
	return &common.MetricQueryResultValue{
		Timestamp: value.Timestamp,
		Value:     value.Value,
	}
}

func ToMetricQueryResultValues(values []*bo.MetricQueryValue) []*common.MetricQueryResultValue {
	return slices.Map(values, ToMetricQueryResultValue)
}
