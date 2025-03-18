package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
)

type Resource interface {
	// GetResources get all resources
	GetResources(ctx context.Context) ([]*system.Resource, error)

	// GetResourceByID get resource by id
	GetResourceByID(ctx context.Context, id uint32) (*system.Resource, error)

	// GetResourceByOperation get resource by operation
	GetResourceByOperation(ctx context.Context, operation string) (*system.Resource, error)

	// SaveResource save resource
	//  exist id update, else insert
	SaveResource(ctx context.Context, resource *system.Resource) error

	// DeleteResource delete resource by id
	DeleteResource(ctx context.Context, id uint32) error
}
