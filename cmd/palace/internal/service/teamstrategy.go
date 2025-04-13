package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamStrategyService struct {
	palacev1.UnimplementedTeamStrategyServer
}

func NewTeamStrategyService() *TeamStrategyService {
	return &TeamStrategyService{}
}

func (s *TeamStrategyService) SaveTeamStrategyGroup(ctx context.Context, req *palacev1.SaveTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) UpdateTeamStrategyGroupStatus(ctx context.Context, req *palacev1.UpdateTeamStrategyGroupStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) DeleteTeamStrategyGroup(ctx context.Context, req *palacev1.DeleteTeamStrategyGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) GetTeamStrategyGroup(ctx context.Context, req *palacev1.GetTeamStrategyGroupRequest) (*palacev1.GetTeamStrategyGroupReply, error) {
	return &palacev1.GetTeamStrategyGroupReply{}, nil
}
func (s *TeamStrategyService) ListTeamStrategyGroup(ctx context.Context, req *palacev1.ListTeamStrategyGroupRequest) (*palacev1.ListTeamStrategyGroupReply, error) {
	return &palacev1.ListTeamStrategyGroupReply{}, nil
}
func (s *TeamStrategyService) SaveTeamStrategy(ctx context.Context, req *palacev1.SaveTeamStrategyRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) UpdateTeamStrategiesStatus(ctx context.Context, req *palacev1.UpdateTeamStrategiesStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) DeleteTeamStrategy(ctx context.Context, req *palacev1.DeleteTeamStrategyRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) GetTeamStrategy(ctx context.Context, req *palacev1.GetTeamStrategyRequest) (*palacev1.GetTeamStrategyReply, error) {
	return &palacev1.GetTeamStrategyReply{}, nil
}
func (s *TeamStrategyService) ListTeamStrategy(ctx context.Context, req *palacev1.ListTeamStrategyRequest) (*palacev1.ListTeamStrategyReply, error) {
	return &palacev1.ListTeamStrategyReply{}, nil
}
func (s *TeamStrategyService) SubscribeTeamStrategy(ctx context.Context, req *palacev1.SubscribeTeamStrategyRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamStrategyService) SubscribeTeamStrategies(ctx context.Context, req *palacev1.SubscribeTeamStrategiesRequest) (*palacev1.SubscribeTeamStrategiesReply, error) {
	return &palacev1.SubscribeTeamStrategiesReply{}, nil
}
