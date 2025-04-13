package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamDashboardService struct {
	palacev1.UnimplementedTeamDashboardServer
}

func NewTeamDashboardService() *TeamDashboardService {
	return &TeamDashboardService{}
}

func (s *TeamDashboardService) SaveTeamDashboard(ctx context.Context, req *palacev1.SaveTeamDashboardRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDashboardService) DeleteTeamDashboard(ctx context.Context, req *palacev1.DeleteTeamDashboardRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDashboardService) GetTeamDashboard(ctx context.Context, req *palacev1.GetTeamDashboardRequest) (*palacev1.GetTeamDashboardReply, error) {
	return &palacev1.GetTeamDashboardReply{}, nil
}
func (s *TeamDashboardService) ListTeamDashboard(ctx context.Context, req *palacev1.ListTeamDashboardRequest) (*palacev1.ListTeamDashboardReply, error) {
	return &palacev1.ListTeamDashboardReply{}, nil
}
func (s *TeamDashboardService) UpdateTeamDashboardStatus(ctx context.Context, req *palacev1.UpdateTeamDashboardStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDashboardService) SaveTeamDashboardChart(ctx context.Context, req *palacev1.SaveTeamDashboardChartRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDashboardService) DeleteTeamDashboardChart(ctx context.Context, req *palacev1.DeleteTeamDashboardChartRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDashboardService) GetTeamDashboardChart(ctx context.Context, req *palacev1.GetTeamDashboardChartRequest) (*palacev1.GetTeamDashboardChartReply, error) {
	return &palacev1.GetTeamDashboardChartReply{}, nil
}
func (s *TeamDashboardService) ListTeamDashboardChart(ctx context.Context, req *palacev1.ListTeamDashboardChartRequest) (*palacev1.ListTeamDashboardChartReply, error) {
	return &palacev1.ListTeamDashboardChartReply{}, nil
}
func (s *TeamDashboardService) UpdateTeamDashboardChartStatus(ctx context.Context, req *palacev1.UpdateTeamDashboardChartStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
