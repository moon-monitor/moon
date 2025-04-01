package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// ToResourceItem 将DO转换为BO
func ToResourceItem(resource *system.Resource) *bo.ResourceItem {
	if resource == nil {
		return nil
	}
	return &bo.ResourceItem{
		ID:     resource.ID,
		Name:   resource.Name,
		Path:   resource.Path,
		Remark: resource.Remark,
		Status: resource.Status,
		Module: vobj.ResourceModule(resource.Module), // 需要类型转换
		Domain: resource.Domain,
		Allow:  resource.Allow,
	}
}

func ToResourceItemList(resources []*system.Resource) []*bo.ResourceItem {
	result := make([]*bo.ResourceItem, 0, len(resources))
	for _, resource := range resources {
		result = append(result, ToResourceItem(resource))
	}
	return result
}

func ToResourceItemProto(resource *bo.ResourceItem) *common.ResourceItem {
	if resource == nil {
		return nil
	}
	return &common.ResourceItem{
		Id:        resource.ID,
		Name:      resource.Name,
		Path:      resource.Path,
		Remark:    resource.Remark,
		Status:    common.ResourceStatus(resource.Status),
		Module:    common.ResourceModule(resource.Module),
		Domain:    common.ResourceDomain(resource.Domain),
		CreatedAt: resource.CreatedAt.Format(time.DateTime),
		UpdatedAt: resource.UpdatedAt.Format(time.DateTime),
	}
}

func ToResourceItemProtoList(resources []*bo.ResourceItem) []*common.ResourceItem {
	result := make([]*common.ResourceItem, 0, len(resources))
	for _, resource := range resources {
		result = append(result, ToResourceItemProto(resource))
	}
	return result
}

func ToResourceDO(resource *bo.SaveResourceReq) *system.Resource {
	return &system.Resource{
		BaseModel: do.BaseModel{ID: resource.ID},
		Name:      resource.Name,
		Path:      resource.Path,
		Module:    vobj.ResourceDomain(resource.Module), // 需要类型转换
		Domain:    resource.Domain,
		Status:    vobj.ResourceStatusEnabled,   // 默认启用
		Allow:     vobj.ResourceAllowSystemRBAC, // 默认系统RBAC
		Remark:    resource.Remark,
	}
}
