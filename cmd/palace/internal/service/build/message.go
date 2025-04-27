package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToListSendMessageLogParams(req *palace.SendMessageLogsRequest) *bo.ListSendMessageLogParams {
	if validate.IsNil(req) {
		panic("SendMessageLogsRequest is nil")
	}
	return &bo.ListSendMessageLogParams{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		TeamID:            0,
		RequestID:         req.GetRequestId(),
		Status:            vobj.SendMessageStatus(req.GetStatus()),
		Keyword:           req.GetKeyword(),
		TimeRange:         []time.Time{},
		MessageType:       vobj.MessageType(req.GetMessageType()),
	}
}
