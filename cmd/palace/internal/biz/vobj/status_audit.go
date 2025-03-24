package vobj

// StatusAudit status audit
//
//go:generate stringer -type=StatusAudit -linecomment -output=status_audit.string.go
type StatusAudit int8

const (
	AuditStatusPending  StatusAudit = iota // Pending
	AuditStatusApproved StatusAudit = iota // Approved
	AuditStatusRejected StatusAudit = iota // Rejected
)
