package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

type SaveTeamMetricStrategyParams struct {
}

type UpdateTeamStrategiesStatusParams struct {
	StrategyIds []uint32
	Status      vobj.GlobalStatus
}

type OperateTeamMetricStrategyParams struct {
	StrategyID       uint32
	MetricStrategyID uint32
}

type ListTeamStrategyParams struct {
	*PaginationRequest
	Keyword string
	Status  []vobj.GlobalStatus
}

func (l *ListTeamStrategyParams) ToListTeamStrategyReply(items []*team.Strategy) *ListTeamStrategyReply {
	return &ListTeamStrategyReply{
		Items:           slices.Map(items, func(item *team.Strategy) do.Strategy { return item }),
		PaginationReply: l.PaginationRequest.ToReply(),
	}
}

type ListTeamStrategyReply = ListReply[do.Strategy]

type ToSubscribeTeamStrategyParams struct {
	StrategyID       uint32
	MetricStrategyID uint32
	SubscribeType    vobj.NoticeType
}

type ToSubscribeTeamStrategiesParams struct {
	StrategyID    uint32
	Subscribers   []uint32
	SubscribeType vobj.NoticeType
	*PaginationRequest
}

func (t *ToSubscribeTeamStrategiesParams) ToSubscribeTeamStrategiesReply(items []*team.StrategySubscriber) *ToSubscribeTeamStrategiesReply {
	return &ToSubscribeTeamStrategiesReply{
		Items:           slices.Map(items, func(item *team.StrategySubscriber) do.TeamStrategySubscriber { return item }),
		PaginationReply: t.PaginationRequest.ToReply(),
	}
}

type ToSubscribeTeamStrategiesReply = ListReply[do.TeamStrategySubscriber]
