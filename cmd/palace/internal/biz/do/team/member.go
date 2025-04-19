package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

var _ do.TeamMember = (*Member)(nil)

const tableNameTeamMember = "team_members"

type Member struct {
	do.TeamModel
	MemberName string            `gorm:"column:member_name;type:varchar(64);not null;comment:成员名" json:"memberName"`
	Remark     string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	UserID     uint32            `gorm:"column:user_id;type:int unsigned;not null;comment:用户ID" json:"userID"`
	InviterID  uint32            `gorm:"column:inviter_id;type:int unsigned;not null;comment:邀请者ID" json:"inviterID"`
	Position   vobj.Role         `gorm:"column:position;type:tinyint(2);not null;comment:职位" json:"position"`
	Status     vobj.MemberStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Roles      []*Role           `gorm:"many2many:sys_team_member_roles" json:"roles"`
}

func (u *Member) GetMemberName() string {
	if u == nil {
		return ""
	}
	return u.MemberName
}

func (u *Member) GetRemark() string {
	if u == nil {
		return ""
	}
	return u.Remark
}

func (u *Member) GetTeamMemberID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *Member) GetUserID() uint32 {
	if u == nil {
		return 0
	}
	return u.UserID
}

func (u *Member) GetInviterID() uint32 {
	if u == nil {
		return 0
	}
	return u.InviterID
}

func (u *Member) GetPosition() vobj.Role {
	if u == nil {
		return vobj.RoleUnknown
	}
	return u.Position
}

func (u *Member) GetStatus() vobj.MemberStatus {
	if u == nil {
		return vobj.MemberStatusUnknown
	}
	return u.Status
}

func (u *Member) GetRoles() []do.TeamRole {
	if u == nil {
		return nil
	}
	return slices.Map(u.Roles, func(r *Role) do.TeamRole { return r })
}

func (u *Member) TableName() string {
	return tableNameTeamMember
}
