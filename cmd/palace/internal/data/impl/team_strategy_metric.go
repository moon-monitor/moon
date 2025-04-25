package impl

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
)

func NewTeamStrategyMetricRepo(d *data.Data) repository.TeamStrategyMetric {
	return &teamStrategyMetricImpl{
		Data: d,
	}
}

type teamStrategyMetricImpl struct {
	*data.Data
}

func (r *teamStrategyMetricImpl) Create(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	return nil
}

func (r *teamStrategyMetricImpl) Update(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	return nil
}

func (r *teamStrategyMetricImpl) Get(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return nil, nil
}

func (r *teamStrategyMetricImpl) Delete(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return nil
}
