package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

var _ SaveTeamStrategy = (*SaveTeamStrategyItem)(nil)

type SaveTeamStrategy interface {
	GetID() uint32
	GetReceiverRoutes() []do.NoticeGroup
	GetName() string
	GetRemark() string
	GetStrategyType() vobj.StrategyType
	GetStrategyGroupID() uint32
	GetStrategyGroup() do.StrategyGroup
	GetStrategy() do.Strategy
}

type SaveTeamStrategyItem struct {
	StrategyGroupID uint32
	StrategyID      uint32
	Name            string
	Remark          string
	StrategyType    vobj.StrategyType
	ReceiverRoutes  []uint32

	receiverRoutes map[uint32]do.NoticeGroup
	strategy       do.Strategy
	strategyGroup  do.StrategyGroup
}

func (r *SaveTeamStrategyItem) GetID() uint32 {
	if validate.IsNil(r) {
		return 0
	}
	if validate.IsNil(r.strategy) {
		return r.StrategyID
	}
	return r.strategy.GetID()
}

func (r *SaveTeamStrategyItem) GetReceiverRoutes() []do.NoticeGroup {
	return kv.New(r.receiverRoutes).Values()
}

func (r *SaveTeamStrategyItem) GetName() string {
	if validate.IsNil(r) {
		return ""
	}
	if validate.IsNotNil(r.strategy) && validate.TextIsNull(r.strategy.GetName()) {
		return r.strategy.GetName()
	}
	return r.Name
}

func (r *SaveTeamStrategyItem) GetRemark() string {
	if validate.IsNil(r) {
		return ""
	}
	if validate.IsNotNil(r.strategy) && validate.TextIsNull(r.strategy.GetRemark()) {
		return r.strategy.GetRemark()
	}
	return r.Remark
}

func (r *SaveTeamStrategyItem) GetStrategyType() vobj.StrategyType {
	if validate.IsNil(r) {
		return vobj.StrategyTypeUnknown
	}
	if validate.IsNotNil(r.strategy) && !r.strategy.GetStrategyType().Exist() {
		return r.strategy.GetStrategyType()
	}
	return r.StrategyType
}

func (r *SaveTeamStrategyItem) GetStrategyGroupID() uint32 {
	if validate.IsNil(r) {
		return 0
	}
	if validate.IsNotNil(r.strategyGroup) {
		return r.strategyGroup.GetID()
	}
	return r.StrategyGroupID
}

func (r *SaveTeamStrategyItem) GetStrategyGroup() do.StrategyGroup {
	if validate.IsNil(r) {
		return nil
	}
	return r.strategyGroup
}

func (r *SaveTeamStrategyItem) withStrategyGroup(strategyGroup do.StrategyGroup) *SaveTeamStrategyItem {
	r.strategyGroup = strategyGroup
	return r
}

func (r *SaveTeamStrategyItem) withReceiverRoutes(routes []do.NoticeGroup) *SaveTeamStrategyItem {
	r.receiverRoutes = slices.ToMap(routes, func(v do.NoticeGroup) uint32 {
		return v.GetID()
	})
	return r
}

func (r *SaveTeamStrategyItem) withStrategy(strategy do.Strategy) *SaveTeamStrategyItem {
	r.strategy = strategy
	return r
}

func (r *SaveTeamStrategyItem) GetStrategy() do.Strategy {
	return r.strategy
}

var _ SaveLabelNotice = (*SaveLabelNoticeItem)(nil)

type SaveLabelNotice interface {
	GetLabelKey() string
	GetLabelValue() string
	GetReceiverRoutes() []do.NoticeGroup
	GetStrategyMetricRuleID() uint32
	GetStrategyMetricRule() do.StrategyMetricRule
	GetID() uint32
	GetLabelNotice() do.StrategyMetricRuleLabelNotice
}

type SaveLabelNoticeItem struct {
	ID                   uint32
	LabelKey             string
	LabelValue           string
	ReceiverRoutes       []uint32
	StrategyMetricRuleID uint32

	labelNotice    do.StrategyMetricRuleLabelNotice
	receiverRoutes map[uint32]do.NoticeGroup
	metricRule     do.StrategyMetricRule
}

// GetID implements SaveLabelNotice.
func (r *SaveLabelNoticeItem) GetID() uint32 {
	if validate.IsNil(r) {
		return 0
	}
	if validate.IsNotNil(r.labelNotice) {
		return r.labelNotice.GetID()
	}
	return r.ID
}

// GetLabelNotice implements SaveLabelNotice.
func (r *SaveLabelNoticeItem) GetLabelNotice() do.StrategyMetricRuleLabelNotice {
	if validate.IsNil(r) {
		return nil
	}
	return r.labelNotice
}

// GetStrategyMetricRule implements SaveLabelNotice.
func (r *SaveLabelNoticeItem) GetStrategyMetricRule() do.StrategyMetricRule {
	if validate.IsNil(r) {
		return nil
	}
	return r.metricRule
}

// GetStrategyMetricRuleID implements SaveLabelNotice.
func (r *SaveLabelNoticeItem) GetStrategyMetricRuleID() uint32 {
	if validate.IsNil(r) {
		return 0
	}
	if validate.IsNotNil(r.metricRule) {
		return r.metricRule.GetID()
	}
	return r.StrategyMetricRuleID
}

func (r *SaveLabelNoticeItem) GetLabelKey() string {
	return r.LabelKey
}

func (r *SaveLabelNoticeItem) GetLabelValue() string {
	return r.LabelValue
}

func (r *SaveLabelNoticeItem) GetReceiverRoutes() []do.NoticeGroup {
	return kv.New(r.receiverRoutes).Values()
}

func (r *SaveLabelNoticeItem) withReceiverRoutes(routes []do.NoticeGroup) *SaveLabelNoticeItem {
	r.receiverRoutes = slices.ToMap(routes, func(v do.NoticeGroup) uint32 {
		return v.GetID()
	})
	return r
}

func (r *SaveLabelNoticeItem) withLabelNotice(labelNotice do.StrategyMetricRuleLabelNotice) *SaveLabelNoticeItem {
	r.labelNotice = labelNotice
	return r
}

var _ SaveTeamMetricStrategyRule = (*SaveTeamMetricStrategyRuleItem)(nil)

type SaveTeamMetricStrategyRule interface {
	GetID() uint32
	GetLevel() do.TeamDict
	GetAlarmPages() []do.TeamDict
	GetReceiverRoutes() []do.NoticeGroup
	GetStatus() vobj.GlobalStatus
	GetLabelNotices() []SaveLabelNotice
	GetSampleMode() vobj.SampleMode
	GetCondition() vobj.ConditionMetric
	GetCount() int64
	GetValues() []float64
	GetDuration() time.Duration
	GetMetricStrategyRule() do.StrategyMetricRule
	GetMetricStrategy() do.StrategyMetric
}

type SaveTeamMetricStrategyRuleItem struct {
	ID             uint32
	LevelID        uint32
	LevelName      string
	SampleMode     vobj.SampleMode
	Condition      vobj.ConditionMetric
	Count          int64
	Values         []float64
	Duration       time.Duration
	ReceiverRoutes []uint32
	Status         vobj.GlobalStatus
	LabelNotices   []*SaveLabelNoticeItem
	AlarmPages     []uint32

	receiverRoutes     map[uint32]do.NoticeGroup
	dicts              map[uint32]do.TeamDict
	metricStrategyRule do.StrategyMetricRule
	metricStrategy     do.StrategyMetric
}

func (r *SaveTeamMetricStrategyRuleItem) GetMetricStrategyRule() do.StrategyMetricRule {
	return r.metricStrategyRule
}

func (r *SaveTeamMetricStrategyRuleItem) GetMetricStrategy() do.StrategyMetric {
	return r.metricStrategy
}

func (r *SaveTeamMetricStrategyRuleItem) GetID() uint32 {
	if validate.IsNil(r) {
		return 0
	}
	if validate.IsNotNil(r.metricStrategy) {
		return r.metricStrategy.GetID()
	}
	return r.ID
}

func (r *SaveTeamMetricStrategyRuleItem) GetLevel() do.TeamDict {
	return r.dicts[r.LevelID]
}

func (r *SaveTeamMetricStrategyRuleItem) GetAlarmPages() []do.TeamDict {
	return slices.Map(r.AlarmPages, func(v uint32) do.TeamDict {
		return r.dicts[v]
	})
}

func (r *SaveTeamMetricStrategyRuleItem) GetReceiverRoutes() []do.NoticeGroup {
	return kv.New(r.receiverRoutes).Values()
}

func (r *SaveTeamMetricStrategyRuleItem) GetStatus() vobj.GlobalStatus {
	if validate.IsNil(r) {
		return vobj.GlobalStatusUnknown
	}
	if validate.IsNotNil(r.metricStrategyRule) && !r.metricStrategyRule.GetStatus().Exist() {
		return r.metricStrategyRule.GetStatus()
	}
	return r.Status
}

func (r *SaveTeamMetricStrategyRuleItem) GetLabelNotices() []SaveLabelNotice {
	return slices.Map(r.LabelNotices, func(v *SaveLabelNoticeItem) SaveLabelNotice {
		return v
	})
}

func (r *SaveTeamMetricStrategyRuleItem) GetCondition() vobj.ConditionMetric {
	if validate.IsNil(r) {
		return vobj.ConditionMetricUnknown
	}

	return r.Condition
}

func (r *SaveTeamMetricStrategyRuleItem) GetSampleMode() vobj.SampleMode {
	if validate.IsNil(r) {
		return vobj.SampleModeUnknown
	}

	return r.SampleMode
}

func (r *SaveTeamMetricStrategyRuleItem) GetCount() int64 {
	if validate.IsNil(r) {
		return 0
	}
	return r.Count
}

func (r *SaveTeamMetricStrategyRuleItem) GetValues() []float64 {
	if validate.IsNil(r) {
		return nil
	}
	return r.Values
}

func (r *SaveTeamMetricStrategyRuleItem) GetDuration() time.Duration {
	if validate.IsNil(r) {
		return 0
	}

	return r.Duration
}

func (r *SaveTeamMetricStrategyRuleItem) withReceiverRoutes(routes []do.NoticeGroup) *SaveTeamMetricStrategyRuleItem {
	r.receiverRoutes = slices.ToMap(routes, func(v do.NoticeGroup) uint32 {
		return v.GetID()
	})
	r.LabelNotices = slices.Map(r.LabelNotices, func(v *SaveLabelNoticeItem) *SaveLabelNoticeItem {
		return v.withReceiverRoutes(r.GetReceiverRoutes())
	})
	return r
}

func (r *SaveTeamMetricStrategyRuleItem) withDicts(dicts []do.TeamDict) *SaveTeamMetricStrategyRuleItem {
	r.dicts = slices.ToMap(dicts, func(v do.TeamDict) uint32 {
		return v.GetID()
	})
	return r
}

func (r *SaveTeamMetricStrategyRuleItem) withMetricStrategyRule(metricStrategyRule do.StrategyMetricRule) *SaveTeamMetricStrategyRuleItem {
	r.metricStrategyRule = metricStrategyRule
	return r
}

func (r *SaveTeamMetricStrategyRuleItem) withMetricStrategy(metricStrategy do.StrategyMetric) *SaveTeamMetricStrategyRuleItem {
	r.metricStrategy = metricStrategy
	return r
}

var _ SaveTeamMetricStrategy = (*SaveTeamMetricStrategyParams)(nil)

type SaveTeamMetricStrategy interface {
	GetStrategy() SaveTeamStrategy
	GetID() uint32
	GetExpr() string
	GetLabels() kv.StringMap
	GetAnnotations() kv.StringMap
	GetDatasourceList() []do.DatasourceMetric
	GetRules() []SaveTeamMetricStrategyRule
	GetStrategyMetric() do.StrategyMetric
}

type SaveTeamMetricStrategyParams struct {
	Strategy       *SaveTeamStrategyItem
	ID             uint32
	Expr           string
	Labels         kv.StringMap
	Annotations    kv.StringMap
	DatasourceList []uint32
	Rules          []*SaveTeamMetricStrategyRuleItem

	datasourceList []do.DatasourceMetric
	strategyMetric do.StrategyMetric
}

func (p *SaveTeamMetricStrategyParams) GetDictIds() []uint32 {
	if validate.IsNil(p) {
		return nil
	}
	list := make([]uint32, 0, len(p.Rules)*2)
	for _, rule := range p.Rules {
		list = append(list, rule.LevelID)
		for _, page := range rule.AlarmPages {
			list = append(list, page)
		}
	}
	return slices.MapFilter(slices.Unique(list), func(v uint32) (uint32, bool) {
		if v <= 0 {
			return 0, false
		}
		return v, true
	})
}

func (p *SaveTeamMetricStrategyParams) GetReceiverIds() []uint32 {
	if validate.IsNil(p) {
		return nil
	}
	list := make([]uint32, 0, len(p.Rules)*2)
	for _, rule := range p.Rules {
		list = append(list, rule.ReceiverRoutes...)
		for _, labelNotice := range rule.LabelNotices {
			list = append(list, labelNotice.ReceiverRoutes...)
		}
	}
	list = append(list, p.Strategy.ReceiverRoutes...)
	return slices.MapFilter(slices.Unique(list), func(v uint32) (uint32, bool) {
		if v <= 0 {
			return 0, false
		}
		return v, true
	})
}

func (p *SaveTeamMetricStrategyParams) GetLabelNoticeIds() []uint32 {
	if validate.IsNil(p) {
		return nil
	}
	list := make([]uint32, 0, len(p.Rules)*2)
	for _, rule := range p.Rules {
		for _, labelNotice := range rule.LabelNotices {
			list = append(list, labelNotice.ID)
		}
	}
	return slices.MapFilter(slices.Unique(list), func(v uint32) (uint32, bool) {
		if v <= 0 {
			return 0, false
		}
		return v, true
	})
}

func (p *SaveTeamMetricStrategyParams) GetStrategy() SaveTeamStrategy {
	if validate.IsNil(p) {
		return nil
	}
	return p.Strategy
}

func (p *SaveTeamMetricStrategyParams) GetID() uint32 {
	if validate.IsNil(p) {
		return 0
	}
	if validate.IsNotNil(p.strategyMetric) {
		return p.strategyMetric.GetID()
	}
	return p.ID
}

func (p *SaveTeamMetricStrategyParams) GetExpr() string {
	return p.Expr
}

func (p *SaveTeamMetricStrategyParams) GetLabels() kv.StringMap {
	return p.Labels
}

func (p *SaveTeamMetricStrategyParams) GetAnnotations() kv.StringMap {
	return p.Annotations
}

func (p *SaveTeamMetricStrategyParams) GetDatasourceList() []do.DatasourceMetric {
	return p.datasourceList
}

func (p *SaveTeamMetricStrategyParams) GetRules() []SaveTeamMetricStrategyRule {
	return slices.Map(p.Rules, func(v *SaveTeamMetricStrategyRuleItem) SaveTeamMetricStrategyRule {
		return v
	})
}

func (p *SaveTeamMetricStrategyParams) GetStrategyMetric() do.StrategyMetric {
	return p.strategyMetric
}

func (p *SaveTeamMetricStrategyParams) WithDatasourceList(datasourceList []do.DatasourceMetric) *SaveTeamMetricStrategyParams {
	p.datasourceList = datasourceList
	return p
}

func (p *SaveTeamMetricStrategyParams) WithStrategyMetric(strategyMetric do.StrategyMetric) *SaveTeamMetricStrategyParams {
	p.strategyMetric = strategyMetric
	p.Rules = slices.Map(p.Rules, func(v *SaveTeamMetricStrategyRuleItem) *SaveTeamMetricStrategyRuleItem {
		return v.withMetricStrategy(strategyMetric)
	})
	return p
}

func (p *SaveTeamMetricStrategyParams) WithStrategy(strategy do.Strategy) *SaveTeamMetricStrategyParams {
	p.Strategy.withStrategy(strategy)
	return p
}

func (p *SaveTeamMetricStrategyParams) WithDicts(dictList []do.TeamDict) *SaveTeamMetricStrategyParams {
	if validate.IsNil(p) {
		return p
	}
	if len(p.Rules) == 0 {
		return p
	}
	for _, rule := range p.Rules {
		rule.withDicts(dictList)
	}
	return p
}

func (p *SaveTeamMetricStrategyParams) WithReceiverRoutes(routes []do.NoticeGroup) *SaveTeamMetricStrategyParams {
	if validate.IsNil(p) {
		return p
	}
	for _, rule := range p.Rules {
		rule.withReceiverRoutes(routes)
	}
	p.Strategy.withReceiverRoutes(routes)
	return p
}

func (p *SaveTeamMetricStrategyParams) WithStrategyGroup(strategyGroup do.StrategyGroup) *SaveTeamMetricStrategyParams {
	p.Strategy.withStrategyGroup(strategyGroup)
	return p
}

func (p *SaveTeamMetricStrategyParams) WithLabelNotices(labelNotices []do.StrategyMetricRuleLabelNotice) *SaveTeamMetricStrategyParams {
	if validate.IsNil(p) {
		return p
	}
	labelNoticeMap := slices.ToMap(labelNotices, func(v do.StrategyMetricRuleLabelNotice) uint32 {
		return v.GetID()
	})
	for _, rule := range p.Rules {
		for _, labelNotice := range rule.LabelNotices {
			labelNotice.withLabelNotice(labelNoticeMap[labelNotice.GetID()])
		}
	}
	return p
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
