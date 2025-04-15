package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

type Resource struct {
}

type BatchUpdateResourceStatusReq struct {
	IDs    []uint32          `json:"ids"`
	Status vobj.GlobalStatus `json:"status"`
}

type ListResourceReq struct {
	Statuses []vobj.GlobalStatus `json:"statuses"`
	Keyword  string              `json:"keyword"`
	*PaginationRequest
}

func (r *ListResourceReq) ToListResourceReply(resources []*system.Resource) *ListResourceReply {
	return &ListResourceReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(resources, func(resource *system.Resource) do.Resource { return resource }),
	}
}

type ListResourceReply = ListReply[do.Resource]
