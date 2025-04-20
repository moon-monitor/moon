package build

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToSaveTeamMetricDatasourceRequest(req *palacev1.SaveTeamMetricDatasourceRequest) *bo.SaveTeamMetricDatasource {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.SaveTeamMetricDatasource{
		ID:             req.GetDatasourceId(),
		Name:           req.GetName(),
		Status:         vobj.GlobalStatusEnable,
		Remark:         req.GetRemark(),
		Driver:         vobj.DatasourceDriverMetric(req.GetMetricDatasourceDriver()),
		Endpoint:       req.GetEndpoint(),
		ScrapeInterval: req.GetScrapeInterval().AsDuration(),
		Headers:        req.GetHeaders(),
		QueryMethod:    vobj.HTTPMethod(req.GetQueryMethod()),
		CA:             req.GetCa(),
		TLS:            ToTLS(req.GetTls()),
		BasicAuth:      ToBasicAuth(req.GetBasicAuth()),
		Extra:          req.GetExtra(),
	}
}

func ToListTeamMetricDatasourceRequest(req *palacev1.ListTeamMetricDatasourceRequest) *bo.ListTeamMetricDatasource {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.ListTeamMetricDatasource{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Status:            vobj.GlobalStatus(req.GetStatus()),
		Keyword:           req.GetKeyword(),
	}
}

func ToTeamMetricDatasourceItem(item do.DatasourceMetric) *common.TeamMetricDatasourceItem {
	if validate.IsNil(item) {
		return nil
	}
	return &common.TeamMetricDatasourceItem{
		TeamID:         item.GetTeamID(),
		DatasourceID:   item.GetID(),
		CreatedAt:      item.GetCreatedAt().Format(time.DateTime),
		UpdatedAt:      item.GetUpdatedAt().Format(time.DateTime),
		Name:           item.GetName(),
		Remark:         item.GetRemark(),
		Driver:         common.DatasourceDriverMetric(item.GetDriver()),
		Endpoint:       item.GetEndpoint(),
		ScrapeInterval: durationpb.New(item.GetScrapeInterval()),
		Headers:        item.GetHeaders(),
		QueryMethod:    common.HTTPMethod(item.GetQueryMethod()),
		Ca:             item.GetCA(),
		Tls:            ToProtoTLS(item.GetTLS()),
		BasicAuth:      ToProtoBasicAuth(item.GetBasicAuth()),
		Extra:          item.GetExtra(),
		Status:         common.GlobalStatus(item.GetStatus()),
		Creator:        ToUserBaseItem(item.GetCreator()),
	}
}

func ToTeamMetricDatasourceItems(items []do.DatasourceMetric) []*common.TeamMetricDatasourceItem {
	return slices.Map(items, ToTeamMetricDatasourceItem)
}
