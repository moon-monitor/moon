package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewTeamStrategyService(
	teamStrategyGroupBiz *biz.TeamStrategyGroupBiz,
	teamStrategyBiz *biz.TeamStrategy,
	teamStrategyMetricBiz *biz.TeamStrategyMetric,
) *TeamStrategyService {
	return &TeamStrategyService{
		teamStrategyGroupBiz:  teamStrategyGroupBiz,
		teamStrategyBiz:       teamStrategyBiz,
		teamStrategyMetricBiz: teamStrategyMetricBiz,
	}
}

type TeamStrategyService struct {
	palacev1.UnimplementedTeamStrategyServer
	teamStrategyGroupBiz  *biz.TeamStrategyGroupBiz
	teamStrategyBiz       *biz.TeamStrategy
	teamStrategyMetricBiz *biz.TeamStrategyMetric
}

func (t *TeamStrategyService) SaveTeamStrategyGroup(ctx context.Context, request *palacev1.SaveTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	params := &bo.SaveTeamStrategyGroupParams{
		ID:     request.GetGroupId(),
		Name:   request.GetName(),
		Remark: request.GetRemark(),
	}
	if err := t.teamStrategyGroupBiz.SaveTeamStrategyGroup(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存策略组成功"}, nil
}

func (t *TeamStrategyService) UpdateTeamStrategyGroupStatus(ctx context.Context, request *palacev1.UpdateTeamStrategyGroupStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateTeamStrategyGroupStatusParams{
		ID:     request.GetGroupId(),
		Status: vobj.GlobalStatus(request.GetStatus()),
	}
	if err := t.teamStrategyGroupBiz.UpdateTeamStrategyGroupStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新策略组状态成功"}, nil
}

func (t *TeamStrategyService) DeleteTeamStrategyGroup(ctx context.Context, request *palacev1.DeleteTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	if err := t.teamStrategyGroupBiz.DeleteTeamStrategyGroup(ctx, request.GetGroupId()); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "删除策略组成功"}, nil
}

func (t *TeamStrategyService) GetTeamStrategyGroup(ctx context.Context, request *palacev1.GetTeamStrategyGroupRequest) (*common.TeamStrategyGroupItem, error) {
	group, err := t.teamStrategyGroupBiz.GetTeamStrategyGroup(ctx, request.GetGroupId())
	if err != nil {
		return nil, err
	}
	return build.ToTeamStrategyGroupItem(group), nil
}

func (t *TeamStrategyService) ListTeamStrategyGroup(ctx context.Context, request *palacev1.ListTeamStrategyGroupRequest) (*palacev1.ListTeamStrategyGroupReply, error) {
	params := &bo.ListTeamStrategyGroupParams{
		Keyword:           request.GetKeyword(),
		Status:            slices.Map(request.GetStatus(), func(status common.GlobalStatus) vobj.GlobalStatus { return vobj.GlobalStatus(status) }),
		PaginationRequest: build.ToPaginationRequest(request.GetPagination()),
	}
	groups, err := t.teamStrategyGroupBiz.ListTeamStrategyGroup(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.ListTeamStrategyGroupReply{
		Items:      build.ToTeamStrategyGroupItems(groups.Items),
		Pagination: build.ToPaginationReply(groups.PaginationReply),
	}, nil
}

func (t *TeamStrategyService) SaveTeamStrategy(ctx context.Context, request *palacev1.SaveTeamStrategyRequest) (*common.EmptyReply, error) {
	params := build.ToSaveTeamStrategyParams(request)
	strategyDo, err := t.teamStrategyBiz.SaveTeamStrategy(ctx, params)
	if err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存策略成功", Id: strategyDo.GetID()}, nil
}

func (t *TeamStrategyService) SaveTeamMetricStrategy(ctx context.Context, request *palacev1.SaveTeamMetricStrategyRequest) (*common.EmptyReply, error) {
	params := build.ToSaveTeamMetricStrategyParams(request)
	metricStrategyDo, err := t.teamStrategyMetricBiz.SaveTeamMetricStrategy(ctx, params)
	if err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存策略成功", Id: metricStrategyDo.GetID()}, nil
}

func (t *TeamStrategyService) SaveTeamMetricStrategyLevels(ctx context.Context, request *palacev1.SaveTeamMetricStrategyLevelsRequest) (*palacev1.SaveTeamMetricStrategyLevelsReply, error) {
	params := build.ToSaveTeamMetricStrategyLevelsParams(request)
	metricStrategyRules, err := t.teamStrategyMetricBiz.SaveTeamMetricStrategyLevels(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.SaveTeamMetricStrategyLevelsReply{
		Levels:  build.ToTeamMetricStrategyItemRules(metricStrategyRules),
		Message: "保存策略成功",
	}, nil
}

func (t *TeamStrategyService) UpdateTeamStrategiesStatus(ctx context.Context, request *palacev1.UpdateTeamStrategiesStatusRequest) (*common.EmptyReply, error) {
	params := build.ToUpdateTeamStrategiesStatusParams(request)
	if err := t.teamStrategyBiz.UpdateTeamStrategiesStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新策略状态成功"}, nil
}

func (t *TeamStrategyService) DeleteTeamStrategy(ctx context.Context, request *palacev1.OperateTeamStrategyRequest) (*common.EmptyReply, error) {
	params := build.ToOperateTeamStrategyParams(request)
	if err := t.teamStrategyBiz.DeleteTeamStrategy(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "删除策略成功"}, nil
}

func (t *TeamStrategyService) GetTeamMetricStrategy(ctx context.Context, request *palacev1.OperateTeamStrategyRequest) (*common.TeamStrategyMetricItem, error) {
	params := build.ToOperateTeamStrategyParams(request)
	strategy, err := t.teamStrategyMetricBiz.GetTeamMetricStrategy(ctx, params)
	if err != nil {
		return nil, err
	}
	return build.ToTeamMetricStrategyItem(strategy), nil
}

func (t *TeamStrategyService) ListTeamStrategy(ctx context.Context, request *palacev1.ListTeamStrategyRequest) (*palacev1.ListTeamStrategyReply, error) {
	params := build.ToListTeamStrategyParams(request)
	strategies, err := t.teamStrategyBiz.ListTeamStrategy(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.ListTeamStrategyReply{
		Items:      build.ToTeamStrategyItems(strategies.Items),
		Pagination: build.ToPaginationReply(strategies.PaginationReply),
	}, nil
}

func (t *TeamStrategyService) SubscribeTeamStrategy(ctx context.Context, request *palacev1.SubscribeTeamStrategyRequest) (*common.EmptyReply, error) {
	params := build.ToSubscribeTeamStrategyParams(request)
	if err := t.teamStrategyBiz.SubscribeTeamStrategy(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "订阅策略成功"}, nil
}

func (t *TeamStrategyService) SubscribeTeamStrategies(ctx context.Context, request *palacev1.SubscribeTeamStrategiesRequest) (*palacev1.SubscribeTeamStrategiesReply, error) {
	params := build.ToSubscribeTeamStrategiesParams(request)
	subscribers, err := t.teamStrategyBiz.SubscribeTeamStrategies(ctx, params)
	if err != nil {
		return nil, err
	}
	return &palacev1.SubscribeTeamStrategiesReply{
		Items:      build.ToSubscribeTeamStrategiesItems(subscribers.Items),
		Pagination: build.ToPaginationReply(subscribers.PaginationReply),
	}, nil
}
