package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

func NewTeamStrategy() *TeamStrategy {
	return &TeamStrategy{}
}

type TeamStrategy struct {
}

func (t *TeamStrategy) SaveTeamMetricStrategy(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	return nil
}

func (t *TeamStrategy) UpdateTeamStrategiesStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error {
	return nil
}

func (t *TeamStrategy) DeleteTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return nil
}

func (t *TeamStrategy) GetTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return nil, nil
}

func (t *TeamStrategy) ListTeamStrategy(ctx context.Context, params *bo.ListTeamStrategyParams) (*bo.ListTeamStrategyReply, error) {
	return nil, nil
}

func (t *TeamStrategy) SubscribeTeamStrategy(ctx context.Context, params *bo.ToSubscribeTeamStrategyParams) error {
	return nil
}

func (t *TeamStrategy) SubscribeTeamStrategies(ctx context.Context, params *bo.ToSubscribeTeamStrategiesParams) (*bo.ToSubscribeTeamStrategiesReply, error) {
	return nil, nil
}
