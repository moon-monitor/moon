package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewPermissionBiz(
	resourceRepo repository.Resource,
	userRepo repository.User,
	teamRepo repository.Team,
	memberRepo repository.Member,
	logger log.Logger,
) *PermissionBiz {
	baseHandler := NewBasePermissionHandler()
	// 构建权限校验链
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
	permissionChain []PermissionHandler // 新增权限校验链
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

// PermissionContext 权限校验上下文
type PermissionContext struct {
	Operation      string
	Resource       *system.Resource
	User           *system.User
	Team           *system.Team
	SystemPosition vobj.Role
	TeamPosition   vobj.Role
	TeamMember     *system.TeamMember
}

// PermissionHandler 权限处理器接口
type PermissionHandler interface {
	Handle(ctx context.Context, pctx *PermissionContext) (stop bool, err error)
}

// PermissionHandlerFunc 处理器函数类型
type PermissionHandlerFunc func(ctx context.Context, pctx *PermissionContext) (stop bool, err error)

func (f PermissionHandlerFunc) Handle(ctx context.Context, pctx *PermissionContext) (bool, error) {
	return f(ctx, pctx)
}

// 基础处理器实现
type basePermissionHandler struct{}

func NewBasePermissionHandler() *basePermissionHandler {
	return &basePermissionHandler{}
}

// OperationHandler 操作校验
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

// ResourceHandler 资源校验
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

// UserHandler 用户校验
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

// AllowCheckHandler 白名单检查
func (h *basePermissionHandler) AllowCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.Resource.Allow.IsNone() || pctx.Resource.Allow.IsUser() {
			return true, nil // 满足条件直接通过
		}
		return false, nil
	})
}

// SystemPositionHandler 系统职位校验
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
		return true, merr.ErrorPermissionDenied("Your current role [%s] is not allowed to access this resource", sysPosition)
	})
}

// SystemAdminCheckHandler 系统管理员检查
func (h *basePermissionHandler) SystemAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.SystemPosition.IsAdminOrSuperAdmin() {
			return true, nil // 管理员直接通过
		}
		return false, nil
	})
}

// SystemRBACHandler 系统RBAC校验
func (h *basePermissionHandler) SystemRBACHandler(checkSystemRBAC func(ctx context.Context, user *system.User, resource *system.Resource) (bool, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		ok, err := checkSystemRBAC(ctx, pctx.User, pctx.Resource)
		if err != nil {
			return false, err
		}
		return ok, nil
	})
}

// TeamIDHandler 团队ID校验
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

// TeamMemberHandler 团队成员校验
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

// TeamPositionHandler 团队职位校验
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
		return true, merr.ErrorPermissionDenied("Your current team role [%s] is not allowed to access this resource", teamPosition)
	})
}

// TeamAdminCheckHandler 团队管理员检查
func (h *basePermissionHandler) TeamAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pctx *PermissionContext) (bool, error) {
		if pctx.TeamPosition.IsAdminOrSuperAdmin() {
			return true, nil // 团队管理员直接通过
		}
		return false, nil
	})
}

// TeamRBACHandler 团队RBAC校验
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
		// 判断角色是否存在，且该角色具备次API权限
		systemRoleDo, ok := validate.SliceFindByValue(user.Roles, sysRoleID, func(role *system.SysRole) uint32 {
			return role.ID
		})
		if !ok {
			return false, merr.ErrorPermissionDenied("User role is invalid.")
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
		return false, merr.ErrorPermissionDenied("User role resource is invalid.")
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
	return false, merr.ErrorPermissionDenied("User role resource is invalid.")
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
