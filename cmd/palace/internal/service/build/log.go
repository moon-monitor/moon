package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

type OperateLogListRequest interface {
	GetPagination() *common.PaginationRequest
	GetOperateTypes() []common.OperateType
	GetKeyword() string
	GetUserId() uint32
}

func ToOperateLogListRequest(req OperateLogListRequest) *bo.OperateLogListRequest {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.OperateLogListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		OperateTypes:      slices.Map(req.GetOperateTypes(), func(operateType common.OperateType) vobj.OperateType { return vobj.OperateType(operateType) }),
		Keyword:           req.GetKeyword(),
		UserID:            req.GetUserId(),
	}
}

func ToOperateLogItem(log do.OperateLog) *common.OperateLogItem {
	if validate.IsNil(log) {
		return nil
	}
	return &common.OperateLogItem{
		Id:          log.GetID(),
		Operator:    UserToUserBaseItemProto(log.GetCreator()),
		Type:        common.OperateType(log.GetOperateType()),
		Module:      common.ResourceModule(log.GetOperateModule()),
		DataID:      log.GetOperateDataID(),
		DataName:    log.GetOperateDataName(),
		OperateTime: log.GetCreatedAt().Format(time.DateTime),
		Title:       log.GetTitle(),
		Before:      log.GetBefore(),
		After:       log.GetAfter(),
		Ip:          log.GetIP(),
	}
}

func ToOperateLogItems(logs []do.OperateLog) []*common.OperateLogItem {
	return slices.Map(logs, ToOperateLogItem)
}
