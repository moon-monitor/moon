package do

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Role interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetRoleID() uint32
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetUsers() []User
	GetMenus() []Menu
}

type TeamRole interface {
	GetTeamID() uint32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetRoleID() uint32
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetMembers() []TeamMember
	GetMenus() []Menu
}
