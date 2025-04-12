package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Resource interface {
	// GetResources get all resources
	GetResources(ctx context.Context) ([]do.Resource, error)

	// GetResourceByID get resource by id
	GetResourceByID(ctx context.Context, id uint32) (do.Resource, error)

	// GetResourceByOperation get resource by operation
	GetResourceByOperation(ctx context.Context, operation string) (do.Resource, error)

	// ListResources list resources
	ListResources(ctx context.Context, req *bo.ListResourceReq) (*bo.ListResourceReply, error)

	// BatchUpdateResourceStatus update multiple resources status
	BatchUpdateResourceStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error

	// GetMenusByUserID get all menus
	GetMenusByUserID(ctx context.Context, userID uint32) ([]do.Menu, error)

	// GetMenus get all menus
	GetMenus(ctx context.Context, t vobj.MenuType) ([]do.Menu, error)
}
