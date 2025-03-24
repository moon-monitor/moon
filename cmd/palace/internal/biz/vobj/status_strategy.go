package vobj

// StrategyStatus strategy status
//
//go:generate stringer -type=StrategyStatus -linecomment -output=status_strategy.string.go
type StrategyStatus int8

const (
	StrategyStatusDisabled StrategyStatus = iota // Disabled
	StrategyStatusEnabled                        // Enabled
)
