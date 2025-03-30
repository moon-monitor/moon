package vobj

// MetricStrategySampleMode metric strategy sample mode
//
//go:generate stringer -type=MetricStrategySampleMode -linecomment -output=metric_strategy_sample_mode.string.go
type MetricStrategySampleMode int8

const (
	MetricStrategySampleModeFor MetricStrategySampleMode = iota + 1 // for
	MetricStrategySampleModeMax                                     // max
	MetricStrategySampleModeMin                                     // min
)
