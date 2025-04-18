package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// ToSaveSMSConfigRequest converts API request to business object
func ToSaveSMSConfigRequest(req *palace.SaveSMSConfigRequest) *bo.SaveSMSConfigRequest {
	return &bo.SaveSMSConfigRequest{
		Config: &do.SMS{
			AccessKeyID:     req.GetAccessKeyID(),
			AccessKeySecret: req.GetAccessKeySecret(),
			SignName:        req.GetSignName(),
			Endpoint:        req.GetEndpoint(),
		},
		ID:       req.GetId(),
		Name:     req.GetName(),
		Remark:   req.GetRemark(),
		Status:   vobj.GlobalStatus(req.GetStatus()),
		Provider: vobj.SMSProviderType(req.GetProvider()),
	}
}

// ToListSMSConfigRequest converts API request to business object
func ToListSMSConfigRequest(req *palace.GetSMSConfigsRequest) *bo.ListSMSConfigRequest {
	return &bo.ListSMSConfigRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Keyword:           req.GetKeyword(),
		Status:            vobj.GlobalStatus(req.GetStatus()),
		Provider:          vobj.SMSProviderType(req.GetProvider()),
	}
}

// ToSMSConfigReply converts business object to API response
func ToSMSConfigReply(reply *bo.ListSMSConfigListReply) *palace.GetSMSConfigsReply {
	return &palace.GetSMSConfigsReply{
		Pagination: ToPaginationReplyProto(reply.PaginationReply),
		Items: slices.MapFilter(reply.Items, func(item do.TeamSMSConfig) (*palace.SMSConfigItem, bool) {
			config := item.GetSMSConfig()
			if validate.IsNil(config) {
				return nil, false
			}
			return &palace.SMSConfigItem{
				ProviderType:    common.SMSProviderType(item.GetProviderType()),
				AccessKeyID:     config.AccessKeyID,
				AccessKeySecret: config.AccessKeySecret,
				SignName:        config.SignName,
				Endpoint:        config.Endpoint,
				Name:            item.GetName(),
				Remark:          item.GetRemark(),
				Id:              item.GetID(),
				Status:          common.GlobalStatus(item.GetStatus()),
			}, true
		}),
	}
}
