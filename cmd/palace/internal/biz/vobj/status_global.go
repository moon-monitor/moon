package vobj

// GlobalStatus is a global status value.
//
//go:generate stringer -type=GlobalStatus -linecomment -output=status_global.string.go
type GlobalStatus int8

const (
	GlobalStatusDisable GlobalStatus = iota
	GlobalStatusEnable
)
