package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func ToPaginationReplyProto(pagination *bo.PaginationReply) *common.PaginationReply {
	return &common.PaginationReply{
		Total: pagination.Total,
		Page:  pagination.Page,
		Limit: pagination.Limit,
	}
}

func ToPaginationRequest(pagination *common.PaginationRequest) *bo.PaginationRequest {
	return &bo.PaginationRequest{
		Page:  pagination.GetPage(),
		Limit: pagination.GetLimit(),
	}
}
