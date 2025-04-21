package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func NewTeamStrategyService() *TeamStrategyService {
	return &TeamStrategyService{}
}

type TeamStrategyService struct {
	palacev1.UnimplementedTeamStrategyServer
}

func (t *TeamStrategyService) SaveTeamStrategyGroup(ctx context.Context, request *palacev1.SaveTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) UpdateTeamStrategyGroupStatus(ctx context.Context, request *palacev1.UpdateTeamStrategyGroupStatusRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) DeleteTeamStrategyGroup(ctx context.Context, request *palacev1.DeleteTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) GetTeamStrategyGroup(ctx context.Context, request *palacev1.GetTeamStrategyGroupRequest) (*palacev1.GetTeamStrategyGroupReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) ListTeamStrategyGroup(ctx context.Context, request *palacev1.ListTeamStrategyGroupRequest) (*palacev1.ListTeamStrategyGroupReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) SaveTeamStrategy(ctx context.Context, request *palacev1.SaveTeamStrategyRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) UpdateTeamStrategiesStatus(ctx context.Context, request *palacev1.UpdateTeamStrategiesStatusRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) DeleteTeamStrategy(ctx context.Context, request *palacev1.DeleteTeamStrategyRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) GetTeamStrategy(ctx context.Context, request *palacev1.GetTeamStrategyRequest) (*palacev1.GetTeamStrategyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) ListTeamStrategy(ctx context.Context, request *palacev1.ListTeamStrategyRequest) (*palacev1.ListTeamStrategyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) SubscribeTeamStrategy(ctx context.Context, request *palacev1.SubscribeTeamStrategyRequest) (*common.EmptyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeamStrategyService) SubscribeTeamStrategies(ctx context.Context, request *palacev1.SubscribeTeamStrategiesRequest) (*palacev1.SubscribeTeamStrategiesReply, error) {
	//TODO implement me
	panic("implement me")
}
