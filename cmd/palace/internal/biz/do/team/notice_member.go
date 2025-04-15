package team

import (
	"time"

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

func (u *NoticeMember) GetID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *NoticeMember) GetTeamID() uint32 {
	if u == nil {
		return 0
	}
	return u.TeamID
}

func (u *NoticeMember) GetNoticeGroupID() uint32 {
	if u == nil {
		return 0
	}
	return u.NoticeGroupID
}

func (u *NoticeMember) GetMemberID() uint32 {
	if u == nil {
		return 0
	}
	return u.MemberID
}

func (u *NoticeMember) GetNoticeType() vobj.NoticeType {
	if u == nil {
		return 0
	}
	return u.NoticeType
}

func (u *NoticeMember) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *NoticeMember) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *NoticeMember) GetNoticeGroup() do.NoticeGroup {
	if u == nil || u.NoticeGroup == nil {
		return nil
	}
	return u.NoticeGroup
}

func (u *NoticeMember) GetMember() do.TeamMember {
	if u == nil || u.Member == nil {
		return nil
	}
	return u.Member
}

func (u *NoticeMember) TableName() string {
	return tableNameNoticeMember
}
