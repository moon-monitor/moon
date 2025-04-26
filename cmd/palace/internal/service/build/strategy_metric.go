package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToSaveTeamStrategyItem(request *palacev1.SaveTeamStrategyItemRequest) *bo.SaveTeamStrategyItem {
	return &bo.SaveTeamStrategyItem{
		StrategyID:     request.GetStrategyId(),
		Name:           request.GetName(),
		Remark:         request.GetRemark(),
		StrategyType:   vobj.StrategyType(request.GetStrategyType()),
		ReceiverRoutes: request.GetReceiverRoutes(),
	}
}

func ToSaveLabelNoticeItem(request *palacev1.LabelNotices) *bo.SaveLabelNoticeItem {
	return &bo.SaveLabelNoticeItem{
		LabelKey:       request.GetKey(),
		LabelValue:     request.GetValue(),
		ReceiverRoutes: request.GetReceiverRoutes(),
	}
}

func ToSaveLabelNoticeItems(request []*palacev1.LabelNotices) []*bo.SaveLabelNoticeItem {
	return slices.Map(request, ToSaveLabelNoticeItem)
}

func ToSaveTeamMetricStrategyRuleItem(request *palacev1.SaveTeamMetricStrategyRequest_MetricStrategyLevelItem) *bo.SaveTeamMetricStrategyRuleItem {
	return &bo.SaveTeamMetricStrategyRuleItem{
		LevelID:        request.GetLevelId(),
		LevelName:      request.GetLevelName(),
		SampleMode:     vobj.SampleMode(request.GetSampleMode()),
		Condition:      vobj.ConditionMetric(request.GetCondition()),
		Count:          request.GetCount(),
		Values:         request.GetValues(),
		Duration:       request.GetDuration().AsDuration(),
		ReceiverRoutes: request.GetReceiverRoutes(),
		Status:         vobj.GlobalStatus(request.GetStatus()),
		LabelNotices:   ToSaveLabelNoticeItems(request.GetLabelNotices()),
	}
}

func ToSaveTeamMetricStrategyRuleItems(request []*palacev1.SaveTeamMetricStrategyRequest_MetricStrategyLevelItem) []*bo.SaveTeamMetricStrategyRuleItem {
	return slices.Map(request, ToSaveTeamMetricStrategyRuleItem)
}

func ToSaveTeamMetricStrategyParams(request *palacev1.SaveTeamMetricStrategyRequest) *bo.SaveTeamMetricStrategyParams {
	return &bo.SaveTeamMetricStrategyParams{
		Strategy:       ToSaveTeamStrategyItem(request.GetStrategy()),
		Expr:           request.GetExpr(),
		Labels:         kv.NewStringMap(request.GetLabels()),
		Annotations:    kv.NewStringMap(request.GetAnnotations()),
		DatasourceList: request.GetDatasource(),
		Rules:          ToSaveTeamMetricStrategyRuleItems(request.GetLevels()),
	}
}

func ToUpdateTeamStrategiesStatusParams(request *palacev1.UpdateTeamStrategiesStatusRequest) *bo.UpdateTeamStrategiesStatusParams {
	return &bo.UpdateTeamStrategiesStatusParams{
		StrategyIds: request.GetStrategyIds(),
		Status:      vobj.GlobalStatus(request.GetStatus()),
	}
}

func ToOperateTeamMetricStrategyParams(request *palacev1.OperateTeamMetricStrategyRequest) *bo.OperateTeamMetricStrategyParams {
	return &bo.OperateTeamMetricStrategyParams{
		StrategyID:       request.GetStrategyId(),
		MetricStrategyID: request.GetMetricStrategyId(),
		Preload:          true,
	}
}

func ToTeamMetricStrategyItem(strategy do.StrategyMetric) *common.TeamStrategyMetricItem {
	return &common.TeamStrategyMetricItem{}
}

func ToTeamMetricStrategyItems(strategies []do.StrategyMetric) []*common.TeamStrategyMetricItem {
	return slices.Map(strategies, ToTeamMetricStrategyItem)
}

func ToListTeamStrategyParams(request *palacev1.ListTeamStrategyRequest) *bo.ListTeamStrategyParams {
	return &bo.ListTeamStrategyParams{
		PaginationRequest: ToPaginationRequest(request.GetPagination()),
		Keyword:           request.GetKeyword(),
		Status:            slices.Map(request.GetStatus(), func(status common.GlobalStatus) vobj.GlobalStatus { return vobj.GlobalStatus(status) }),
	}
}

func ToTeamStrategyItem(strategy do.Strategy) *common.TeamStrategyItem {
	return &common.TeamStrategyItem{}
}

func ToTeamStrategyItems(strategies []do.Strategy) []*common.TeamStrategyItem {
	return slices.Map(strategies, ToTeamStrategyItem)
}

func ToSubscribeTeamStrategyParams(request *palacev1.SubscribeTeamStrategyRequest) *bo.ToSubscribeTeamStrategyParams {
	return &bo.ToSubscribeTeamStrategyParams{
		StrategyID:       request.GetStrategyId(),
		MetricStrategyID: request.GetMetricStrategyId(),
		SubscribeType:    vobj.NoticeType(request.GetSubscribeType()),
	}
}

func ToSubscribeTeamStrategiesParams(request *palacev1.SubscribeTeamStrategiesRequest) *bo.ToSubscribeTeamStrategiesParams {
	return &bo.ToSubscribeTeamStrategiesParams{
		StrategyID:        request.GetStrategyId(),
		Subscribers:       request.GetSubscribers(),
		SubscribeType:     vobj.NoticeType(request.GetSubscribeType()),
		PaginationRequest: ToPaginationRequest(request.GetPagination()),
	}
}

func ToSubscribeTeamStrategiesItem(subscriber do.TeamStrategySubscriber) *common.SubscriberItem {
	return &common.SubscriberItem{}
}

func ToSubscribeTeamStrategiesItems(subscribers []do.TeamStrategySubscriber) []*common.SubscriberItem {
	return slices.Map(subscribers, ToSubscribeTeamStrategiesItem)
}
