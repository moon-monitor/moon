package vobj

// MetricDatasourceDriver metric datasource driver
//
//go:generate stringer -type=MetricDatasourceDriver -linecomment -output=metric_datasource_driver.string.go
type MetricDatasourceDriver int8

const (
	MetricDatasourceDriverPrometheus MetricDatasourceDriver = iota + 1
	MetricDatasourceDriverVictoriaMetrics
)
