package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

const tableNameNoticeHook = "team_notice_hooks"

type NoticeHook struct {
	do.TeamModel
	Name         string            `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark       string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status       vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	URL          string            `gorm:"column:url;type:varchar(255);not null;comment:URL" json:"url"`
	Method       vobj.HTTPMethod   `gorm:"column:method;type:tinyint(2);not null;comment:请求方法" json:"method"`
	Secret       string            `gorm:"column:secret;type:varchar(255);not null;comment:密钥" json:"secret"`
	Headers      kv.StringMap      `gorm:"column:headers;type:text;not null;comment:请求头" json:"headers"`
	NoticeGroups []*NoticeGroup    `gorm:"many2many:team_notice_group_hooks" json:"noticeGroups"`
	APP          vobj.HookApp      `gorm:"column:app;type:tinyint(2);not null;comment:应用" json:"app"`
}

func (n *NoticeHook) TableName() string {
	return tableNameNoticeHook
}
