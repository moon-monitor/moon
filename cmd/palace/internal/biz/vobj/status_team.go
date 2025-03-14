package vobj

// TeamStatus user status
//
//go:generate stringer -type=TeamStatus -linecomment -output=status_team.string.go
type TeamStatus int8

const (
	TeamStatusUnknown   TeamStatus = iota // unknown
	TeamStatusNormal                      // normal
	TeamStatusForbidden                   // forbidden
	TeamStatusDeleted                     // deleted
)
