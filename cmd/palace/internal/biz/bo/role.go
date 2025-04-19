package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

type Role interface {
	GetID() uint32
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetMenus() []do.Menu
	GetMenuIds() []uint32
}

type SaveTeamRoleReq struct {
	teamRole do.TeamRole
	menus    []do.Menu
	ID       uint32
	Name     string
	Remark   string
	MenuIds  []uint32
}

func (r *SaveTeamRoleReq) GetID() uint32 {
	if r == nil {
		return 0
	}
	if r.teamRole == nil {
		return r.ID
	}
	return r.teamRole.GetID()
}

func (r *SaveTeamRoleReq) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *SaveTeamRoleReq) GetRemark() string {
	if r == nil {
		return ""
	}
	return r.Remark
}

func (r *SaveTeamRoleReq) GetStatus() vobj.GlobalStatus {
	if r == nil {
		return vobj.GlobalStatusUnknown
	}
	if validate.IsNil(r.teamRole) {
		return vobj.GlobalStatusEnable
	}
	return r.teamRole.GetStatus()
}

func (r *SaveTeamRoleReq) GetMenus() []do.Menu {
	if r == nil {
		return nil
	}
	return r.menus
}

func (r *SaveTeamRoleReq) GetMenuIds() []uint32 {
	if r == nil {
		return nil
	}
	return r.MenuIds
}

func (r *SaveTeamRoleReq) WithMenus(menus []do.Menu) Role {
	r.menus = menus
	return r
}

func (r *SaveTeamRoleReq) WithRole(role do.TeamRole) Role {
	r.teamRole = role
	return r
}

type SaveRoleReq struct {
	role    do.Role
	menus   []do.Menu
	ID      uint32
	Name    string
	Remark  string
	MenuIds []uint32
}

func (r *SaveRoleReq) GetID() uint32 {
	if r == nil {
		return 0
	}
	if r.role == nil {
		return r.ID
	}
	return r.role.GetID()
}

func (r *SaveRoleReq) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *SaveRoleReq) GetRemark() string {
	if r == nil {
		return ""
	}
	return r.Remark
}

func (r *SaveRoleReq) GetStatus() vobj.GlobalStatus {
	if r == nil {
		return vobj.GlobalStatusUnknown
	}
	if validate.IsNil(r.role) {
		return vobj.GlobalStatusEnable
	}
	return r.role.GetStatus()
}

func (r *SaveRoleReq) GetMenus() []do.Menu {
	if r == nil {
		return nil
	}
	return r.menus
}

func (r *SaveRoleReq) GetMenuIds() []uint32 {
	if r == nil {
		return nil
	}
	return r.MenuIds
}

func (r *SaveRoleReq) WithRole(role do.Role) Role {
	r.role = role
	return r
}

func (r *SaveRoleReq) WithMenus(menus []do.Menu) Role {
	r.menus = menus
	return r
}

type ListRoleReq struct {
	*PaginationRequest
	Status  vobj.GlobalStatus `json:"status"`
	Keyword string            `json:"keyword"`
}

func (r *ListRoleReq) ToListTeamRoleReply(roles []*team.Role) *ListTeamRoleReply {
	return &ListTeamRoleReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(roles, func(role *team.Role) do.TeamRole { return role }),
	}
}

func (r *ListRoleReq) ToListRoleReply(roles []*system.Role) *ListRoleReply {
	return &ListRoleReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(roles, func(role *system.Role) do.Role { return role }),
	}
}

type ListTeamRoleReply = ListReply[do.TeamRole]

type ListRoleReply = ListReply[do.Role]

type UpdateRoleStatusReq struct {
	RoleID uint32            `json:"roleId"`
	Status vobj.GlobalStatus `json:"status"`
}

type UpdateRoleUsers interface {
	GetRoleID() uint32
	GetUserIds() []uint32
	GetUsers() []do.User
}

type UpdateRoleUsersReq struct {
	RoleID  uint32   `json:"roleId"`
	UserIDs []uint32 `json:"userIds"`
	users   []do.User
}

func (r *UpdateRoleUsersReq) GetRoleID() uint32 {
	if r == nil {
		return 0
	}
	return r.RoleID
}

func (r *UpdateRoleUsersReq) GetUserIds() []uint32 {
	if r == nil {
		return nil
	}
	return r.UserIDs
}

func (r *UpdateRoleUsersReq) GetUsers() []do.User {
	if r == nil {
		return nil
	}
	return r.users
}

func (r *UpdateRoleUsersReq) WithUsers(users []do.User) UpdateRoleUsers {
	r.users = users
	return r
}
