package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewTeam(
	userRepo repository.User,
	teamRepo repository.Team,
	teamEmailConfigRepo repository.TeamEmailConfig,
	teamSMSConfigRepo repository.TeamSMSConfig,
	teamRoleRepo repository.TeamRole,
	menuRepo repository.Menu,
	operateLogRepo repository.OperateLog,
	memberRepo repository.Member,
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
		operateLogRepo:      operateLogRepo,
		memberRepo:          memberRepo,
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
	operateLogRepo      repository.OperateLog
	memberRepo          repository.Member
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

func (t *Team) OperateLogList(ctx context.Context, req *bo.OperateLogListRequest) (*bo.OperateLogListReply, error) {
	return t.operateLogRepo.TeamList(ctx, req)
}

func (t *Team) GetTeamByID(ctx context.Context, teamID uint32) (do.Team, error) {
	return t.teamRepo.FindByID(ctx, teamID)
}

func (t *Team) GetTeamMembers(ctx context.Context, req *bo.TeamMemberListRequest) (*bo.TeamMemberListReply, error) {
	return t.memberRepo.List(ctx, req)
}

func (t *Team) UpdateMemberPosition(ctx context.Context, req *bo.UpdateMemberPositionReq) error {
	userId, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}
	operatorDo, err := t.memberRepo.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}
	req.WithOperator(operatorDo)
	memberDo, err := t.memberRepo.Get(ctx, req.MemberID)
	if err != nil {
		return err
	}
	req.WithMember(memberDo)
	if err := req.Validate(); err != nil {
		return err
	}
	return t.memberRepo.UpdatePosition(ctx, req)
}

func (t *Team) UpdateMemberStatus(ctx context.Context, req *bo.UpdateMemberStatusReq) error {
	userId, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}
	operatorDo, err := t.memberRepo.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}
	req.WithOperator(operatorDo)
	members, err := t.memberRepo.Find(ctx, req.MemberIds)
	if err != nil {
		return err
	}
	req.WithMembers(members)
	if err := req.Validate(); err != nil {
		return err
	}
	return t.memberRepo.UpdateStatus(ctx, req)
}

func (t *Team) UpdateMemberRoles(ctx context.Context, req *bo.UpdateMemberRolesReq) error {
	userId, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}
	operatorDo, err := t.memberRepo.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}
	req.WithOperator(operatorDo)
	memberDo, err := t.memberRepo.Get(ctx, req.MemberId)
	if err != nil {
		return err
	}
	req.WithMember(memberDo)
	roles, err := t.teamRoleRepo.Find(ctx, req.RoleIds)
	if err != nil {
		return err
	}
	req.WithRoles(roles)
	if err := req.Validate(); err != nil {
		return err
	}
	return t.memberRepo.UpdateRoles(ctx, req)
}

func (t *Team) RemoveMember(ctx context.Context, req *bo.RemoveMemberReq) error {
	userId, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}
	operatorDo, err := t.memberRepo.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}
	req.WithOperator(operatorDo)
	memberDo, err := t.memberRepo.Get(ctx, req.MemberID)
	if err != nil {
		return err
	}
	req.WithMember(memberDo)
	if err := req.Validate(); err != nil {
		return err
	}
	return t.memberRepo.UpdateStatus(ctx, req)
}
