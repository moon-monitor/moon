package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func NewSystemService(
	userBiz *biz.UserBiz,
	messageBiz *biz.Message,
	teamBiz *biz.Team,
	systemBiz *biz.System,
) *SystemService {
	return &SystemService{
		userBiz:    userBiz,
		messageBiz: messageBiz,
		teamBiz:    teamBiz,
		systemBiz:  systemBiz,
	}
}

type SystemService struct {
	palacev1.UnimplementedSystemServer
	userBiz    *biz.UserBiz
	messageBiz *biz.Message
	teamBiz    *biz.Team
	systemBiz  *biz.System
}

func (s *SystemService) UpdateUser(ctx context.Context, req *palacev1.UpdateUserRequest) (*common.EmptyReply, error) {
	params := build.ToUserUpdateInfo(req)
	if err := s.userBiz.UpdateUserBaseInfo(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新用户信息成功"}, nil
}

func (s *SystemService) UpdateUserStatus(ctx context.Context, req *palacev1.UpdateUserStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateUserStatusRequest{
		UserIds: req.GetUserIds(),
		Status:  vobj.UserStatus(req.GetStatus()),
	}
	if err := s.userBiz.UpdateUserStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新用户状态成功"}, nil
}

func (s *SystemService) ResetUserPassword(ctx context.Context, req *palacev1.ResetUserPasswordRequest) (*common.EmptyReply, error) {
	params := &bo.ResetUserPasswordRequest{
		UserId:       req.GetUserId(),
		SendEmailFun: s.messageBiz.SendEmail,
	}
	if err := s.userBiz.ResetUserPassword(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "重置用户密码成功"}, nil
}

func (s *SystemService) UpdateUserPosition(ctx context.Context, req *palacev1.UpdateUserPositionRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateUserPositionRequest{
		UserId:   req.GetUserId(),
		Position: vobj.Role(req.GetPosition()),
	}
	if err := s.userBiz.UpdateUserPosition(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新用户职位成功"}, nil
}

func (s *SystemService) GetUser(ctx context.Context, req *palacev1.GetUserRequest) (*common.UserItem, error) {
	userDo, err := s.userBiz.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return build.ToUserItem(userDo), nil
}

func (s *SystemService) GetUserList(ctx context.Context, req *palacev1.GetUserListRequest) (*palacev1.GetUserListReply, error) {
	params := build.ToUserListRequest(req)
	userReply, err := s.userBiz.ListUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetUserListReply{
		Items:      build.ToUserItems(userReply.Items),
		Pagination: build.ToPaginationReply(userReply.PaginationReply),
	}, nil
}

func (s *SystemService) GetTeamList(ctx context.Context, req *palacev1.GetTeamListRequest) (*palacev1.GetTeamListReply, error) {
	params := build.ToTeamListRequest(req)
	teamReply, err := s.teamBiz.ListTeam(ctx, params)
	if err != nil {
		return nil, err
	}

	return &palacev1.GetTeamListReply{
		Items:      build.ToTeamItems(teamReply.Items),
		Pagination: build.ToPaginationReply(teamReply.PaginationReply),
	}, nil
}

func (s *SystemService) GetSystemRole(ctx context.Context, req *palacev1.GetSystemRoleRequest) (*common.SystemRoleItem, error) {
	roleDo, err := s.systemBiz.GetRole(ctx, req.GetRoleId())
	if err != nil {
		return nil, err
	}
	return build.ToSystemRoleItem(roleDo), nil
}

func (s *SystemService) GetSystemRoles(ctx context.Context, req *palacev1.GetSystemRolesRequest) (*palacev1.GetSystemRolesReply, error) {
	params := build.ToListRoleRequest(req)
	roleReply, err := s.systemBiz.GetRoles(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetSystemRolesReply{
		Items:      build.ToSystemRoleItems(roleReply.Items),
		Pagination: build.ToPaginationReply(roleReply.PaginationReply),
	}, nil
}

func (s *SystemService) SaveRole(ctx context.Context, req *palacev1.SaveRoleRequest) (*common.EmptyReply, error) {
	params := build.ToSaveRoleRequest(req)
	if err := s.systemBiz.SaveRole(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存角色成功"}, nil
}

func (s *SystemService) UpdateRoleStatus(ctx context.Context, req *palacev1.UpdateRoleStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateRoleStatusReq{
		RoleID: req.GetRoleId(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}
	if err := s.systemBiz.UpdateRoleStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新角色状态成功"}, nil
}

func (s *SystemService) UpdateUserRoles(ctx context.Context, req *palacev1.UpdateUserRolesRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateUserRolesReq{
		RoleIDs: req.GetRoleIds(),
		UserID:  req.GetUserId(),
	}
	if err := s.systemBiz.UpdateUserRoles(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新用户角色成功"}, nil
}

func (s *SystemService) UpdateRoleUsers(ctx context.Context, req *palacev1.UpdateRoleUsersRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateRoleUsersReq{
		RoleID:  req.GetRoleId(),
		UserIDs: req.GetUserIds(),
	}
	if err := s.systemBiz.UpdateRoleUsers(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新角色用户成功"}, nil
}

func (s *SystemService) GetTeamAuditList(ctx context.Context, req *palacev1.GetTeamAuditListRequest) (*palacev1.GetTeamAuditListReply, error) {
	params := build.ToTeamAuditListRequest(req)
	teamAuditReply, err := s.systemBiz.GetTeamAuditList(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetTeamAuditListReply{
		Items:      build.ToTeamAuditItems(teamAuditReply.Items),
		Pagination: build.ToPaginationReply(teamAuditReply.PaginationReply),
	}, nil
}

func (s *SystemService) UpdateTeamAuditStatus(ctx context.Context, req *palacev1.UpdateTeamAuditStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateTeamAuditStatusReq{
		AuditID: req.GetAuditId(),
		Status:  vobj.StatusAudit(req.GetStatus()),
		Reason:  req.GetReason(),
	}
	if err := s.systemBiz.UpdateTeamAuditStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新团队审核状态成功"}, nil
}

func (s *SystemService) OperateLogList(ctx context.Context, req *palacev1.OperateLogListRequest) (*palacev1.OperateLogListReply, error) {
	params := build.ToOperateLogListRequest(req)
	operateLogReply, err := s.systemBiz.OperateLogList(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.OperateLogListReply{
		Items:      build.ToOperateLogItems(operateLogReply.Items),
		Pagination: build.ToPaginationReply(operateLogReply.PaginationReply),
	}, nil
}
