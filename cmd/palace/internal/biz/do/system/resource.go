package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameResource = "sys_resources"

type Resource struct {
	do.BaseModel
	Name   string              `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__api__name,priority:1;comment:api名称" json:"name"`
	Path   string              `gorm:"column:path;type:varchar(255);not null;uniqueIndex:idx__api__path,priority:1;comment:api路径" json:"path"`
	Status vobj.ResourceStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Domain vobj.ResourceDomain `gorm:"column:domain;type:tinyint(2);not null;comment:领域" json:"domain"`
	Module vobj.ResourceDomain `gorm:"column:module;type:tinyint(2);not null;comment:模块" json:"module"`
	Allow  vobj.ResourceAllow  `gorm:"column:allow;type:tinyint(2);not null;comment:放行规则" json:"allow"`
}

func (u *Resource) TableName() string {
	return tableNameResource
}
