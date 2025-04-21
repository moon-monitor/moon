package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToRole(roleDo do.Role) *system.Role {
	if validate.IsNil(roleDo) {
		return nil
	}
	role, ok := roleDo.(*system.Role)
	if ok {
		return role
	}
	return &system.Role{
		CreatorModel: ToCreatorModel(roleDo),
		Name:         role.GetName(),
		Remark:       role.GetRemark(),
		Status:       role.GetStatus(),
		Users:        ToUsers(role.GetUsers()),
		Menus:        nil,
	}
}

func ToRoles(roles []do.Role) []*system.Role {
	return slices.Map(roles, ToRole)
}

func ToTeamRole(roleDo do.TeamRole) *system.TeamRole {
	if validate.IsNil(roleDo) {
		return nil
	}
	role, ok := roleDo.(*system.TeamRole)
	if ok {
		return role
	}
	return &system.TeamRole{
		TeamModel: ToTeamModel(roleDo),
		Name:      role.GetName(),
		Remark:    role.GetRemark(),
		Status:    role.GetStatus(),
		Members:   nil,
		Menus:     nil,
	}
}

func ToTeamRoles(roles []do.TeamRole) []*system.TeamRole {
	return slices.Map(roles, ToTeamRole)
}
