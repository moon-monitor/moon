package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeamStrategy(
	teamStrategyRepo repository.TeamStrategy,
) *TeamStrategy {
	return &TeamStrategy{
		teamStrategyRepo: teamStrategyRepo,
	}
}

type TeamStrategy struct {
	teamStrategyRepo repository.TeamStrategy
}

func (t *TeamStrategy) UpdateTeamStrategiesStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error {
	return t.teamStrategyRepo.UpdateStatus(ctx, params)
}

func (t *TeamStrategy) ListTeamStrategy(ctx context.Context, params *bo.ListTeamStrategyParams) (*bo.ListTeamStrategyReply, error) {
	return t.teamStrategyRepo.List(ctx, params)
}

func (t *TeamStrategy) SubscribeTeamStrategy(ctx context.Context, params *bo.ToSubscribeTeamStrategyParams) error {
	return t.teamStrategyRepo.Subscribe(ctx, params)
}

func (t *TeamStrategy) SubscribeTeamStrategies(ctx context.Context, params *bo.ToSubscribeTeamStrategiesParams) (*bo.ToSubscribeTeamStrategiesReply, error) {
	return t.teamStrategyRepo.SubscribeList(ctx, params)
}
