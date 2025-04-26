package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToStrategyGroup(strategyGroup do.StrategyGroup) *team.StrategyGroup {
	if validate.IsNil(strategyGroup) {
		return nil
	}
	if strategyGroup, ok := strategyGroup.(*team.StrategyGroup); ok {
		return strategyGroup
	}
	return &team.StrategyGroup{
		Name:       strategyGroup.GetName(),
		TeamModel:  ToTeamModel(strategyGroup),
		Remark:     strategyGroup.GetRemark(),
		Status:     vobj.GlobalStatusEnable,
		Strategies: ToStrategies(strategyGroup.GetStrategies()),
	}
}

func ToStrategy(params do.Strategy) *team.Strategy {
	if validate.IsNil(params) {
		return nil
	}
	if strategy, ok := params.(*team.Strategy); ok {
		return strategy
	}
	return &team.Strategy{
		StrategyGroupID: params.GetStrategyGroupID(),
		Name:            params.GetName(),
		Remark:          params.GetRemark(),
		Status:          vobj.GlobalStatusEnable,
		StrategyType:    params.GetStrategyType(),
		Notices:         ToStrategyNotices(params.GetNotices()),
		TeamModel:       ToTeamModel(params),
		StrategyGroup:   ToStrategyGroup(params.GetStrategyGroup()),
	}
}

func ToStrategies(params []do.Strategy) []*team.Strategy {
	return slices.Map(params, ToStrategy)
}

func ToStrategyMetric(params do.StrategyMetric) *team.StrategyMetric {
	if validate.IsNil(params) {
		return nil
	}
	if strategyMetric, ok := params.(*team.StrategyMetric); ok {
		return strategyMetric
	}
	return &team.StrategyMetric{
		StrategyID:          params.GetID(),
		Expr:                params.GetExpr(),
		Labels:              params.GetLabels(),
		Annotations:         params.GetAnnotations(),
		TeamModel:           ToTeamModel(params),
		Strategy:            ToStrategy(params.GetStrategy()),
		StrategyMetricRules: ToStrategyMetricRules(params.GetRules()),
		Datasource:          ToDatasourceMetrics(params.GetDatasourceList()),
	}
}

func ToStrategyMetricRules(params []do.StrategyMetricRule) []*team.StrategyMetricRule {
	return slices.Map(params, ToStrategyMetricRule)
}

func ToStrategyMetricRule(params do.StrategyMetricRule) *team.StrategyMetricRule {
	if validate.IsNil(params) {
		return nil
	}
	if strategyMetricRule, ok := params.(*team.StrategyMetricRule); ok {
		return strategyMetricRule
	}
	return &team.StrategyMetricRule{
		TeamModel:        ToTeamModel(params),
		StrategyMetricID: params.GetID(),
		StrategyMetric:   ToStrategyMetric(params.GetStrategyMetric()),
		LevelID:          params.GetLevelID(),
		Level:            ToDict(params.GetLevel()),
		SampleMode:       params.GetSampleMode(),
		Condition:        params.GetCondition(),
		Count:            params.GetCount(),
		Values:           params.GetValues(),
		Duration:         params.GetDuration(),
		Status:           params.GetStatus(),
		Notices:          ToStrategyNotices(params.GetNotices()),
		LabelNotices:     ToStrategyMetricRuleLabelNotices(params.GetLabelNotices()),
		AlarmPages:       ToDicts(params.GetAlarmPages()),
	}
}
