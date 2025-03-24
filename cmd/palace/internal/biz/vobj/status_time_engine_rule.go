package vobj

// TimeEngineRuleStatus time engine rule status
//
//go:generate stringer -type=TimeEngineRuleStatus -linecomment -output=status_time_engine_rule.string.go
type TimeEngineRuleStatus int8

const (
	TimeEngineRuleStatusEnabled  TimeEngineRuleStatus = iota // Enabled
	TimeEngineRuleStatusDisabled                             // Disabled
)
