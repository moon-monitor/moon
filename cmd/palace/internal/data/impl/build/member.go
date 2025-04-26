package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToStrategyMember(member do.NoticeMember) *team.NoticeMember {
	if validate.IsNil(member) {
		return nil
	}
	if member, ok := member.(*team.NoticeMember); ok {
		return member
	}
	item := &team.NoticeMember{
		TeamModel:     ToTeamModel(member),
		NoticeGroupID: member.GetNoticeGroupID(),
		UserID:        member.GetUserID(),
		NoticeType:    member.GetNoticeType(),
		NoticeGroup:   ToStrategyNotice(member.GetNoticeGroup()),
	}
	return item
}

func ToStrategyMembers(members []do.NoticeMember) []*team.NoticeMember {
	return slices.Map(members, ToStrategyMember)
}
