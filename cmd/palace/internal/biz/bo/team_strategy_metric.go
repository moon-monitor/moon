package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

type SaveTeamStrategyItem struct {
	StrategyID   uint32
	Name         string
	Remark       string
	StrategyType vobj.StrategyType
}

type SaveLabelNoticesItem struct {
	LabelKey       string
	LabelValue     string
	ReceiverRoutes []uint32
}

type SaveTeamMetricStrategyRuleItem struct {
	LevelID        uint32
	LevelName      string
	SampleMode     vobj.SampleMode
	Condition      vobj.ConditionMetric
	Count          int64
	Values         []float64
	Duration       time.Duration
	ReceiverRoutes []uint32
	Status         vobj.GlobalStatus
	LabelNotices   []*SaveLabelNoticesItem
}

type SaveTeamMetricStrategyParams struct {
	Strategy       *SaveTeamStrategyItem
	Expr           string
	Labels         kv.StringMap
	Annotations    kv.StringMap
	ReceiverRoutes []uint32
	DatasourceList []uint32
	Rules          []*SaveTeamMetricStrategyRuleItem

	datasource     []do.DatasourceMetric
	receiverRoutes []do.NoticeGroup
	levels         []do.TeamDict
	strategy       do.Strategy
}

func (s *SaveTeamMetricStrategyParams) GetLevelIds() []uint32 {
	return slices.Map(s.Rules, func(rule *SaveTeamMetricStrategyRuleItem) uint32 { return rule.LevelID })
}

func (s *SaveTeamMetricStrategyParams) GetReceiverRouteIds() []uint32 {
	routes := s.ReceiverRoutes
	for _, rule := range s.Rules {
		routes = append(routes, rule.ReceiverRoutes...)
	}
	return slices.Unique(routes)
}

func (s *SaveTeamMetricStrategyParams) WithStrategy(strategy do.Strategy) *SaveTeamMetricStrategyParams {
	s.strategy = strategy
	return s
}

func (s *SaveTeamMetricStrategyParams) WithDatasourceList(datasourceList []do.DatasourceMetric) *SaveTeamMetricStrategyParams {
	s.datasource = datasourceList
	return s
}

func (s *SaveTeamMetricStrategyParams) WithReceiverRoutes(receiverRoutes []do.NoticeGroup) *SaveTeamMetricStrategyParams {
	s.receiverRoutes = receiverRoutes
	return s
}

func (s *SaveTeamMetricStrategyParams) WithLevels(levels []do.TeamDict) *SaveTeamMetricStrategyParams {
	s.levels = levels
	return s
}

func (s *SaveTeamMetricStrategyParams) Validate() error {
	if validate.IsNil(s.strategy) {
		return merr.ErrorParamsError("策略不存在")
	}
	if len(s.receiverRoutes) != len(s.GetReceiverRouteIds()) {
		return merr.ErrorParamsError("告警组不能为空")
	}
	if len(s.levels) != len(s.GetLevelIds()) {
		return merr.ErrorParamsError("告警级别不存在")
	}
	if len(s.datasource) != len(s.DatasourceList) {
		return merr.ErrorParamsError("数据源不存在")
	}
	return nil
}

type UpdateTeamStrategiesStatusParams struct {
	StrategyIds []uint32
	Status      vobj.GlobalStatus
}

type OperateTeamMetricStrategyParams struct {
	StrategyID       uint32
	MetricStrategyID uint32
	Preload          bool
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
