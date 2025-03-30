package vobj

// MetricStrategyCondition metric strategy sample mode
//
//go:generate stringer -type=MetricStrategyCondition -linecomment -output=metric_strategy_condition.string.go
type MetricStrategyCondition int8

const (
	MetricStrategyConditionEQ    MetricStrategyCondition = iota + 1 // eq
	MetricStrategyConditionNE                                       // ne
	MetricStrategyConditionGT                                       // gt
	MetricStrategyConditionGTE                                      // gte
	MetricStrategyConditionLT                                       // lt
	MetricStrategyConditionLTE                                      // lte
	MetricStrategyConditionIN                                       // in
	MetricStrategyConditionNotIN                                    // not in
)
