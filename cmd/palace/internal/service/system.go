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

func NewSystemService(userBiz *biz.UserBiz, messageBiz *biz.Message) *SystemService {
	return &SystemService{
		userBiz:    userBiz,
		messageBiz: messageBiz,
	}
}

type SystemService struct {
	palacev1.UnimplementedSystemServer
	userBiz    *biz.UserBiz
	messageBiz *biz.Message
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
	return &common.EmptyReply{}, nil
}

func (s *SystemService) GetUser(ctx context.Context, req *palacev1.GetUserRequest) (*palacev1.GetUserReply, error) {
	return &palacev1.GetUserReply{}, nil
}

func (s *SystemService) GetUserList(ctx context.Context, req *palacev1.GetUserListRequest) (*palacev1.GetUserListReply, error) {
	return &palacev1.GetUserListReply{}, nil
}

func (s *SystemService) GetTeamList(ctx context.Context, req *palacev1.GetTeamListRequest) (*palacev1.GetTeamListReply, error) {
	return &palacev1.GetTeamListReply{}, nil
}

func (s *SystemService) GetSystemRole(ctx context.Context, req *palacev1.GetSystemRoleRequest) (*palacev1.GetSystemRoleReply, error) {
	return &palacev1.GetSystemRoleReply{}, nil
}

func (s *SystemService) SaveRole(ctx context.Context, req *palacev1.SaveRoleRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) UpdateRoleStatus(ctx context.Context, req *palacev1.UpdateRoleStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) UpdateUserRoles(ctx context.Context, req *palacev1.UpdateUserRolesRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) UpdateRoleUsers(ctx context.Context, req *palacev1.UpdateRoleUsersRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) GetTeamAuditList(ctx context.Context, req *palacev1.GetTeamAuditListRequest) (*palacev1.GetTeamAuditListReply, error) {
	return &palacev1.GetTeamAuditListReply{}, nil
}

func (s *SystemService) UpdateTeamAuditStatus(ctx context.Context, req *palacev1.UpdateTeamAuditStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) OperateLogList(ctx context.Context, req *palacev1.OperateLogListRequest) (*palacev1.OperateLogListReply, error) {
	return &palacev1.OperateLogListReply{}, nil
}
