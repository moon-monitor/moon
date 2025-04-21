package build

import (
	"github.com/google/uuid"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
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
		CreatorModel:  do.CreatorModel{},
		Name:          "",
		Status:        0,
		Remark:        "",
		Logo:          "",
		LeaderID:      0,
		UUID:          uuid.UUID{},
		Capacity:      0,
		Leader:        nil,
		Admins:        nil,
		Resources:     nil,
		BizDBConfig:   nil,
		AlarmDBConfig: nil,
	}
}

func ToTeams(teamDos []do.Team) []*system.Team {
	return slices.Map(teamDos, ToTeam)
}
