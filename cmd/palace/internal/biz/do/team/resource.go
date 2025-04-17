package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

var _ do.Resource = (*Resource)(nil)

const tableNameResource = "team_resources"

type Resource struct {
	do.TeamModel
	Name   string             `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__api__name,priority:1;comment:api名称" json:"name"`
	Path   string             `gorm:"column:path;type:varchar(255);not null;uniqueIndex:idx__api__path,priority:1;comment:api路径" json:"path"`
	Status vobj.GlobalStatus  `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Allow  vobj.ResourceAllow `gorm:"column:allow;type:tinyint(2);not null;comment:放行规则" json:"allow"`
	Remark string             `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	MenuID uint32             `gorm:"column:menu_id;type:int unsigned;not null;comment:菜单id" json:"menuID"`
	Menus  []*Menu            `gorm:"many2many:team_menu_resources" json:"menus"`
}

func (u *Resource) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *Resource) GetPath() string {
	if u == nil {
		return ""
	}
	return u.Path
}

func (u *Resource) GetStatus() vobj.GlobalStatus {
	if u == nil {
		return vobj.GlobalStatusUnknown
	}
	return u.Status
}

func (u *Resource) GetAllow() vobj.ResourceAllow {
	if u == nil {
		return vobj.ResourceAllowUnknown
	}
	return u.Allow
}

func (u *Resource) GetRemark() string {
	if u == nil {
		return ""
	}
	return u.Remark
}

func (u *Resource) GetMenus() []do.Menu {
	if u == nil {
		return nil
	}
	return slices.Map(u.Menus, func(m *Menu) do.Menu { return m })
}

func (u *Resource) TableName() string {
	return tableNameResource
}
