package vobj

// NoticeHookStatus notice hook status
//
//go:generate stringer -type=NoticeHookStatus -linecomment -output=status_notice_hook.string.go
type NoticeHookStatus int8

const (
	NoticeHookStatusEnabled  NoticeHookStatus = iota // Enabled
	NoticeHookStatusDisabled                         // Disabled
)
