package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamService struct {
	palacev1.UnimplementedTeamServer

	teamBiz *biz.Team
}

func NewTeamService(teamBiz *biz.Team) *TeamService {
	return &TeamService{
		teamBiz: teamBiz,
	}
}

func (s *TeamService) SaveTeam(ctx context.Context, req *palacev1.SaveTeamRequest) (*common.EmptyReply, error) {
	params := bo.NewSaveOneTeamRequest(req, req.GetId())
	if err := s.teamBiz.SaveTeam(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存团队信息成功"}, nil
}

func (s *TeamService) GetTeam(ctx context.Context, req *common.EmptyRequest) (*palacev1.GetTeamReply, error) {
	return &palacev1.GetTeamReply{}, nil
}

func (s *TeamService) GetTeamResources(ctx context.Context, req *common.EmptyRequest) (*palacev1.GetTeamResourcesReply, error) {
	return &palacev1.GetTeamResourcesReply{}, nil
}

func (s *TeamService) TransferTeam(ctx context.Context, req *palacev1.TransferTeamRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) InviteMember(ctx context.Context, req *palacev1.InviteMemberRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) RemoveMember(ctx context.Context, req *palacev1.RemoveMemberRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) GetTeamMembers(ctx context.Context, req *common.EmptyRequest) (*palacev1.GetTeamMembersReply, error) {
	return &palacev1.GetTeamMembersReply{}, nil
}

func (s *TeamService) UpdateMemberPosition(ctx context.Context, req *palacev1.UpdateMemberPositionRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) UpdateMemberStatus(ctx context.Context, req *palacev1.UpdateMemberStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) UpdateMemberRoles(ctx context.Context, req *palacev1.UpdateMemberRolesRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) GetTeamRoles(ctx context.Context, req *common.EmptyRequest) (*palacev1.GetTeamRolesReply, error) {
	return &palacev1.GetTeamRolesReply{}, nil
}

func (s *TeamService) SaveTeamRole(ctx context.Context, req *palacev1.SaveTeamRoleRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) DeleteTeamRole(ctx context.Context, req *palacev1.DeleteTeamRoleRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamService) SaveEmailConfig(ctx context.Context, req *palacev1.SaveEmailConfigRequest) (*common.EmptyReply, error) {
	if err := s.teamBiz.SaveEmailConfig(ctx, build.ToSaveEmailConfigRequest(req)); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存邮件配置成功"}, nil
}

func (s *TeamService) GetEmailConfigs(ctx context.Context, req *palacev1.GetEmailConfigsRequest) (*palacev1.GetEmailConfigsReply, error) {
	params := build.ToListEmailConfigRequest(req)
	config, err := s.teamBiz.GetEmailConfigs(ctx, params)
	if err != nil {
		return nil, err
	}
	return build.ToEmailConfigReply(config), nil
}

func (s *TeamService) SaveSMSConfig(ctx context.Context, req *palacev1.SaveSMSConfigRequest) (*common.EmptyReply, error) {
	if err := s.teamBiz.SaveSMSConfig(ctx, build.ToSaveSMSConfigRequest(req)); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存短信配置成功"}, nil
}

func (s *TeamService) GetSMSConfigs(ctx context.Context, req *palacev1.GetSMSConfigsRequest) (*palacev1.GetSMSConfigsReply, error) {
	params := build.ToListSMSConfigRequest(req)
	config, err := s.teamBiz.GetSMSConfigs(ctx, params)
	if err != nil {
		return nil, err
	}
	return build.ToSMSConfigReply(config), nil
}
