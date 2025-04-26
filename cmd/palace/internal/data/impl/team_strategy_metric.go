package impl

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/impl/build"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func NewTeamStrategyMetricRepo(d *data.Data) repository.TeamStrategyMetric {
	return &teamStrategyMetricImpl{
		Data: d,
	}
}

type teamStrategyMetricImpl struct {
	*data.Data
}

func (r *teamStrategyMetricImpl) Create(ctx context.Context, params bo.SaveTeamMetricStrategy) error {
	tx, _, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return err
	}
	strategyDo := build.ToStrategyDo(ctx, params.GetStrategy())
	strategyMutation := tx.Strategy
	if err := strategyMutation.WithContext(ctx).Create(strategyDo); err != nil {
		return err
	}

	strategyMetricDo := build.ToStrategyMetricDo(ctx, params)
	strategyMetricDo.StrategyID = strategyDo.ID
	strategyMetricDo.Strategy = strategyDo

	strategyMetricMutation := tx.StrategyMetric
	if err := strategyMetricMutation.WithContext(ctx).Create(strategyMetricDo); err != nil {
		return err
	}

	rules := build.ToStrategyMetricRulesDo(ctx, params.GetRules())
	for _, rule := range rules {
		rule.StrategyMetricID = strategyMetricDo.ID
		rule.StrategyMetric = strategyMetricDo
	}
	strategyMetricRuleMutation := tx.StrategyMetricRule
	if err := strategyMetricRuleMutation.WithContext(ctx).Create(rules...); err != nil {
		return err
	}

	return nil
}

func (r *teamStrategyMetricImpl) Update(ctx context.Context, params bo.SaveTeamMetricStrategy) error {
	tx, teamId, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return err
	}
	strategyDo := build.ToStrategyDo(ctx, params.GetStrategy())
	strategyMutation := tx.Strategy
	strategyMutations := []field.AssignExpr{
		strategyMutation.Name.Value(strategyDo.Name),
		strategyMutation.Remark.Value(strategyDo.Remark),
		strategyMutation.Status.Value(strategyDo.Status.GetValue()),
		strategyMutation.StrategyType.Value(strategyDo.StrategyType.GetValue()),
		strategyMutation.StrategyGroupID.Value(strategyDo.GetStrategyGroupID()),
	}
	strategyWrapper := []gen.Condition{
		strategyMutation.TeamID.Eq(teamId),
		strategyMutation.ID.Eq(params.GetStrategy().GetID()),
	}
	if _, err := strategyMutation.WithContext(ctx).Where(strategyWrapper...).UpdateSimple(strategyMutations...); err != nil {
		return err
	}
	strategyMetricDo := build.ToStrategyMetricDo(ctx, params)
	strategyMetricMutation := tx.StrategyMetric
	strategyMetricWrapper := []gen.Condition{
		strategyMetricMutation.TeamID.Eq(teamId),
		strategyMetricMutation.ID.Eq(params.GetStrategyMetric().GetID()),
	}
	strategyMetricMutations := []field.AssignExpr{
		strategyMetricMutation.Expr.Value(strategyMetricDo.Expr),
		strategyMetricMutation.Labels.Value(strategyMetricDo.Labels),
		strategyMetricMutation.Annotations.Value(strategyMetricDo.Annotations),
	}
	if _, err := strategyMetricMutation.WithContext(ctx).Where(strategyMetricWrapper...).UpdateSimple(strategyMetricMutations...); err != nil {
		return err
	}
	strategyMetricRuleMutation := tx.StrategyMetricRule
	strategyMetricRuleDo := build.ToStrategyMetricRulesDo(ctx, params.GetRules())
	for _, rule := range strategyMetricRuleDo {
		rule.StrategyMetricID = strategyMetricDo.ID
		rule.StrategyMetric = strategyMetricDo
	}
	if err := strategyMetricRuleMutation.
		WithContext(ctx).
		Where(strategyMetricWrapper...).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: strategyMetricRuleMutation.ID.ColumnName().String()}},
			UpdateAll: true,
		}).
		Create(strategyMetricRuleDo...); err != nil {
		return err
	}
	return nil
}

func (r *teamStrategyMetricImpl) Get(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return nil, nil
}

func (r *teamStrategyMetricImpl) Delete(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return nil
}
