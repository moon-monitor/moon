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
	transaction repository.Transaction,
) *TeamStrategyMetric {
	return &TeamStrategyMetric{
		teamStrategyGroupRepo:  teamStrategyGroupRepo,
		teamStrategyRepo:       teamStrategyRepo,
		teamStrategyMetricRepo: teamStrategyMetricRepo,
		dictRepo:               dictRepo,
		noticeGroupRepo:        noticeGroupRepo,
		datasourceRepo:         datasourceRepo,
		transaction:            transaction,
	}
}

type TeamStrategyMetric struct {
	teamStrategyGroupRepo  repository.TeamStrategyGroup
	teamStrategyRepo       repository.TeamStrategy
	teamStrategyMetricRepo repository.TeamStrategyMetric
	dictRepo               repository.TeamDict
	noticeGroupRepo        repository.TeamNotice
	datasourceRepo         repository.TeamDatasourceMetric
	transaction            repository.Transaction
}

func (t *TeamStrategyMetric) SaveTeamMetricStrategy(ctx context.Context, params *bo.SaveTeamMetricStrategyParams) error {
	datasourceDoList, err := t.datasourceRepo.FindByIds(ctx, params.DatasourceList)
	if err != nil {
		return err
	}
	params.WithDatasourceList(datasourceDoList)
	dictDoList, err := t.dictRepo.FindByIds(ctx, params.GetDictIds())
	if err != nil {
		return err
	}
	params.WithDicts(dictDoList)
	noticeGroupDoList, err := t.noticeGroupRepo.FindByIds(ctx, params.GetReceiverIds())
	if err != nil {
		return err
	}
	params.WithReceiverRoutes(noticeGroupDoList)
	labelNotices, err := t.noticeGroupRepo.FindLabelNotices(ctx, params.GetLabelNoticeIds())
	if err != nil {
		return err
	}
	params.WithLabelNotices(labelNotices)
	return t.transaction.BizExec(ctx, func(ctx context.Context) error {
		if params.GetID() <= 0 {
			return t.teamStrategyMetricRepo.Create(ctx, params)
		}
		strategyDo, err := t.teamStrategyRepo.Get(ctx, params.GetStrategy().GetID())
		if err != nil {
			return err
		}
		params.WithStrategy(strategyDo)
		stratgyGroup, err := t.teamStrategyGroupRepo.Get(ctx, params.GetStrategy().GetStrategyGroupID())
		if err != nil {
			return err
		}
		params.WithStrategyGroup(stratgyGroup)
		return t.teamStrategyMetricRepo.Update(ctx, params)
	})
}

func (t *TeamStrategyMetric) GetTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) (do.StrategyMetric, error) {
	return t.teamStrategyMetricRepo.Get(ctx, params)
}

func (t *TeamStrategyMetric) DeleteTeamMetricStrategy(ctx context.Context, params *bo.OperateTeamMetricStrategyParams) error {
	return t.teamStrategyMetricRepo.Delete(ctx, params)
}
