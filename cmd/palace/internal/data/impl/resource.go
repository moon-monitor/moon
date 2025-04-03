package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
)

func NewResourceRepo(d *data.Data, logger log.Logger) repository.Resource {
	return &resourceImpl{
		Data:   d,
		Query:  systemgen.Use(d.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.resource")),
	}
}

type resourceImpl struct {
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func (r *resourceImpl) GetResources(ctx context.Context) ([]*system.Resource, error) {
	return r.Resource.WithContext(ctx).Find()
}

func (r *resourceImpl) GetResourceByID(ctx context.Context, id uint32) (*system.Resource, error) {
	resourceDo, err := r.Resource.WithContext(ctx).Where(r.Resource.ID.Eq(id)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
}

func (r *resourceImpl) GetResourceByOperation(ctx context.Context, operation string) (*system.Resource, error) {
	resourceDo, err := r.Resource.WithContext(ctx).Where(r.Resource.Path.Eq(operation)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
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

func (r *resourceImpl) BatchUpdateResourceStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error {
	if len(ids) == 0 {
		return nil
	}
	_, err := r.Resource.WithContext(ctx).
		Where(r.Resource.ID.In(ids...)).
		Update(r.Resource.Status, int8(status))
	return err
}

func (r *resourceImpl) ListResources(ctx context.Context, req *bo.ListResourceReq) ([]*system.Resource, error) {
	resourceQuery := r.Resource.WithContext(ctx)

	if len(req.Statuses) > 0 {
		statusValues := make([]int8, 0, len(req.Statuses))
		for _, s := range req.Statuses {
			statusValues = append(statusValues, int8(s))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Status.In(statusValues...))
	}

	if len(req.Modules) > 0 {
		moduleValues := make([]int8, 0, len(req.Modules))
		for _, m := range req.Modules {
			moduleValues = append(moduleValues, int8(m))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Module.In(moduleValues...))
	}

	if len(req.Domains) > 0 {
		domainValues := make([]int8, 0, len(req.Domains))
		for _, d := range req.Domains {
			domainValues = append(domainValues, int8(d))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Domain.In(domainValues...))
	}

	if req.Keyword != "" {
		keywordPattern := "%" + req.Keyword + "%"
		opts := []gen.Condition{
			r.Resource.Name.Like(keywordPattern),
			r.Resource.Path.Like(keywordPattern),
			r.Resource.Remark.Like(keywordPattern),
		}
		resourceQuery = resourceQuery.Where(resourceQuery.Or(opts...))
	}

	return resourceQuery.Order(r.Resource.ID.Desc()).Find()
}
