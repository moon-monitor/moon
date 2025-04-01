package vobj

//go:generate stringer -type=AuditAction -linecomment -output=action_audit.string.go
type AuditAction int8

const (
	AuditActionJoin  AuditAction = iota // Join
	AuditActionLeave                    // Leave
)
