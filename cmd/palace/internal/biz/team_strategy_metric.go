package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

func NewTeamStrategyMetric() *TeamStrategyMetric {
	return &TeamStrategyMetric{}
}

type TeamStrategyMetric struct {
}

func (t *TeamStrategyMetric) SaveTeamMetricStrategy(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	return nil
}

func (t *TeamStrategyMetric) GetTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return nil, nil
}

func (t *TeamStrategyMetric) DeleteTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return nil
}
