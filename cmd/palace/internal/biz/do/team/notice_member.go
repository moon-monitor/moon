package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameNoticeMember = "team_notice_members"

type NoticeMember struct {
	do.TeamModel
	NoticeGroupID uint32          `gorm:"column:notice_group_id;type:int(10) unsigned;not null;comment:通知组ID" json:"noticeGroupID"`
	MemberID      uint32          `gorm:"column:member_id;type:int(10) unsigned;not null;comment:成员ID" json:"memberID"`
	NoticeType    vobj.NoticeType `gorm:"column:notice_type;type:int(10) unsigned;not null;comment:通知类型" json:"noticeType"`
	NoticeGroup   *NoticeGroup    `gorm:"foreignKey:NoticeGroupID;references:ID" json:"noticeGroup"`
	Member        *Member         `gorm:"foreignKey:MemberID;references:ID" json:"member"`
}

func (u *NoticeMember) TableName() string {
	return tableNameNoticeMember
}
