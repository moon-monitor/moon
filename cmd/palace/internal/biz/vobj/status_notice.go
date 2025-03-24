package vobj

// NoticeStatus notice status
//
//go:generate stringer -type=NoticeStatus -linecomment -output=status_notice.string.go
type NoticeStatus int8

const (
	NoticeStatusEnabled  NoticeStatus = iota // Enabled
	NoticeStatusDisabled                     // Disabled
)
