package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameNoticeGroup = "team_notice_groups"

type NoticeGroup struct {
	do.TeamModel
	Name    string            `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark  string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Members []*NoticeMember   `gorm:"many2many:team_notice_group_members" json:"members"`
	Hooks   []*NoticeHook     `gorm:"many2many:team_notice_group_hooks" json:"hooks"`
}

func (n *NoticeGroup) TableName() string {
	return tableNameNoticeGroup
}
