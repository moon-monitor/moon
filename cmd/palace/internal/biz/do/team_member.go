package do

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type TeamMember interface {
	TeamBase
	GetTeamMemberID() uint32
	GetUserID() uint32
	GetInviterID() uint32
	GetPosition() vobj.Role
	GetStatus() vobj.MemberStatus
	GetRoles() []TeamRole
}
