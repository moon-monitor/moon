package vobj

// TimeEngineStatus time engine status
//
//go:generate stringer -type=TimeEngineStatus -linecomment -output=status_time_engine.string.go
type TimeEngineStatus int8

const (
	TimeEngineStatusEnabled  TimeEngineStatus = iota // Enabled
	TimeEngineStatusDisabled                         // Disabled
)
