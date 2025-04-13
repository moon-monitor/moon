package vobj

// ConditionMetric represents the condition metric of a palace object.
//
//go:generate stringer -type=ConditionMetric -linecomment -output condition_metric.string.go
type ConditionMetric int8

const (
	ConditionMetricEQ    ConditionMetric = iota // EQ
	ConditionMetricNE                           // NE
	ConditionMetricGT                           // GT
	ConditionMetricGTE                          // GTE
	ConditionMetricLT                           // LT
	ConditionMetricLTE                          // LTE
	ConditionMetricIn                           // In
	ConditionMetricNotIn                        // NotIn
)
