package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeamStrategy(
	teamStrategyGroupRepo repository.TeamStrategyGroup,
	teamStrategyRepo repository.TeamStrategy,
	teamNoticeRepo repository.TeamNotice,
) *TeamStrategy {
	return &TeamStrategy{
		teamStrategyGroupRepo: teamStrategyGroupRepo,
		teamStrategyRepo:      teamStrategyRepo,
		teamNoticeRepo:        teamNoticeRepo,
	}
}

type TeamStrategy struct {
	teamStrategyGroupRepo repository.TeamStrategyGroup
	teamStrategyRepo      repository.TeamStrategy
	teamNoticeRepo        repository.TeamNotice
}

func (t *TeamStrategy) SaveTeamStrategy(ctx context.Context, params *bo.SaveTeamStrategyParams) (do.Strategy, error) {
	strategyGroup, err := t.teamStrategyGroupRepo.Get(ctx, params.StrategyGroupID)
	if err != nil {
		return nil, err
	}
	receiverRoutes, err := t.teamNoticeRepo.FindByIds(ctx, params.ReceiverRoutes)
	if err != nil {
		return nil, err
	}

	if params.ID <= 0 {
		req := params.ToCreateTeamStrategyParams(strategyGroup, receiverRoutes)
		if err := req.Validate(); err != nil {
			return nil, err
		}
		return t.teamStrategyRepo.Create(ctx, req)
	}

	strategyDo, err := t.teamStrategyRepo.Get(ctx, &bo.OperateTeamStrategyParams{StrategyId: params.ID})
	if err != nil {
		return nil, err
	}

	req := params.ToUpdateTeamStrategyParams(strategyGroup, strategyDo, receiverRoutes)
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return t.teamStrategyRepo.Update(ctx, req)
}

func (t *TeamStrategy) DeleteTeamStrategy(ctx context.Context, params *bo.OperateTeamStrategyParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	return t.teamStrategyRepo.Delete(ctx, params)
}

func (t *TeamStrategy) UpdateTeamStrategiesStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	return t.teamStrategyRepo.UpdateStatus(ctx, params)
}

func (t *TeamStrategy) ListTeamStrategy(ctx context.Context, params *bo.ListTeamStrategyParams) (*bo.ListTeamStrategyReply, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return t.teamStrategyRepo.List(ctx, params)
}

func (t *TeamStrategy) SubscribeTeamStrategy(ctx context.Context, params *bo.SubscribeTeamStrategyParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	return t.teamStrategyRepo.Subscribe(ctx, params)
}

func (t *TeamStrategy) SubscribeTeamStrategies(ctx context.Context, params *bo.SubscribeTeamStrategiesParams) (*bo.SubscribeTeamStrategiesReply, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return t.teamStrategyRepo.SubscribeList(ctx, params)
}
