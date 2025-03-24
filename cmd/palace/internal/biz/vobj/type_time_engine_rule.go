package vobj

// TimeEngineRuleType time engine rule type
//
//go:generate stringer -type=TimeEngineRuleType -linecomment -output=type_time_engine_rule.string.go
type TimeEngineRuleType int8

const (
	TimeEngineRuleTypeHourRange   TimeEngineRuleType = iota // hour
	TimeEngineRuleTypeDaysOfWeek                            // week
	TimeEngineRuleTypeDaysOfMonth                           // day
	TimeEngineRuleTypeMonths                                // month
)
