package team

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameNoticeHook = "team_notice_hooks"

type NoticeHook struct {
	do.TeamModel

	Name    string            `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark  string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	URL     string            `gorm:"column:url;type:varchar(255);not null;comment:URL" json:"url"`
	Secret  string            `gorm:"column:secret;type:varchar(255);not null;comment:密钥" json:"secret"`
	Headers NoticeHookHeaders `gorm:"column:headers;type:text;not null;comment:请求头" json:"headers"`
}

func (n *NoticeHook) TableName() string {
	return tableNameNoticeHook
}

type NoticeHookHeaders map[string]string

func (n NoticeHookHeaders) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *NoticeHookHeaders) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), n)
}
