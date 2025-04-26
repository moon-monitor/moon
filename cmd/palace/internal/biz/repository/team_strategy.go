package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

type TeamStrategy interface {
	UpdateStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error
	List(ctx context.Context, params *bo.ListTeamStrategyParams) (*bo.ListTeamStrategyReply, error)
	Subscribe(ctx context.Context, params *bo.ToSubscribeTeamStrategyParams) error
	SubscribeList(ctx context.Context, params *bo.ToSubscribeTeamStrategiesParams) (*bo.ToSubscribeTeamStrategiesReply, error)
	Get(ctx context.Context, strategyID uint32) (do.Strategy, error)
}

type TeamStrategyMetric interface {
	Create(ctx context.Context, params bo.SaveTeamMetricStrategy) error
	Update(ctx context.Context, params bo.SaveTeamMetricStrategy) error
	Get(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error)
	Delete(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error
}
