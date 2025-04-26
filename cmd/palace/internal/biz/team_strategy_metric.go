package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeamStrategyMetric(
	teamStrategyGroupRepo repository.TeamStrategyGroup,
	teamStrategyRepo repository.TeamStrategy,
	teamStrategyMetricRepo repository.TeamStrategyMetric,
	dictRepo repository.TeamDict,
	noticeGroupRepo repository.TeamNotice,
	datasourceRepo repository.TeamDatasourceMetric,
) *TeamStrategyMetric {
	return &TeamStrategyMetric{
		teamStrategyGroupRepo:  teamStrategyGroupRepo,
		teamStrategyRepo:       teamStrategyRepo,
		teamStrategyMetricRepo: teamStrategyMetricRepo,
		dictRepo:               dictRepo,
		noticeGroupRepo:        noticeGroupRepo,
		datasourceRepo:         datasourceRepo,
	}
}

type TeamStrategyMetric struct {
	teamStrategyGroupRepo  repository.TeamStrategyGroup
	teamStrategyRepo       repository.TeamStrategy
	teamStrategyMetricRepo repository.TeamStrategyMetric
	dictRepo               repository.TeamDict
	noticeGroupRepo        repository.TeamNotice
	datasourceRepo         repository.TeamDatasourceMetric
}

func (t *TeamStrategyMetric) SaveTeamMetricStrategy(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) (do.StrategyMetric, error) {
	strategyDo, err := t.teamStrategyRepo.Get(ctx, &bo.OperateTeamStrategyParams{StrategyId: params.StrategyID})
	if err != nil {
		return nil, err
	}
	datasourceDos, err := t.datasourceRepo.FindByIds(ctx, params.Datasource)
	if err != nil {
		return nil, err
	}
	if params.ID <= 0 {
		req := params.ToCreateTeamMetricStrategyParams(strategyDo, datasourceDos)
		if err := req.Validate(); err != nil {
			return nil, err
		}
		return t.teamStrategyMetricRepo.Create(ctx, req)
	}
	strategyMetricDo, err := t.teamStrategyMetricRepo.Get(ctx, &bo.OperateTeamStrategyParams{StrategyId: params.StrategyID})
	if err != nil {
		return nil, err
	}
	req := params.ToUpdateTeamMetricStrategyParams(strategyDo, datasourceDos, strategyMetricDo)
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return t.teamStrategyMetricRepo.Update(ctx, req)
}

func (t *TeamStrategyMetric) SaveTeamMetricStrategyLevels(ctx context.Context, params *bo.SaveTeamMetricStrategyLevelsParams) ([]do.StrategyMetricRule, error) {
	return nil, nil
}

func (t *TeamStrategyMetric) GetTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamStrategyParams) (do.StrategyMetric, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return t.teamStrategyMetricRepo.Get(ctx, params)
}
