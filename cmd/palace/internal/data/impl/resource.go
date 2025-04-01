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

// BatchUpdateResourceStatus 批量更新资源状态
func (r *resourceImpl) BatchUpdateResourceStatus(ctx context.Context, ids []uint32, status vobj.ResourceStatus) error {
	if len(ids) == 0 {
		return nil
	}
	_, err := r.Resource.WithContext(ctx).
		Where(r.Resource.ID.In(ids...)).
		Update(r.Resource.Status, int8(status))
	return err
}

// ListResources 查询资源列表
func (r *resourceImpl) ListResources(ctx context.Context, req *bo.ListResourceReq) ([]*system.Resource, error) {
	resourceQuery := r.Resource.WithContext(ctx)

	// 添加状态过滤
	if len(req.Statuses) > 0 {
		statusValues := make([]int8, 0, len(req.Statuses))
		for _, s := range req.Statuses {
			statusValues = append(statusValues, int8(s))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Status.In(statusValues...))
	}

	// 添加模块过滤
	if len(req.Modules) > 0 {
		moduleValues := make([]int8, 0, len(req.Modules))
		for _, m := range req.Modules {
			moduleValues = append(moduleValues, int8(m))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Module.In(moduleValues...))
	}

	// 添加领域过滤
	if len(req.Domains) > 0 {
		domainValues := make([]int8, 0, len(req.Domains))
		for _, d := range req.Domains {
			domainValues = append(domainValues, int8(d))
		}
		resourceQuery = resourceQuery.Where(r.Resource.Domain.In(domainValues...))
	}

	// 添加关键字搜索
	if req.Keyword != "" {
		keywordPattern := "%" + req.Keyword + "%"
		opts := []gen.Condition{
			r.Resource.Name.Like(keywordPattern),
			r.Resource.Path.Like(keywordPattern),
			r.Resource.Remark.Like(keywordPattern),
		}
		resourceQuery = resourceQuery.Where(resourceQuery.Or(opts...))
	}

	// 排序和查询
	return resourceQuery.Order(r.Resource.ID.Desc()).Find()
}
