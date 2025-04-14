package team

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

var _ do.TeamRole = (*Role)(nil)

const tableNameTeamRole = "team_roles"

type Role struct {
	do.TeamModel
	Name    string            `gorm:"column:name;type:varchar(64);not null;comment:角色名" json:"name"`
	Remark  string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Members []*Member         `gorm:"many2many:team_member_roles" json:"members"`
	Menus   []*Menu           `gorm:"many2many:team_role_menus" json:"menus"`
}

func (u *Role) GetTeamID() uint32 {
	if u == nil {
		return 0
	}
	return u.TeamID
}

func (u *Role) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *Role) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *Role) GetRoleID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *Role) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *Role) GetRemark() string {
	if u == nil {
		return ""
	}
	return u.Remark
}

func (u *Role) GetStatus() vobj.GlobalStatus {
	if u == nil {
		return vobj.GlobalStatusUnknown
	}
	return u.Status
}

func (u *Role) GetMembers() []do.TeamMember {
	if u == nil {
		return nil
	}
	return slices.Map(u.Members, func(m *Member) do.TeamMember { return m })
}

func (u *Role) GetMenus() []do.Menu {
	if u == nil {
		return nil
	}
	return slices.Map(u.Menus, func(m *Menu) do.Menu { return m })
}

func (u *Role) TableName() string {
	return tableNameTeamRole
}
