package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

// NewPermissionBiz create a new permission biz
func NewPermissionBiz(
	resourceRepo repository.Resource,
	userRepo repository.User,
	teamRepo repository.Team,
	memberRepo repository.Member,
	logger log.Logger,
) *PermissionBiz {
	baseHandler := &basePermissionHandler{}
	// build permission chain
	permissionChain := []PermissionHandler{
		baseHandler.UserHandler(userRepo.FindByID),
		baseHandler.OperationHandler(),
		baseHandler.ResourceHandler(resourceRepo.GetResourceByOperation),
		baseHandler.AllowCheckHandler(),
		baseHandler.SystemAdminCheckHandler(),
		baseHandler.SystemRBACHandler(checkSystemRBAC),
		baseHandler.TeamIDHandler(teamRepo.FindByID),
		baseHandler.TeamMemberHandler(memberRepo.FindByUserID),
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
	pCtx := &PermissionContext{}
	for _, handler := range a.permissionChain {
		stop, err := handler.Handle(ctx, pCtx)
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
	Handle(ctx context.Context, pCtx *PermissionContext) (stop bool, err error)
}

// PermissionHandlerFunc permission handler function type
type PermissionHandlerFunc func(ctx context.Context, pCtx *PermissionContext) (stop bool, err error)

func (f PermissionHandlerFunc) Handle(ctx context.Context, pCtx *PermissionContext) (bool, error) {
	return f(ctx, pCtx)
}

// base permission handler implementation
type basePermissionHandler struct{}

// OperationHandler operation check
func (h *basePermissionHandler) OperationHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		operation, ok := permission.GetOperationByContext(ctx)
		if !ok {
			return true, merr.ErrorBadRequest("operation is invalid")
		}
		pCtx.Operation = operation
		return false, nil
	})
}

// ResourceHandler resourceRepo check
func (h *basePermissionHandler) ResourceHandler(getResourceByOperation func(ctx context.Context, operation string) (*system.Resource, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		resource, err := getResourceByOperation(ctx, pCtx.Operation)
		if err != nil {
			return true, err
		}
		if !resource.Status.IsEnable() {
			return true, merr.ErrorPermissionDenied("permission denied")
		}
		pCtx.Resource = resource
		return false, nil
	})
}

// UserHandler user check
func (h *basePermissionHandler) UserHandler(findUserByID func(ctx context.Context, userID uint32) (*system.User, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
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
		pCtx.User = user
		pCtx.SystemPosition = user.Position
		return false, nil
	})
}

// AllowCheckHandler allow check
func (h *basePermissionHandler) AllowCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		if pCtx.Resource.Allow.IsNone() || pCtx.Resource.Allow.IsUser() {
			return true, nil // satisfy condition directly pass
		}
		return false, nil
	})
}

// SystemAdminCheckHandler system admin check
func (h *basePermissionHandler) SystemAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		if pCtx.SystemPosition.IsAdminOrSuperAdmin() {
			return true, nil // 管理员直接通过
		}
		return false, nil
	})
}

// SystemRBACHandler system rbac check
func (h *basePermissionHandler) SystemRBACHandler(checkSystemRBAC func(ctx context.Context, user *system.User, resource *system.Resource) (bool, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		ok, err := checkSystemRBAC(ctx, pCtx.User, pCtx.Resource)
		if err != nil {
			return false, err
		}
		return ok, nil
	})
}

// TeamIDHandler team id check
func (h *basePermissionHandler) TeamIDHandler(findTeamByID func(ctx context.Context, teamID uint32) (*system.Team, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
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
		pCtx.Team = team
		return false, nil
	})
}

// TeamMemberHandler team member check
func (h *basePermissionHandler) TeamMemberHandler(findTeamMemberByUserID func(ctx context.Context, userID uint32) (*system.TeamMember, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		member, err := findTeamMemberByUserID(ctx, pCtx.User.ID)
		if err != nil {
			return true, err
		}
		if !member.Status.IsNormal() {
			return true, merr.ErrorPermissionDenied("team member is invalid [%s]", member.Status)
		}
		if pCtx.Team.ID != member.TeamID {
			return true, merr.ErrorPermissionDenied("team id is invalid")
		}
		pCtx.TeamMember = member
		pCtx.TeamPosition = member.Position
		return false, nil
	})
}

// TeamAdminCheckHandler team admin check
func (h *basePermissionHandler) TeamAdminCheckHandler() PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		if pCtx.TeamPosition.IsAdminOrSuperAdmin() {
			return true, nil // team admin directly pass
		}
		return false, nil
	})
}

// TeamRBACHandler team rbac check
func (h *basePermissionHandler) TeamRBACHandler(checkTeamRBAC func(ctx context.Context, member *system.TeamMember, resource *system.Resource) (bool, error)) PermissionHandler {
	return PermissionHandlerFunc(func(ctx context.Context, pCtx *PermissionContext) (bool, error) {
		ok, err := checkTeamRBAC(ctx, pCtx.TeamMember, pCtx.Resource)
		if err != nil {
			return false, err
		}
		return ok, nil
	})
}

func checkSystemRBAC(_ context.Context, user *system.User, resource *system.Resource) (bool, error) {
	if !resource.Allow.IsSystemRBAC() {
		return false, nil
	}
	resources := make([]*system.Resource, 0, len(user.Roles)*10)
	for _, role := range user.Roles {
		if role.Status.IsEnable() {
			for _, menu := range role.Menus {
				if !menu.Status.IsEnable() {
					continue
				}
				resources = append(resources, menu.Resources...)
			}
		}
	}
	_, ok := slices.FindByValue(resources, resource.ID, func(role *system.Resource) uint32 { return role.ID })
	if ok {
		return true, nil
	}
	return false, merr.ErrorPermissionDenied("user role resourceRepo is invalid.")
}

func checkTeamRBAC(_ context.Context, member *system.TeamMember, resource *system.Resource) (bool, error) {
	if !resource.Allow.IsTeamRBAC() {
		return false, nil
	}
	resources := make([]*system.Resource, 0, len(member.Roles)*10)
	for _, role := range member.Roles {
		if role.Status.IsEnable() {
			for _, menu := range role.Menus {
				if !menu.Status.IsEnable() {
					continue
				}
				resources = append(resources, menu.Resources...)
			}
		}
	}
	_, ok := slices.FindByValue(resources, resource.ID, func(role *system.Resource) uint32 { return role.ID })
	if ok {
		return true, nil
	}
	return false, merr.ErrorPermissionDenied("team role resourceRepo is invalid.")
}
