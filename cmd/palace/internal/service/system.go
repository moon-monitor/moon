package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type SystemService struct {
	palacev1.UnimplementedSystemServer
}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) UpdateUser(ctx context.Context, req *palacev1.UpdateUserRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) UpdateUserStatus(ctx context.Context, req *palacev1.UpdateUserStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SystemService) ResetUserPassword(ctx context.Context, req *palacev1.ResetUserPasswordRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
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
