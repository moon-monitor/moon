package team

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameNoticeGroup = "team_notice_group"

type NoticeGroup struct {
	do.TeamModel

	Name   string            `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`

	Members NoticeMembers `gorm:"column:members;type:text;not null;comment:通知成员列表" json:"members"`
	Hooks   []*NoticeHook `gorm:"many2many:team_notice_group_hooks" json:"hooks"`
}

func (n *NoticeGroup) TableName() string {
	return tableNameNoticeGroup
}

type NoticeMembers []*NoticeMember

type NoticeMember struct {
	UserID   uint32          `json:"user_id"`
	MemberID uint32          `json:"member_id"`
	Notice   vobj.NoticeType `json:"notice"`
}

func (n NoticeMembers) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *NoticeMembers) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), n)
}
