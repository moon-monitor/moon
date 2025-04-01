package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
)

// NewResourceBiz 创建资源业务逻辑
func NewResourceBiz(
	repo repository.Resource,
	logger log.Logger,
) *ResourceBiz {
	return &ResourceBiz{
		repo:   repo,
		helper: log.NewHelper(log.With(logger, "module", "biz.resource")),
	}
}

// ResourceBiz 资源业务逻辑
type ResourceBiz struct {
	repo   repository.Resource
	helper *log.Helper
}

// SaveResource 保存资源
func (r *ResourceBiz) SaveResource(ctx context.Context, req *bo.SaveResourceReq) error {
	// 转换为数据对象
	resource := build.ToResourceDO(req)
	resource.WithContext(ctx)

	// 保存资源
	return r.repo.SaveResource(ctx, resource)
}

// BatchUpdateResourceStatus 批量更新资源状态
func (r *ResourceBiz) BatchUpdateResourceStatus(ctx context.Context, req *bo.BatchUpdateResourceStatusReq) error {
	return r.repo.BatchUpdateResourceStatus(ctx, req.IDs, req.Status)
}

// DeleteResource 删除资源
func (r *ResourceBiz) DeleteResource(ctx context.Context, id uint32) error {
	return r.repo.DeleteResource(ctx, id)
}

// GetResource 获取资源详情
func (r *ResourceBiz) GetResource(ctx context.Context, id uint32) (*bo.ResourceItem, error) {
	// 查询资源
	resource, err := r.repo.GetResourceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 转换为业务对象
	return build.ToResourceItem(resource), nil
}

// ListResource 查询资源列表
func (r *ResourceBiz) ListResource(ctx context.Context, req *bo.ListResourceReq) ([]*bo.ResourceItem, error) {
	// 查询资源列表
	resources, err := r.repo.ListResources(ctx, req)
	if err != nil {
		return nil, err
	}

	return build.ToResourceItemList(resources), nil
}
