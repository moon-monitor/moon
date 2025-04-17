package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeam(
	userRepo repository.User,
	teamRepo repository.Team,
	transaction repository.Transaction,
	logger log.Logger,
) *Team {
	return &Team{
		helper:      log.NewHelper(log.With(logger, "module", "biz.team")),
		userRepo:    userRepo,
		teamRepo:    teamRepo,
		transaction: transaction,
	}
}

type Team struct {
	helper      *log.Helper
	userRepo    repository.User
	teamRepo    repository.Team
	transaction repository.Transaction
}

func (t *Team) SaveTeam(ctx context.Context, req *bo.SaveOneTeamRequest) error {
	return t.transaction.MainExec(ctx, func(ctx context.Context) error {
		var (
			teamDo do.Team
			err    error
		)
		defer func() {
			if err != nil {
				t.helper.Errorw("msg", "save team fail", "err", err)
				return
			}
			if err = t.userRepo.AppendTeam(ctx, teamDo); err != nil {
				t.helper.Errorw("msg", "append team to user fail", "err", err)
				return
			}
		}()
		if req.GetID() <= 0 {
			createParams, err := req.WithCreateTeamRequest(ctx)
			if err != nil {
				return err
			}
			teamDo, err = t.teamRepo.Create(ctx, createParams)
			return err
		}
		teamInfo, err := t.teamRepo.FindByID(ctx, req.GetID())
		if err != nil {
			return err
		}
		updateTeamParams := req.WithUpdateTeamRequest(teamInfo)
		teamDo, err = t.teamRepo.Update(ctx, updateTeamParams)
		return err
	})
}
