package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
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

type ListResourceReply struct {
	*PaginationReply
	Resources []do.Resource `json:"resources"`
}
