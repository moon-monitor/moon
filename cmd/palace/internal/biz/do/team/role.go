package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeamRole = "team_roles"

type Role struct {
	do.TeamModel
	Name    string            `gorm:"column:name;type:varchar(64);not null;comment:角色名" json:"name"`
	Remark  string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Members []*Member         `gorm:"many2many:team_member_roles" json:"members"`
	Menus   []*Menu           `gorm:"many2many:team_role_menus" json:"menus"`
}

func (u *Role) TableName() string {
	return tableNameTeamRole
}
