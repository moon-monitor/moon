package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
)

func NewTeam(
	teamRepo repository.Team,
	logger log.Logger,
) *Team {
	return &Team{
		helper:   log.NewHelper(log.With(logger, "module", "biz.team")),
		teamRepo: teamRepo,
	}
}

type Team struct {
	helper   *log.Helper
	teamRepo repository.Team
}

func (t *Team) SaveTeam(ctx context.Context, req *bo.SaveOneTeamRequest) error {
	if req.GetID() <= 0 {
		params, err := req.WithCreateTeamRequest(ctx)
		if err != nil {
			return err
		}
		return t.teamRepo.Create(ctx, params)
	}
	teamInfo, err := t.teamRepo.FindByID(ctx, req.GetID())
	if err != nil {
		return err
	}
	updateTeamParams := req.WithUpdateTeamRequest(teamInfo)
	return t.teamRepo.Update(ctx, updateTeamParams)
}
