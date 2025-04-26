package impl

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/impl/build"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewTeamStrategyMetricRepo(d *data.Data) repository.TeamStrategyMetric {
	return &teamStrategyMetricImpl{
		Data: d,
	}
}

type teamStrategyMetricImpl struct {
	*data.Data
}

// Create implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Create(ctx context.Context, params bo.CreateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	tx, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	strategyMetricDo := &team.StrategyMetric{
		StrategyID:  params.GetStrategy().GetID(),
		Strategy:    build.ToStrategy(params.GetStrategy()),
		Expr:        params.GetExpr(),
		Labels:      params.GetLabels(),
		Annotations: params.GetAnnotations(),
		Datasource:  build.ToDatasourceMetrics(params.GetDatasource()),
	}

	if err = tx.StrategyMetric.WithContext(ctx).Create(strategyMetricDo); err != nil {
		return nil, err
	}

	if len(strategyMetricDo.Datasource) > 0 {
		datasource := tx.StrategyMetric.Datasource.WithContext(ctx).Model(strategyMetricDo)
		if err := datasource.Append(strategyMetricDo.Datasource...); err != nil {
			return nil, err
		}
	}

	return strategyMetricDo, nil
}

// Update implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Update(ctx context.Context, params bo.UpdateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	strategyMetricMutation := tx.StrategyMetric
	wrapper := []gen.Condition{
		strategyMetricMutation.StrategyID.Eq(params.GetStrategy().GetID()),
		strategyMetricMutation.TeamID.Eq(teamId),
	}

	strategyMetricMutations := []field.AssignExpr{
		strategyMetricMutation.Expr.Value(params.GetExpr()),
		strategyMetricMutation.Labels.Value(params.GetLabels()),
		strategyMetricMutation.Annotations.Value(params.GetAnnotations()),
	}
	if _, err := strategyMetricMutation.WithContext(ctx).Where(wrapper...).UpdateSimple(strategyMetricMutations...); err != nil {
		return nil, err
	}
	strategyMetricDo, err := t.Get(ctx, &bo.OperateTeamStrategyParams{
		StrategyId: params.GetStrategyMetric().GetID(),
	})
	if err != nil {
		return nil, err
	}

	datasourceDos := build.ToDatasourceMetrics(params.GetDatasource())
	datasourceMutation := tx.StrategyMetric.Datasource.WithContext(ctx).Model(build.ToStrategyMetric(strategyMetricDo))
	if len(datasourceDos) > 0 {
		if err := datasourceMutation.Replace(datasourceDos...); err != nil {
			return nil, err
		}
	} else {
		if err := datasourceMutation.Clear(); err != nil {
			return nil, err
		}
	}

	return strategyMetricDo, nil
}

// UpdateLevels implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) UpdateLevels(ctx context.Context, params bo.SaveTeamMetricStrategyLevels) ([]do.StrategyMetricRule, error) {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	strategyMetricRuleMutation := tx.StrategyMetricRule
	wrapper := []gen.Condition{
		strategyMetricRuleMutation.StrategyMetricID.Eq(params.GetStrategyMetric().GetID()),
		strategyMetricRuleMutation.TeamID.Eq(teamId),
	}
	strategyMetricRuleDos := make([]do.StrategyMetricRule, 0, len(params.GetLevels()))
	for _, rule := range params.GetLevels() {
		mutations := []field.AssignExpr{
			strategyMetricRuleMutation.LevelID.Value(rule.GetLevel().GetID()),
			strategyMetricRuleMutation.SampleMode.Value(rule.GetSampleMode().GetValue()),
			strategyMetricRuleMutation.Count_.Value(rule.GetCount()),
			strategyMetricRuleMutation.Condition.Value(rule.GetCondition().GetValue()),
			strategyMetricRuleMutation.Values.Value(team.Values(rule.GetValues())),
			strategyMetricRuleMutation.Status.Value(rule.GetStatus().GetValue()),
			strategyMetricRuleMutation.Duration.Value(int64(rule.GetDuration())),
		}
		if _, err := strategyMetricRuleMutation.WithContext(ctx).Where(wrapper...).UpdateSimple(mutations...); err != nil {
			return nil, err
		}
		ruleDo, err := strategyMetricRuleMutation.WithContext(ctx).Where(wrapper...).First()
		if err != nil {
			return nil, err
		}
		strategyMetricRuleDos = append(strategyMetricRuleDos, ruleDo)
		ruleDoItem := build.ToStrategyMetricRule(ruleDo)
		ruleDoItem.WithContext(ctx)
		alarmPages := build.ToDicts(rule.GetAlarmPages())
		alarmPagesMutation := tx.StrategyMetricRule.AlarmPages.WithContext(ctx).Model(ruleDoItem)
		if len(alarmPages) > 0 {
			if err := alarmPagesMutation.Replace(alarmPages...); err != nil {
				return nil, err
			}
		} else {
			if err := alarmPagesMutation.Clear(); err != nil {
				return nil, err
			}
		}
		noticeGroups := build.ToStrategyNotices(rule.GetReceiverRoutes())
		noticeGroupsMutation := tx.StrategyMetricRule.Notices.WithContext(ctx).Model(ruleDoItem)
		if len(noticeGroups) > 0 {
			if err := noticeGroupsMutation.Replace(noticeGroups...); err != nil {
				return nil, err
			}
		} else {
			if err := noticeGroupsMutation.Clear(); err != nil {
				return nil, err
			}
		}
		labelNotices := slices.Map(rule.GetLabelNotices(), func(notice bo.LabelNotice) *team.StrategyMetricRuleLabelNotice {
			labelNoticeDo := &team.StrategyMetricRuleLabelNotice{
				LabelKey:             notice.GetKey(),
				LabelValue:           notice.GetValue(),
				Notices:              build.ToStrategyNotices(notice.GetReceiverRoutes()),
				StrategyMetricRuleID: ruleDoItem.GetID(),
			}
			labelNoticeDo.WithContext(ctx)
			return labelNoticeDo
		})
		strategyMetricRuleLabelNoticeMutation := tx.StrategyMetricRuleLabelNotice
		strategyMetricRuleLabelNoticeWrapper := []gen.Condition{
			strategyMetricRuleLabelNoticeMutation.StrategyMetricRuleID.Eq(ruleDoItem.GetID()),
		}
		if _, err := strategyMetricRuleLabelNoticeMutation.WithContext(ctx).Where(strategyMetricRuleLabelNoticeWrapper...).Delete(); err != nil {
			return nil, err
		}
		if err := strategyMetricRuleLabelNoticeMutation.WithContext(ctx).Create(labelNotices...); err != nil {
			return nil, err
		}
	}
	return strategyMetricRuleDos, nil
}

// CreateLevels implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) CreateLevels(ctx context.Context, params bo.SaveTeamMetricStrategyLevels) ([]do.StrategyMetricRule, error) {
	tx, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	strategyMetricRuleDos := slices.Map(params.GetLevels(), func(rule bo.SaveTeamMetricStrategyLevel) *team.StrategyMetricRule {
		ruleItem := &team.StrategyMetricRule{
			StrategyMetricID: params.GetStrategyMetric().GetID(),
			LevelID:          rule.GetLevel().GetID(),
			SampleMode:       rule.GetSampleMode(),
			Count:            rule.GetCount(),
			Condition:        rule.GetCondition(),
			Values:           rule.GetValues(),
			StrategyMetric:   build.ToStrategyMetric(params.GetStrategyMetric()),
			Level:            build.ToDict(rule.GetLevel()),
			Duration:         rule.GetDuration(),
			Status:           rule.GetStatus(),
			Notices:          build.ToStrategyNotices(rule.GetReceiverRoutes()),
			LabelNotices: slices.Map(rule.GetLabelNotices(), func(notice bo.LabelNotice) *team.StrategyMetricRuleLabelNotice {
				item := &team.StrategyMetricRuleLabelNotice{
					LabelKey:   notice.GetKey(),
					LabelValue: notice.GetValue(),
					Notices:    build.ToStrategyNotices(notice.GetReceiverRoutes()),
				}
				item.WithContext(ctx)
				return item
			}),
			AlarmPages: build.ToDicts(rule.GetAlarmPages()),
		}
		ruleItem.WithContext(ctx)
		return ruleItem
	})

	if err = tx.StrategyMetricRule.WithContext(ctx).Create(strategyMetricRuleDos...); err != nil {
		return nil, err
	}

	return slices.Map(strategyMetricRuleDos, func(rule *team.StrategyMetricRule) do.StrategyMetricRule {
		return rule
	}), nil
}

// Delete implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Delete(ctx context.Context, params *bo.OperateTeamStrategyParams) error {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}

	strategyMetricMutation := tx.StrategyMetric
	wrapper := []gen.Condition{
		strategyMetricMutation.StrategyID.Eq(params.StrategyId),
		strategyMetricMutation.TeamID.Eq(teamId),
	}

	_, err = strategyMetricMutation.WithContext(ctx).Where(wrapper...).Delete()
	return err
}

// DeleteLevels implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) DeleteLevels(ctx context.Context, params *bo.OperateTeamStrategyParams) error {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}

	strategyMetricMutation := tx.StrategyMetric
	wrapper := []gen.Condition{
		strategyMetricMutation.StrategyID.Eq(params.StrategyId),
		strategyMetricMutation.TeamID.Eq(teamId),
	}
	var strategyMetricIds []uint32
	if err := strategyMetricMutation.WithContext(ctx).Select(strategyMetricMutation.ID).Where(wrapper...).Scan(&strategyMetricIds); err != nil {
		return err
	}
	if len(strategyMetricIds) == 0 {
		return nil
	}
	strategyMetricRuleMutation := tx.StrategyMetricRule
	wrapper = []gen.Condition{
		strategyMetricRuleMutation.TeamID.Eq(teamId),
		strategyMetricRuleMutation.StrategyMetricID.In(strategyMetricIds...),
	}
	_, err = strategyMetricRuleMutation.WithContext(ctx).Where(wrapper...).Delete()
	return err
}

// Get implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Get(ctx context.Context, params *bo.OperateTeamStrategyParams) (do.StrategyMetric, error) {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	strategyMetricMutation := tx.StrategyMetric
	wrapper := []gen.Condition{
		strategyMetricMutation.StrategyID.Eq(params.StrategyId),
		strategyMetricMutation.TeamID.Eq(teamId),
	}

	strategyMetricDo, err := strategyMetricMutation.WithContext(ctx).Preload(field.Associations).Where(wrapper...).First()
	if err != nil {
		return nil, err
	}

	return build.ToStrategyMetric(strategyMetricDo), nil
}
