package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// ToSaveEmailConfigRequest converts proto request to business object
func ToSaveEmailConfigRequest(req *palacev1.SaveEmailConfigRequest) *bo.SaveEmailConfigRequest {
	if req == nil {
		return nil
	}

	return &bo.SaveEmailConfigRequest{
		Config: &do.Email{
			User: req.GetUser(),
			Pass: req.GetPass(),
			Host: req.GetHost(),
			Port: req.GetPort(),
			Name: req.GetName(),
		},
		ID:     req.GetId(),
		Name:   req.GetName(),
		Remark: req.GetRemark(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}
}

// ToEmailConfigReply converts business object to proto reply
func ToEmailConfigReply(configs *bo.ListEmailConfigListReply) *palacev1.GetEmailConfigsReply {
	if validate.IsNil(configs) {
		return &palacev1.GetEmailConfigsReply{}
	}

	items := make([]*palacev1.EmailConfigItem, 0, len(configs.Items))
	for _, config := range configs.Items {
		items = append(items, &palacev1.EmailConfigItem{
			User:   config.GetUser(),
			Pass:   config.GetPass(),
			Host:   config.GetHost(),
			Port:   config.GetPort(),
			Status: common.GlobalStatus(config.GetStatus()),
			Name:   config.GetName(),
			Remark: config.GetRemark(),
			Id:     config.GetID(),
		})
	}

	return &palacev1.GetEmailConfigsReply{
		Items:      items,
		Pagination: ToPaginationReplyProto(configs.PaginationReply),
	}
}

// ToListEmailConfigRequest converts proto request to business object
func ToListEmailConfigRequest(req *palacev1.GetEmailConfigsRequest) *bo.ListEmailConfigRequest {
	if req == nil {
		return nil
	}

	return &bo.ListEmailConfigRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Keyword:           req.GetKeyword(),
		Status:            vobj.GlobalStatus(req.GetStatus()),
	}
}
