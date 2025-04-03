package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameMenu = "sys_menus"

type Menu struct {
	do.BaseModel
	Name     string            `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__menu__name,priority:1;comment:菜单名称" json:"name"`
	Path     string            `gorm:"column:path;type:varchar(255);not null;uniqueIndex:idx__menu__path,priority:1;comment:菜单路径" json:"path"`
	Status   vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Icon     string            `gorm:"column:icon;type:varchar(64);not null;comment:图标" json:"icon"`
	ParentID uint32            `gorm:"column:parent_id;type:int unsigned;not null;default:0;comment:父级id" json:"parentID"`
	Type     vobj.MenuType     `gorm:"column:type;type:tinyint(2);not null;comment:菜单类型" json:"type"`

	Resources []*Resource `gorm:"foreignKey:MenuID;references:ID" json:"resources"`
	Parent    *Menu       `gorm:"foreignKey:ParentID;references:ID" json:"parent"`
}

func (u *Menu) TableName() string {
	return tableNameMenu
}
