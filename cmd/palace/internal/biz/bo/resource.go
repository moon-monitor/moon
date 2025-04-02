package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

// SaveResourceReq 保存资源请求
type SaveResourceReq struct {
	ID     uint32              `json:"id"`
	Name   string              `json:"name"`
	Path   string              `json:"path"`
	Remark string              `json:"remark"`
	Module vobj.ResourceModule `json:"module"`
	Domain vobj.ResourceDomain `json:"domain"`
	Allow  vobj.ResourceAllow  `json:"allow"`
}

// BatchUpdateResourceStatusReq 批量更新资源状态请求
type BatchUpdateResourceStatusReq struct {
	IDs    []uint32          `json:"ids"`
	Status vobj.GlobalStatus `json:"status"`
}

// ListResourceReq 资源列表查询请求
type ListResourceReq struct {
	Statuses []vobj.GlobalStatus   `json:"statuses"`
	Modules  []vobj.ResourceModule `json:"modules"`
	Domains  []vobj.ResourceDomain `json:"domains"`
	Keyword  string                `json:"keyword"`
}

// ResourceItem 资源项
type ResourceItem struct {
	ID        uint32              `json:"id"`
	Name      string              `json:"name"`
	Path      string              `json:"path"`
	Remark    string              `json:"remark"`
	Status    vobj.GlobalStatus   `json:"status"`
	Module    vobj.ResourceModule `json:"module"` // 注意：系统中实际储存为ResourceDomain类型
	Domain    vobj.ResourceDomain `json:"domain"`
	Allow     vobj.ResourceAllow  `json:"allow"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
