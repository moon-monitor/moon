package vobj

// ResourceStatus resource status
//
//go:generate stringer -type=ResourceStatus -linecomment -output=status_resource.string.go
type ResourceStatus int8

const (
	ResourceStatusDisabled ResourceStatus = iota // Disabled
	ResourceStatusEnabled
)
