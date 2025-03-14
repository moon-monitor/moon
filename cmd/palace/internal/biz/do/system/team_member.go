package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeamMember = "sys_team_members"

type TeamMember struct {
	do.TeamModel
	UserID    uint32    `gorm:"column:user_id;type:int unsigned;not null;comment:用户ID" json:"user_id,omitempty"`
	InviterID uint32    `gorm:"column:inviter_id;type:int unsigned;not null;comment:邀请者ID" json:"inviter_id,omitempty"`
	Position  vobj.Role `gorm:"column:position;type:tinyint(2);not null;comment:职位" json:"position"`

	User    *User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Inviter *User       `gorm:"foreignKey:InvitorID;references:ID" json:"inviter"`
	Roles   []*TeamRole `gorm:"many2many:sys_team_member_roles;foreignKey:ID;joinForeignKey:TeamMemberID;references:ID;joinReferences:RoleID" json:"roles"`
}

func (u *TeamMember) TableName() string {
	return tableNameTeamMember
}
