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

// Create implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Create(ctx context.Context, params bo.CreateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	panic("unimplemented")
}

// CreateLevels implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) CreateLevels(ctx context.Context, params *bo.SaveTeamMetricStrategyLevelsParams) ([]do.StrategyMetricRule, error) {
	panic("unimplemented")
}

// Delete implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Delete(ctx context.Context, params *bo.OperateTeamStrategyParams) error {
	panic("unimplemented")
}

// Get implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Get(ctx context.Context, params *bo.OperateTeamStrategyParams) (do.StrategyMetric, error) {
	panic("unimplemented")
}

// Update implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) Update(ctx context.Context, params bo.UpdateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	panic("unimplemented")
}

// UpdateLevels implements repository.TeamStrategyMetric.
func (t *teamStrategyMetricImpl) UpdateLevels(ctx context.Context, params *bo.SaveTeamMetricStrategyLevelsParams) ([]do.StrategyMetricRule, error) {
	panic("unimplemented")
}
