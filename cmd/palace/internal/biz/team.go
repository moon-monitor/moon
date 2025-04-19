package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewTeam(
	userRepo repository.User,
	teamRepo repository.Team,
	teamEmailConfigRepo repository.TeamEmailConfig,
	teamSMSConfigRepo repository.TeamSMSConfig,
	teamRoleRepo repository.TeamRole,
	menuRepo repository.Menu,
	transaction repository.Transaction,
	logger log.Logger,
) *Team {
	return &Team{
		helper:              log.NewHelper(log.With(logger, "module", "biz.team")),
		userRepo:            userRepo,
		teamRepo:            teamRepo,
		teamEmailConfigRepo: teamEmailConfigRepo,
		teamSMSConfigRepo:   teamSMSConfigRepo,
		teamRoleRepo:        teamRoleRepo,
		menuRepo:            menuRepo,
		transaction:         transaction,
	}
}

type Team struct {
	helper              *log.Helper
	userRepo            repository.User
	teamRepo            repository.Team
	teamEmailConfigRepo repository.TeamEmailConfig
	teamSMSConfigRepo   repository.TeamSMSConfig
	teamRoleRepo        repository.TeamRole
	menuRepo            repository.Menu
	transaction         repository.Transaction
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

// SaveEmailConfig saves the email configuration for a team
func (t *Team) SaveEmailConfig(ctx context.Context, req *bo.SaveEmailConfigRequest) error {
	return t.transaction.BizExec(ctx, func(ctx context.Context) error {
		if req.ID <= 0 {
			return t.teamEmailConfigRepo.Create(ctx, req)
		}
		emailConfig, err := t.teamEmailConfigRepo.Get(ctx, req.ID)
		if err != nil {
			return err
		}
		return t.teamEmailConfigRepo.Update(ctx, req.WithEmailConfig(emailConfig))
	})
}

// GetEmailConfigs retrieves the email configuration for a team
func (t *Team) GetEmailConfigs(ctx context.Context, req *bo.ListEmailConfigRequest) (*bo.ListEmailConfigListReply, error) {
	configListReply, err := t.teamEmailConfigRepo.List(ctx, req)
	if err != nil {
		return nil, merr.ErrorInternalServerError("failed to get email config").WithCause(err)
	}

	return configListReply, nil
}

// SaveSMSConfig saves the SMS configuration for a team
func (t *Team) SaveSMSConfig(ctx context.Context, req *bo.SaveSMSConfigRequest) error {
	return t.transaction.BizExec(ctx, func(ctx context.Context) error {
		if req.ID <= 0 {
			return t.teamSMSConfigRepo.Create(ctx, req)
		}
		smsConfig, err := t.teamSMSConfigRepo.Get(ctx, req.ID)
		if err != nil {
			return err
		}
		return t.teamSMSConfigRepo.Update(ctx, req.WithSMSConfig(smsConfig))
	})
}

// GetSMSConfigs retrieves SMS configurations for a team
func (t *Team) GetSMSConfigs(ctx context.Context, req *bo.ListSMSConfigRequest) (*bo.ListSMSConfigListReply, error) {
	return t.teamSMSConfigRepo.List(ctx, req)
}

func (t *Team) SaveTeamRole(ctx context.Context, req *bo.SaveTeamRoleReq) error {
	return t.transaction.BizExec(ctx, func(ctx context.Context) error {
		if req.GetID() <= 0 {
			return t.teamRoleRepo.Create(ctx, req)
		}
		teamRoleDo, err := t.teamRoleRepo.Get(ctx, req.GetID())
		if err != nil {
			return err
		}
		if len(req.GetMenuIds()) > 0 {
			menuDos, err := t.menuRepo.Find(ctx, req.GetMenuIds())
			if err != nil {
				return err
			}
			req.WithMenus(menuDos)
		}

		return t.teamRoleRepo.Update(ctx, req.WithRole(teamRoleDo))
	})
}

func (t *Team) GetTeamRoles(ctx context.Context, req *bo.ListRoleReq) (*bo.ListTeamRoleReply, error) {
	return t.teamRoleRepo.List(ctx, req)
}

func (t *Team) DeleteTeamRole(ctx context.Context, roleID uint32) error {
	return t.teamRoleRepo.Delete(ctx, roleID)
}

func (t *Team) UpdateTeamRoleStatus(ctx context.Context, req *bo.UpdateRoleStatusReq) error {
	return t.teamRoleRepo.UpdateStatus(ctx, req)
}

func (t *Team) ListTeam(ctx context.Context, req *bo.TeamListRequest) (*bo.TeamListReply, error) {
	return t.teamRepo.List(ctx, req)
}
