package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

const tableNameTeamRole = "sys_team_roles"

type TeamRole struct {
	do.TeamModel
	Name   string `gorm:"column:name;type:varchar(64);not null;comment:角色名" json:"name"`
	Remark string `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`

	Members   []*TeamMember `gorm:"many2many:sys_team_member_roles;foreignKey:ID;joinForeignKey:TeamRoleID;references:ID;joinReferences:TeamMemberID" json:"members"`
	Resources []*Resource   `gorm:"many2many:sys_team_role_resources;foreignKey:ID;joinForeignKey:TeamRoleID;references:ID;joinReferences:ResourceID" json:"resources"`
}

func (u *TeamRole) TableName() string {
	return tableNameTeamRole
}
