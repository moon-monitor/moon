package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToListRoleRequest(req *palace.GetTeamRolesRequest) *bo.ListRoleReq {
	return &bo.ListRoleReq{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Status:            vobj.GlobalStatus(req.GetStatus()),
		Keyword:           req.GetKeyword(),
	}
}

func ToTeamRoleItem(role do.TeamRole) *common.TeamRoleItem {
	if validate.IsNil(role) {
		return nil
	}
	return &common.TeamRoleItem{
		Id:        role.GetID(),
		Name:      role.GetName(),
		Remark:    role.GetRemark(),
		Status:    common.RoleStatus(role.GetStatus()),
		Resources: nil,
		Members:   nil,
		CreatedAt: role.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: role.GetUpdatedAt().Format(time.RFC3339),
	}
}

func ToListTeamRoleReply(reply *bo.ListRoleReply) *palace.GetTeamRolesReply {
	if validate.IsNil(reply) {
		return &palace.GetTeamRolesReply{}
	}
	return &palace.GetTeamRolesReply{
		Items:      slices.Map(reply.Items, ToTeamRoleItem),
		Pagination: ToPaginationReplyProto(reply.PaginationReply),
	}
}

func ToSaveTeamRoleRequest(req *palace.SaveTeamRoleRequest) *bo.SaveTeamRoleReq {
	return &bo.SaveTeamRoleReq{
		ID:      req.GetRoleID(),
		Name:    req.GetName(),
		Remark:  req.GetRemark(),
		MenuIds: req.GetResourceIDs(),
	}
}
