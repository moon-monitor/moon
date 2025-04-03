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
)

// NewResourceBiz 创建资源业务逻辑
func NewResourceBiz(
	resourceRepo repository.Resource,
	logger log.Logger,
) *ResourceBiz {
	return &ResourceBiz{
		resourceRepo: resourceRepo,
		helper:       log.NewHelper(log.With(logger, "module", "biz.resource")),
	}
}

type ResourceBiz struct {
	resourceRepo repository.Resource
	helper       *log.Helper
}

func (r *ResourceBiz) BatchUpdateResourceStatus(ctx context.Context, req *bo.BatchUpdateResourceStatusReq) error {
	return r.resourceRepo.BatchUpdateResourceStatus(ctx, req.IDs, req.Status)
}

func (r *ResourceBiz) GetResource(ctx context.Context, id uint32) (*system.Resource, error) {
	return r.resourceRepo.GetResourceByID(ctx, id)
}

func (r *ResourceBiz) ListResource(ctx context.Context, req *bo.ListResourceReq) (*bo.ListResourceReply, error) {
	return r.resourceRepo.ListResources(ctx, req)
}

func (r *ResourceBiz) SelfMenus(ctx context.Context) ([]*system.Menu, error) {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("user id not found")
	}
	return r.resourceRepo.GetMenusByUserID(ctx, userID)
}

func (r *ResourceBiz) Menus(ctx context.Context, t vobj.MenuType) ([]*system.Menu, error) {
	return r.resourceRepo.GetMenus(ctx, t)
}
