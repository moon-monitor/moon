package vobj

// StrategyGroupStatus strategy group status
//
//go:generate stringer -type=StrategyGroupStatus -linecomment -output=status_strategy_group.string.go
type StrategyGroupStatus int8

const (
	StrategyGroupStatusDisabled      StrategyGroupStatus = iota // Disabled
	StrategyGroupStatusEnabled                                  // Enabled
	StrategyGroupStatusImporting                                // Importing
	StrategyGroupStatusImportFailed                             // ImportFailed
	StrategyGroupStatusImportSuccess                            // ImportSuccess
)
