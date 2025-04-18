package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

type ListRoleRequest interface {
	GetPagination() *common.PaginationRequest
	GetStatus() common.GlobalStatus
	GetKeyword() string
}

func ToListRoleRequest(req ListRoleRequest) *bo.ListRoleReq {
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
		Status:    common.GlobalStatus(role.GetStatus()),
		Resources: nil,
		Members:   nil,
		CreatedAt: role.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: role.GetUpdatedAt().Format(time.RFC3339),
		Creator:   UserToUserBaseItemProto(role.GetCreator()),
	}
}

func ToTeamRoleItems(roles []do.TeamRole) []*common.TeamRoleItem {
	return slices.Map(roles, ToTeamRoleItem)
}

func ToSystemRoleItem(role do.Role) *common.SystemRoleItem {
	if validate.IsNil(role) {
		return nil
	}
	return &common.SystemRoleItem{
		Id:        role.GetID(),
		Name:      role.GetName(),
		Remark:    role.GetRemark(),
		Status:    common.GlobalStatus(role.GetStatus()),
		CreatedAt: role.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: role.GetUpdatedAt().Format(time.RFC3339),
		Resources: nil,
		Users:     nil,
		Creator:   UserToUserBaseItemProto(role.GetCreator()),
	}
}

func ToSystemRoleItems(roles []do.Role) []*common.SystemRoleItem {
	return slices.Map(roles, ToSystemRoleItem)
}

type SaveTeamRoleRequest interface {
	GetRoleId() uint32
	GetName() string
	GetRemark() string
	GetMenuIds() []uint32
}

func ToSaveTeamRoleRequest(req SaveTeamRoleRequest) *bo.SaveTeamRoleReq {
	return &bo.SaveTeamRoleReq{
		ID:      req.GetRoleId(),
		Name:    req.GetName(),
		Remark:  req.GetRemark(),
		MenuIds: req.GetMenuIds(),
	}
}

func ToSaveRoleRequest(req SaveTeamRoleRequest) *bo.SaveRoleReq {
	return &bo.SaveRoleReq{
		ID:      req.GetRoleId(),
		Name:    req.GetName(),
		Remark:  req.GetRemark(),
		MenuIds: req.GetMenuIds(),
	}
}
