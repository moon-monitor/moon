package do

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type TeamMember interface {
	GetTeamID() uint32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetTeamMemberID() uint32
	GetUserID() uint32
	GetInviterID() uint32
	GetPosition() vobj.Role
	GetStatus() vobj.MemberStatus
	GetRoles() []TeamRole
}
