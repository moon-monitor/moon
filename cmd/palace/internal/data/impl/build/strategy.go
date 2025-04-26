package build

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
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

func ToStrategyDo(ctx context.Context, params bo.SaveTeamStrategy) *team.Strategy {
	if validate.IsNil(params) {
		return nil
	}
	if strategyDo, ok := params.GetStrategy().(*team.Strategy); ok {
		strategyDo.WithContext(ctx)
		return strategyDo
	}
	strategy := &team.Strategy{
		StrategyGroupID: params.GetStrategyGroupID(),
		Name:            params.GetName(),
		TeamModel:       ToTeamModel(params.GetStrategy()),
		Remark:          params.GetRemark(),
		Status:          vobj.GlobalStatusEnable,
		StrategyType:    params.GetStrategyType(),
		StrategyGroup:   ToStrategyGroup(params.GetStrategyGroup()),
		Notices:         ToStrategyNotices(params.GetReceiverRoutes()),
	}
	strategy.WithContext(ctx)
	return strategy
}

func ToStrategyMetricDo(ctx context.Context, params bo.SaveTeamMetricStrategy) *team.StrategyMetric {
	if validate.IsNil(params) {
		return nil
	}
	if strategyMetricDo, ok := params.GetStrategyMetric().(*team.StrategyMetric); ok {
		strategyMetricDo.WithContext(ctx)
		return strategyMetricDo
	}
	strategyMetric := &team.StrategyMetric{
		StrategyID:          params.GetStrategy().GetID(),
		Expr:                params.GetExpr(),
		Labels:              params.GetLabels(),
		Annotations:         params.GetAnnotations(),
		TeamModel:           ToTeamModel(params.GetStrategyMetric()),
		Strategy:            ToStrategy(params.GetStrategy().GetStrategy()),
		StrategyMetricRules: ToStrategyMetricRulesDo(ctx, params.GetRules()),
		Datasource:          ToDatasourceMetrics(params.GetDatasourceList()),
	}
	strategyMetric.WithContext(ctx)
	return strategyMetric
}

func ToStrategyMetricRulesDo(ctx context.Context, params []bo.SaveTeamMetricStrategyRule) []*team.StrategyMetricRule {
	return slices.Map(params, func(v bo.SaveTeamMetricStrategyRule) *team.StrategyMetricRule {
		return ToStrategyMetricRuleDo(ctx, v)
	})
}

func ToStrategyMetricRuleDo(ctx context.Context, params bo.SaveTeamMetricStrategyRule) *team.StrategyMetricRule {
	if validate.IsNil(params) {
		return nil
	}
	if strategyMetricRuleDo, ok := params.GetMetricStrategyRule().(*team.StrategyMetricRule); ok {
		strategyMetricRuleDo.WithContext(ctx)
		return strategyMetricRuleDo
	}
	strategyMetricRule := &team.StrategyMetricRule{
		TeamModel:        ToTeamModel(params.GetMetricStrategy()),
		StrategyMetricID: params.GetMetricStrategy().GetID(),
		StrategyMetric:   ToStrategyMetric(params.GetMetricStrategy()),
		LevelID:          params.GetID(),
		Level:            ToDict(params.GetLevel()),
		SampleMode:       params.GetSampleMode(),
		Condition:        params.GetCondition(),
		Count:            params.GetCount(),
		Values:           params.GetValues(),
		Duration:         params.GetDuration(),
		Status:           params.GetStatus(),
		Notices:          ToStrategyNotices(params.GetReceiverRoutes()),
		LabelNotices:     ToStrategyMetricRuleLabelNoticesDo(params.GetLabelNotices()),
		AlarmPages:       ToDicts(params.GetAlarmPages()),
	}
	strategyMetricRule.WithContext(ctx)
	return strategyMetricRule
}
