package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// ToTeamItem 将系统Team对象转换为TeamItem proto对象
func ToTeamItem(team do.Team) *common.TeamItem {
	if validate.IsNil(team) {
		return nil
	}

	return &common.TeamItem{
		Id:              team.GetID(),
		Uuid:            team.GetUUID().String(),
		Name:            team.GetName(),
		Remark:          team.GetRemark(),
		Logo:            team.GetLogo(),
		Status:          common.TeamStatus(team.GetStatus()),
		Creator:         UserToUserBaseItemProto(team.GetCreator()),
		Leader:          UserToUserBaseItemProto(team.GetLeader()),
		Admins:          UsersToUserBaseItemsProto(team.GetAdmins()),
		CreatedAt:       team.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:       team.GetUpdatedAt().Format(time.RFC3339),
		MemberCount:     0,
		StrategyCount:   0,
		DatasourceCount: 0,
	}
}

func ToTeamBaseItem(team do.Team) *common.TeamBaseItem {
	if validate.IsNil(team) {
		return nil
	}

	return &common.TeamBaseItem{
		Id:     team.GetID(),
		Name:   team.GetName(),
		Remark: team.GetRemark(),
		Logo:   team.GetLogo(),
	}
}

// ToTeamItems 将系统Team对象切片转换为TeamItem proto对象切片
func ToTeamItems(teams []do.Team) []*common.TeamItem {
	return slices.Map(teams, ToTeamItem)
}

func ToTeamBaseItems(teams []do.Team) []*common.TeamBaseItem {
	return slices.Map(teams, ToTeamBaseItem)
}

func ToTeamListRequest(req *palace.GetTeamListRequest) *bo.TeamListRequest {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.TeamListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Keyword:           req.GetKeyword(),
		Status:            slices.Map(req.GetStatus(), func(status common.TeamStatus) vobj.TeamStatus { return vobj.TeamStatus(status) }),
		UserIds:           nil,
		LeaderId:          req.GetLeaderId(),
		CreatorId:         req.GetCreatorId(),
	}
}

func ToTeamMemberListRequest(req *palace.GetTeamMembersRequest, teamId uint32) *bo.TeamMemberListRequest {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.TeamMemberListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Keyword:           req.GetKeyword(),
		Status:            slices.Map(req.GetStatus(), func(status common.MemberStatus) vobj.MemberStatus { return vobj.MemberStatus(status) }),
		Positions:         slices.Map(req.GetPositions(), func(position common.MemberPosition) vobj.Role { return vobj.Role(position) }),
		TeamId:            teamId,
	}
}

func ToTeamMemberItem(member do.TeamMember) *common.TeamMemberItem {
	if validate.IsNil(member) {
		return nil
	}
	return &common.TeamMemberItem{
		Id:        member.GetTeamMemberID(),
		User:      UserToUserBaseItemProto(nil),
		Position:  common.MemberPosition(member.GetPosition()),
		Status:    common.MemberStatus(member.GetStatus()),
		Inviter:   UserToUserBaseItemProto(nil),
		Roles:     ToTeamRoleItems(member.GetRoles()),
		CreatedAt: member.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: member.GetUpdatedAt().Format(time.RFC3339),
	}
}

func ToTeamMemberItems(members []do.TeamMember) []*common.TeamMemberItem {
	return slices.Map(members, ToTeamMemberItem)
}
