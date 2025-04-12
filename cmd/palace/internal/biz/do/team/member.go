package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeamMember = "team_members"

type Member struct {
	do.TeamModel
	UserID    uint32            `gorm:"column:user_id;type:int unsigned;not null;comment:用户ID" json:"userID"`
	InviterID uint32            `gorm:"column:inviter_id;type:int unsigned;not null;comment:邀请者ID" json:"inviterID"`
	Position  vobj.Role         `gorm:"column:position;type:tinyint(2);not null;comment:职位" json:"position"`
	Status    vobj.MemberStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Roles     []*Role           `gorm:"many2many:sys_team_member_roles" json:"roles"`
}

func (u *Member) TableName() string {
	return tableNameTeamMember
}
