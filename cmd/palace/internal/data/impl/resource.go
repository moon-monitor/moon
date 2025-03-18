package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
)

func NewResourceRepo(d *data.Data, logger log.Logger) repository.Resource {
	return &resourceImpl{
		Data:   d,
		Query:  systemQuery.Use(d.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.resource")),
	}
}

type resourceImpl struct {
	*data.Data

	*systemQuery.Query

	helper *log.Helper
}

func (r *resourceImpl) GetResources(ctx context.Context) ([]*system.Resource, error) {
	return r.Resource.WithContext(ctx).Find()
}

func (r *resourceImpl) GetResourceByID(ctx context.Context, id uint32) (*system.Resource, error) {
	return r.Resource.WithContext(ctx).Where(r.Resource.ID.Eq(id)).First()
}

func (r *resourceImpl) GetResourceByOperation(ctx context.Context, operation string) (*system.Resource, error) {
	return r.Resource.WithContext(ctx).Where(r.Resource.Path.Eq(operation)).First()
}

func (r *resourceImpl) SaveResource(ctx context.Context, resource *system.Resource) error {
	if resource.ID == 0 {
		return r.createResource(ctx, resource)
	}
	return r.updateByID(ctx, resource)
}

func (r *resourceImpl) createResource(ctx context.Context, resource *system.Resource) error {
	return r.Resource.WithContext(ctx).Create(resource)
}

func (r *resourceImpl) updateByID(ctx context.Context, resource *system.Resource) error {
	_, err := r.Resource.WithContext(ctx).Where(r.Resource.ID.Eq(resource.ID)).Updates(resource)
	return err
}

func (r *resourceImpl) DeleteResource(ctx context.Context, id uint32) error {
	_, err := r.Resource.WithContext(ctx).Where(r.Resource.ID.Eq(id)).Delete()
	return err
}
