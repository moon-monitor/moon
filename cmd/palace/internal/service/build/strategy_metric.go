package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToSaveTeamMetricStrategyParams(request *palacev1.SaveTeamMetricStrategyRequest) *bo.SaveTeamMetricStrategyParams {
	return &bo.SaveTeamMetricStrategyParams{}
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
