package vobj

// RoleStatus role status
//
//go:generate stringer -type=RoleStatus -linecomment -output=status_role.string.go
type RoleStatus int8

const (
	RoleStatusDisabled RoleStatus = iota
	RoleStatusNormal
	RoleStatusForbidden
)
