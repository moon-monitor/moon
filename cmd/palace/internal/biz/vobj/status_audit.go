package vobj

// StatusAudit status audit
//
//go:generate stringer -type=StatusAudit -linecomment -output=status_audit.string.go
type StatusAudit int8

const (
	AuditStatusUnknown  StatusAudit = iota // Unknown
	AuditStatusPending                     // Pending
	AuditStatusApproved                    // Approved
	AuditStatusRejected                    // Rejected
)
