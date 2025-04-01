package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// NewPermissionBiz create a new permission biz
func NewPermissionBiz(
	resourceRepo repository.Resource,
	userRepo repository.User,
	teamRepo repository.Team,
	memberRepo repository.Member,
	logger log.Logger,
) *PermissionBiz {
	baseHandler := NewBasePermissionHandler()
	// build permission chain
	permissionChain := []PermissionHandler{
		baseHandler.UserHandler(userRepo.FindByID),
		baseHandler.OperationHandler(),
		baseHandler.ResourceHandler(resourceRepo.GetResourceByOperation),
		baseHandler.AllowCheckHandler(),
		baseHandler.SystemPositionHandler(),
		baseHandler.SystemAdminCheckHandler(),
		baseHandler.SystemRBACHandler(checkSystemRBAC),
		baseHandler.TeamIDHandler(teamRepo.FindByID),
		baseHandler.TeamMemberHandler(memberRepo.FindByUserID),
		baseHandler.TeamPositionHandler(),
		baseHandler.TeamAdminCheckHandler(),
		baseHandler.TeamRBACHandler(checkTeamRBAC),
	}
	return &PermissionBiz{
		helper:          log.NewHelper(log.With(logger, "module", "biz.permission")),
		permissionChain: permissionChain,
	}
}

type PermissionBiz struct {
	permissionChain []PermissionHandler // add permission chain
	helper          *log.Helper
}

func (a *PermissionBiz) VerifyPermission(ctx context.Context) error {
	pctx := &PermissionContext{}
	for _, handler := range a.permissionChain {
		stop, err := handler.Handle(ctx, pctx)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
	}
	return nil
}

// PermissionContext permission check context
type PermissionContext struct {
	Operation      string
	Resource       *system.Resource
	User           *system.User
	Team           *system.Team
	SystemPosition vobj.Role
	TeamPosition   vobj.Role
	TeamMember     *system.TeamMember
}

// PermissionHandler permission handler interface
type PermissionHandler interface {
	Handle(ctx context.Context, pctx *PermissionContext) (stop bool, err error)
}

// PermissionHandlerFunc permission handler function type
type PermissionHandlerFunc func(ctx context.Context, pctx *PermissionContext) (stop bool, err error)

func (f PermissionHandlerFunc) Handle(ctx context.Context, pctx *PermissionContext) (bool, error) {
	return f(ctx, pctx)
}

// base permission handler implementation
type basePermissionHandler struct{}

func NewBasePermissionHandler() *basePermissionHandler {
	return &basePermissionHandler{}
}

// OperationHandler operation check
func (h *basePermissionHandler) OperationHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		operation, ok := permission.GetOperationByContext(ctx)
		if !ok {
			return true, merr.ErrorBadRequest("operation is invalid")
		}
		pctx.Operation = operation
		return false, nil
	})
}

// ResourceHandler resource check
func (h *basePermissionHandler) ResourceHandler(getResourceByOperation func(ctx context.Context, operation string) (*system.Resource, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		resource, err := getResourceByOperation(ctx, pctx.Operation)
		if err != nil {
			return true, err
		}
		if !resource.Status.IsEnabled() {
			return true, merr.ErrorPermissionDenied("permission denied")
		}
		pctx.Resource = resource
		return false, nil
	})
}

// UserHandler user check
func (h *basePermissionHandler) UserHandler(findUserByID func(ctx context.Context, userID uint32) (*system.User, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		userID, ok := permission.GetUserIDByContext(ctx)
		if !ok {
			return true, merr.ErrorBadRequest("user id is invalid")
		}
		user, err := findUserByID(ctx, userID)
		if err != nil {
			return true, err
		}
		if !user.Status.IsNormal() {
			return true, merr.ErrorUserForbidden("user forbidden")
		}
		pctx.User = user
		return false, nil
	})
}

// AllowCheckHandler allow check
func (h *basePermissionHandler) AllowCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.Resource.Allow.IsNone() || pctx.Resource.Allow.IsUser() {
			return true, nil // satisfy condition directly pass
		}
		return false, nil
	})
}

// SystemPositionHandler system position check
func (h *basePermissionHandler) SystemPositionHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		sysPosition, ok := permission.GetSysPositionByContext(ctx)
		if !ok {
			pctx.SystemPosition = pctx.User.Position
			return false, nil
		}
		if pctx.User.Position.GTE(sysPosition) {
			pctx.SystemPosition = sysPosition
			return false, nil
		}
		return true, merr.ErrorPermissionDenied("your current role [%s] is not allowed to access this resource", sysPosition)
	})
}

// SystemAdminCheckHandler system admin check
func (h *basePermissionHandler) SystemAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.SystemPosition.IsAdminOrSuperAdmin() {
			return true, nil // 管理员直接通过
		}
		return false, nil
	})
}

// SystemRBACHandler system rbac check
func (h *basePermissionHandler) SystemRBACHandler(checkSystemRBAC func(ctx context.Context, user *system.User, resource *system.Resource) (bool, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		ok, err := checkSystemRBAC(ctx, pctx.User, pctx.Resource)
		if err != nil {
			return false, err
		}
		return ok, nil
	})
}

// TeamIDHandler team id check
func (h *basePermissionHandler) TeamIDHandler(findTeamByID func(ctx context.Context, teamID uint32) (*system.Team, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		teamID, ok := permission.GetTeamIDByContext(ctx)
		if !ok {
			return true, merr.ErrorPermissionDenied("team id is invalid")
		}
		team, err := findTeamByID(ctx, teamID)
		if err != nil {
			return true, err
		}
		if !team.Status.IsNormal() {
			return true, merr.ErrorPermissionDenied("team is invalid")
		}
		pctx.Team = team
		return false, nil
	})
}

// TeamMemberHandler team member check
func (h *basePermissionHandler) TeamMemberHandler(findTeamMemberByUserID func(ctx context.Context, userID uint32) (*system.TeamMember, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		member, err := findTeamMemberByUserID(ctx, pctx.User.ID)
		if err != nil {
			return true, err
		}
		if !member.Status.IsNormal() {
			return true, merr.ErrorPermissionDenied("team member is invalid [%s]", member.Status)
		}
		if pctx.Team.ID != member.TeamID {
			return true, merr.ErrorPermissionDenied("team id is invalid")
		}
		pctx.TeamMember = member
		return false, nil
	})
}

// TeamPositionHandler team position check
func (h *basePermissionHandler) TeamPositionHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		teamPosition, ok := permission.GetTeamPositionByContext(ctx)
		if !ok {
			pctx.TeamPosition = pctx.TeamMember.Position
			return false, nil
		}
		if pctx.TeamMember.Position.GTE(teamPosition) {
			pctx.TeamPosition = teamPosition
			return false, nil
		}
		return true, merr.ErrorPermissionDenied("your current team role [%s] is not allowed to access this resource", teamPosition)
	})
}

// TeamAdminCheckHandler team admin check
func (h *basePermissionHandler) TeamAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.TeamPosition.IsAdminOrSuperAdmin() {
			return true, nil // team admin directly pass
		}
		return false, nil
	})
}

// TeamRBACHandler team rbac check
func (h *basePermissionHandler) TeamRBACHandler(checkTeamRBAC func(ctx context.Context, member *system.TeamMember, resource *system.Resource) (bool, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		ok, err := checkTeamRBAC(ctx, pctx.TeamMember, pctx.Resource)
		if err != nil {
			return false, err
		}
		return ok, nil
	})
}

func checkSystemRBAC(ctx context.Context, user *system.User, resource *system.Resource) (bool, error) {
	if !resource.Allow.IsSystemRBAC() {
		return false, nil
	}
	sysRoleID, ok := permission.GetSysRoleIDByContext(ctx)
	if ok {
		// check if the role exists and has this API permission
		systemRoleDo, ok := validate.SliceFindByValue(user.Roles, sysRoleID, func(role *system.SysRole) uint32 {
			return role.ID
		})
		if !ok {
			return false, merr.ErrorPermissionDenied("user role is invalid.")
		}
		if !systemRoleDo.Status.IsNormal() {
			return false, merr.ErrorPermissionDenied("role is invalid [%s]", systemRoleDo.Status)
		}
		_, ok = validate.SliceFindByValue(systemRoleDo.Resources, resource.ID, func(role *system.Resource) uint32 {
			return role.ID
		})
		if ok {
			return true, nil
		}
		return false, merr.ErrorPermissionDenied("user role resource is invalid.")
	}
	resources := make([]*system.Resource, 0, len(user.Roles)*10)
	for _, role := range user.Roles {
		if role.Status.IsNormal() {
			resources = append(resources, role.Resources...)
		}
	}
	_, ok = validate.SliceFindByValue(resources, resource.ID, func(role *system.Resource) uint32 {
		return role.ID
	})
	if ok {
		return true, nil
	}
	return false, merr.ErrorPermissionDenied("user role resource is invalid.")
}

func checkTeamRBAC(ctx context.Context, member *system.TeamMember, resource *system.Resource) (bool, error) {
	if !resource.Allow.IsTeamRBAC() {
		return false, nil
	}
	teamRoleID, ok := permission.GetTeamRoleIDByContext(ctx)
	if ok {
		teamRoleDo, ok := validate.SliceFindByValue(member.Roles, teamRoleID, func(role *system.TeamRole) uint32 {
			return role.ID
		})
		if !ok {
			return false, merr.ErrorPermissionDenied("team role is invalid")
		}
		if !teamRoleDo.Status.IsNormal() {
			return false, merr.ErrorPermissionDenied("team role is invalid [%s]", teamRoleDo.Status)
		}
		_, ok = validate.SliceFindByValue(teamRoleDo.Resources, resource.ID, func(role *system.Resource) uint32 {
			return role.ID
		})
		if ok {
			return true, nil
		}
		return false, merr.ErrorPermissionDenied("team role resource is invalid.")
	}
	resources := make([]*system.Resource, 0, len(member.Roles)*10)
	for _, role := range member.Roles {
		if role.Status.IsNormal() {
			resources = append(resources, role.Resources...)
		}
	}
	_, ok = validate.SliceFindByValue(resources, resource.ID, func(role *system.Resource) uint32 {
		return role.ID
	})
	if ok {
		return true, nil
	}
	return false, merr.ErrorPermissionDenied("team role resource is invalid.")
}

// VerifyNewPermission 验证新权限
func (a *PermissionBiz) VerifyNewPermission(ctx context.Context, permissionReq *bo.VerifyNewPermission) error {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorBadRequest("user id is invalid")
	}

	// 1. 获取用户信息并验证
	user, err := a.findUser(ctx, userID)
	if err != nil {
		return err
	}

	if !user.Status.IsNormal() {
		return merr.ErrorUserForbidden("user forbidden")
	}

	// 1.1 验证系统角色权限
	if permissionReq.SystemPosition.IsAdminOrSuperAdmin() && !user.Position.IsAdminOrSuperAdmin() {
		return merr.ErrorPermissionDenied("cannot switch to admin role without admin permission")
	}

	// 1.2 验证当前用户是否有权限切换到目标系统角色
	if !user.Position.GTE(permissionReq.SystemPosition) {
		return merr.ErrorPermissionDenied("your role [%s] is not allowed to switch to [%s]",
			user.Position, permissionReq.SystemPosition)
	}

	// 1.3 验证系统角色ID是否有效
	if permissionReq.SystemRoleID > 0 {
		systemRoleDo, ok := validate.SliceFindByValue(user.Roles, permissionReq.SystemRoleID, func(role *system.SysRole) uint32 {
			return role.ID
		})
		if !ok {
			return merr.ErrorPermissionDenied("system role is invalid")
		}
		if !systemRoleDo.Status.IsNormal() {
			return merr.ErrorPermissionDenied("system role is invalid status: [%s]", systemRoleDo.Status)
		}
	}

	// 2. 如果有团队ID，验证团队权限
	if permissionReq.TeamID > 0 {
		// 2.1 验证团队是否存在且正常
		team, err := a.findTeam(ctx, permissionReq.TeamID)
		if err != nil {
			return err
		}
		if !team.Status.IsNormal() {
			return merr.ErrorPermissionDenied("team is invalid")
		}

		// 2.2 验证用户是否是团队成员
		member, err := a.findMember(ctx, userID, permissionReq.TeamID)
		if err != nil {
			return err
		}
		if !member.Status.IsNormal() {
			return merr.ErrorPermissionDenied("team member is invalid [%s]", member.Status)
		}

		// 2.3 验证成员是否属于当前团队
		if team.ID != member.TeamID {
			return merr.ErrorPermissionDenied("team id is invalid")
		}

		// 2.4 验证团队角色权限
		if permissionReq.TeamPosition.IsAdminOrSuperAdmin() && !member.Position.IsAdminOrSuperAdmin() {
			return merr.ErrorPermissionDenied("cannot switch to team admin role without team admin permission")
		}

		// 2.5 验证当前用户是否有权限切换到目标团队角色
		if !member.Position.GTE(permissionReq.TeamPosition) {
			return merr.ErrorPermissionDenied("your team role [%s] is not allowed to switch to [%s]",
				member.Position, permissionReq.TeamPosition)
		}

		// 2.6 验证团队角色ID是否有效
		if permissionReq.TeamRoleID > 0 {
			teamRoleDo, ok := validate.SliceFindByValue(member.Roles, permissionReq.TeamRoleID, func(role *system.TeamRole) uint32 {
				return role.ID
			})
			if !ok {
				return merr.ErrorPermissionDenied("team role is invalid")
			}
			if !teamRoleDo.Status.IsNormal() {
				return merr.ErrorPermissionDenied("team role is invalid status: [%s]", teamRoleDo.Status)
			}
		}
	}

	return nil
}

// 查询用户
func (a *PermissionBiz) findUser(ctx context.Context, userID uint32) (*system.User, error) {
	for _, handler := range a.permissionChain {
		if userHandler, ok := handler.(PermissionHandlerFunc); ok {
			pctx := &PermissionContext{
				User: &system.User{},
			}
			pctx.User.ID = userID
			_, err := userHandler.Handle(ctx, pctx)
			if err == nil && pctx.User != nil && pctx.User.ID > 0 {
				return pctx.User, nil
			}
		}
	}
	return nil, merr.ErrorUserNotFound("user not found")
}

// 查询团队
func (a *PermissionBiz) findTeam(ctx context.Context, teamID uint32) (*system.Team, error) {
	for _, handler := range a.permissionChain {
		if teamHandler, ok := handler.(PermissionHandlerFunc); ok {
			pctx := &PermissionContext{
				Team: &system.Team{},
			}
			pctx.Team.ID = teamID
			_, err := teamHandler.Handle(ctx, pctx)
			if err == nil && pctx.Team != nil && pctx.Team.ID > 0 {
				return pctx.Team, nil
			}
		}
	}
	return nil, merr.ErrorPermissionDenied("team not found")
}

// 查询团队成员
func (a *PermissionBiz) findMember(ctx context.Context, userID uint32, teamID uint32) (*system.TeamMember, error) {
	for _, handler := range a.permissionChain {
		if memberHandler, ok := handler.(PermissionHandlerFunc); ok {
			pctx := &PermissionContext{
				User: &system.User{},
				Team: &system.Team{},
			}
			pctx.User.ID = userID
			pctx.Team.ID = teamID
			_, err := memberHandler.Handle(ctx, pctx)
			if err == nil && pctx.TeamMember != nil {
				return pctx.TeamMember, nil
			}
		}
	}
	return nil, merr.ErrorPermissionDenied("team member not found")
}
