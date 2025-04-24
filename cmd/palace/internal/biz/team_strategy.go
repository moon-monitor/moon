package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
)

func NewTeamStrategy() *TeamStrategy {
	return &TeamStrategy{}
}

type TeamStrategy struct {
}

func (t *TeamStrategy) UpdateTeamStrategiesStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error {
	return nil
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
