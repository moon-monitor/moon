package impl

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
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
	strategyNotices := slices.Map(params.GetNotice(params.Strategy.ReceiverRoutes), func(route do.NoticeGroup) *team.NoticeGroup {
		return &team.NoticeGroup{
			ID:            route.GetID(),
			Name:          route.GetName(),
			Remark:        route.GetRemark(),
			Status:        route.GetStatus(),
			CreateTime:    route.GetCreateTime(),
			UpdateTime:    route.GetUpdateTime(),
			TeamModel:     do.TeamModel{},
			Members:       []*team.NoticeMember{},
			Hooks:         []*team.NoticeHook{},
			EmailConfigID: 0,
			EmailConfig:   &team.EmailConfig{},
			SMSConfigID:   0,
			SMSConfig:     &team.SmsConfig{},
		}
	}
	strategyDo := &team.Strategy{
		StrategyGroupID: params.Strategy.StrategyGroupID,
		Name:            params.Strategy.Name,
		Remark:          params.Strategy.Remark,
		Status:          vobj.GlobalStatusEnable,
		StrategyType:    params.Strategy.StrategyType,
		Notices:         strategyNotices,
	}
	strategyMetricDo := &team.StrategyMetric{
		StrategyID:          0,
		Strategy:            strategyDo,
		Expr:                "",
		Labels:              kv.StringMap{},
		Annotations:         kv.StringMap{},
		StrategyMetricRules: []*team.StrategyMetricRule{},
		Datasource:          []*team.DatasourceMetric{},
	}
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
