package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToTeam(teamDo do.Team) *system.Team {
	if validate.IsNil(teamDo) {
		return nil
	}
	team, ok := teamDo.(*system.Team)
	if ok {
		return team
	}
	return &system.Team{
		CreatorModel:  ToCreatorModel(teamDo),
		Name:          teamDo.GetName(),
		Status:        teamDo.GetStatus(),
		Remark:        teamDo.GetRemark(),
		Logo:          teamDo.GetLogo(),
		LeaderID:      teamDo.GetLeaderID(),
		UUID:          teamDo.GetUUID(),
		Capacity:      teamDo.GetCapacity(),
		Leader:        ToUser(teamDo.GetLeader()),
		Admins:        ToUsers(teamDo.GetAdmins()),
		Resources:     nil,
		BizDBConfig:   crypto.NewObject(teamDo.GetBizDBConfig()),
		AlarmDBConfig: crypto.NewObject(teamDo.GetAlarmDBConfig()),
	}
}

func ToTeams(teamDos []do.Team) []*system.Team {
	return slices.Map(teamDos, ToTeam)
}

func ToTeamMember(memberDo do.TeamMember) *system.TeamMember {
	if validate.IsNil(memberDo) {
		return nil
	}
	member, ok := memberDo.(*system.TeamMember)
	if ok {
		return member
	}
	return &system.TeamMember{
		TeamModel:  ToTeamModel(memberDo),
		MemberName: memberDo.GetMemberName(),
		Remark:     memberDo.GetRemark(),
		UserID:     memberDo.GetUserID(),
		InviterID:  memberDo.GetInviterID(),
		Position:   memberDo.GetPosition(),
		Status:     memberDo.GetStatus(),
		Roles:      ToTeamRoles(memberDo.GetRoles()),
		User:       ToUser(memberDo.GetUser()),
		Inviter:    ToUser(memberDo.GetInviter()),
	}
}

func ToTeamMembers(memberDos []do.TeamMember) []*system.TeamMember {
	return slices.Map(memberDos, ToTeamMember)
}
