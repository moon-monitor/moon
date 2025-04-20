package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToNoticeMemberItem(noticeMember do.NoticeMember) *common.NoticeMemberItem {
	if validate.IsNil(noticeMember) {
		return nil
	}
	return &common.NoticeMemberItem{
		NoticeGroupId: noticeMember.GetNoticeGroupID(),
		UserId:        noticeMember.GetUserID(),
		NoticeType:    common.NoticeType(noticeMember.GetNoticeType()),
		NoticeGroup:   ToNoticeGroupItem(noticeMember.GetNoticeGroup()),
		Member:        ToTeamMemberBaseItem(noticeMember.GetMember()),
	}
}

func ToNoticeMemberItems(noticeMembers []do.NoticeMember) []*common.NoticeMemberItem {
	return slices.Map(noticeMembers, ToNoticeMemberItem)
}

func ToNoticeGroupItem(noticeGroup do.NoticeGroup) *common.NoticeGroupItem {
	return &common.NoticeGroupItem{
		NoticeGroupId: noticeGroup.GetID(),
		CreatedAt:     noticeGroup.GetCreatedAt().Format(time.DateTime),
		UpdatedAt:     noticeGroup.GetUpdatedAt().Format(time.DateTime),
		Name:          noticeGroup.GetName(),
		Remark:        noticeGroup.GetRemark(),
		Status:        common.GlobalStatus(noticeGroup.GetStatus()),
		NoticeMembers: ToNoticeMemberItems(noticeGroup.GetNoticeMembers()),
		Hooks:         ToNoticeHookItems(noticeGroup.GetHooks()),
		Creator:       ToUserBaseItem(noticeGroup.GetCreator()),
	}
}

func ToNoticeGroupItems(noticeGroups []do.NoticeGroup) []*common.NoticeGroupItem {
	return slices.Map(noticeGroups, ToNoticeGroupItem)
}
