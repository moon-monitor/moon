package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

var _ do.NoticeMember = (*NoticeMember)(nil)

const tableNameNoticeMember = "team_notice_members"

type NoticeMember struct {
	do.TeamModel
	NoticeGroupID uint32          `gorm:"column:notice_group_id;type:int(10) unsigned;not null;comment:通知组ID" json:"noticeGroupID"`
	MemberID      uint32          `gorm:"column:member_id;type:int(10) unsigned;not null;comment:成员ID" json:"memberID"`
	NoticeType    vobj.NoticeType `gorm:"column:notice_type;type:int(10) unsigned;not null;comment:通知类型" json:"noticeType"`
	NoticeGroup   *NoticeGroup    `gorm:"foreignKey:NoticeGroupID;references:ID" json:"noticeGroup"`
	Member        *Member         `gorm:"foreignKey:MemberID;references:ID" json:"member"`
}

func (n *NoticeMember) GetNoticeGroupID() uint32 {
	if n == nil {
		return 0
	}
	return n.NoticeGroupID
}

func (n *NoticeMember) GetMemberID() uint32 {
	if n == nil {
		return 0
	}
	return n.MemberID
}

func (n *NoticeMember) GetNoticeType() vobj.NoticeType {
	if n == nil {
		return vobj.NoticeTypeNone
	}
	return n.NoticeType
}

func (n *NoticeMember) GetNoticeGroup() do.NoticeGroup {
	if n == nil || n.NoticeGroup == nil {
		return nil
	}
	return n.NoticeGroup
}

func (n *NoticeMember) GetMember() do.TeamMember {
	if n == nil || n.Member == nil {
		return nil
	}
	return n.Member
}

func (n *NoticeMember) TableName() string {
	return tableNameNoticeMember
}
