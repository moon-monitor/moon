package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeamStrategyMetric(
	teamStrategyRepo repository.TeamStrategy,
	teamStrategyMetricRepo repository.TeamStrategyMetric,
	dictRepo repository.TeamDict,
	noticeGroupRepo repository.TeamNotice,
	datasourceRepo repository.TeamDatasourceMetric,
	transaction repository.Transaction,
) *TeamStrategyMetric {
	return &TeamStrategyMetric{
		teamStrategyRepo:       teamStrategyRepo,
		teamStrategyMetricRepo: teamStrategyMetricRepo,
		dictRepo:               dictRepo,
		noticeGroupRepo:        noticeGroupRepo,
		datasourceRepo:         datasourceRepo,
		transaction:            transaction,
	}
}

type TeamStrategyMetric struct {
	teamStrategyRepo       repository.TeamStrategy
	teamStrategyMetricRepo repository.TeamStrategyMetric
	dictRepo               repository.TeamDict
	noticeGroupRepo        repository.TeamNotice
	datasourceRepo         repository.TeamDatasourceMetric
	transaction            repository.Transaction
}

func (t *TeamStrategyMetric) SaveTeamMetricStrategy(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	strategyDo, err := t.teamStrategyRepo.Get(ctx, params.Strategy.StrategyID)
	if err != nil {
		return err
	}
	params.WithStrategy(strategyDo)
	levelIds := params.GetLevelIds()
	levels, err := t.dictRepo.FindByIds(ctx, levelIds)
	if err != nil {
		return err
	}
	params.WithLevels(levels)
	receiverRoutes, err := t.noticeGroupRepo.FindByIds(ctx, params.GetReceiverRouteIds())
	if err != nil {
		return err
	}
	params.WithReceiverRoutes(receiverRoutes)
	datasourceList, err := t.datasourceRepo.FindByIds(ctx, params.DatasourceList)
	if err != nil {
		return err
	}
	params.WithDatasourceList(datasourceList)
	if err := params.Validate(); err != nil {
		return err
	}
	return t.transaction.BizExec(ctx, func(ctx context.Context) error {
		if params.Strategy.StrategyID > 0 {
			return t.teamStrategyMetricRepo.Update(ctx, params)
		}
		return t.teamStrategyMetricRepo.Create(ctx, params)
	})
}

func (t *TeamStrategyMetric) GetTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return t.teamStrategyMetricRepo.Get(ctx, params)
}

func (t *TeamStrategyMetric) DeleteTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return t.teamStrategyMetricRepo.Delete(ctx, params)
}
