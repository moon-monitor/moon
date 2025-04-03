package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameRole = "sys_roles"

type SysRole struct {
	do.BaseModel
	Name   string            `gorm:"column:name;type:varchar(64);not null;comment:角色名" json:"name"`
	Remark string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`

	Users []*User `gorm:"many2many:sys_user_roles;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:UserID" json:"users"`
	Menus []*Menu `gorm:"many2many:sys_role_menus;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:MenuID" json:"menus"`
}

func (u *SysRole) TableName() string {
	return tableNameRole
}
