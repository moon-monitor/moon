package build

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToStrategyMember(ctx context.Context, member do.NoticeMember) *team.NoticeMember {
	if member, ok := member.(*team.NoticeMember); ok {
		member.WithContext(ctx)
		return member
	}
	item := &team.NoticeMember{
		TeamModel: do.TeamModel{
			CreatorModel: do.CreatorModel{
				BaseModel: do.BaseModel{
					ID:        member.GetID(),
					CreatedAt: member.GetCreatedAt(),
					UpdatedAt: member.GetUpdatedAt(),
					DeletedAt: member.GetDeletedAt(),
				},
				CreatorID: member.GetCreatorID(),
			},
			TeamID: member.GetTeamID(),
		},
		NoticeGroupID: member.GetNoticeGroupID(),
		UserID:        member.GetUserID(),
		NoticeType:    member.GetNoticeType(),
		NoticeGroup:   ToStrategyNotice(ctx, member.GetNoticeGroup()),
	}
	item.WithContext(ctx)
	return item
}

func ToStrategyMembers(ctx context.Context, members []do.NoticeMember) []*team.NoticeMember {
	return slices.Map(members, func(member do.NoticeMember) *team.NoticeMember {
		return ToStrategyMember(ctx, member)
	})
}
