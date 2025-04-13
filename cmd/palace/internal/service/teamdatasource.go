package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamDatasourceService struct {
	palacev1.UnimplementedTeamDatasourceServer
}

func NewTeamDatasourceService() *TeamDatasourceService {
	return &TeamDatasourceService{}
}

func (s *TeamDatasourceService) SaveTeamMetricDatasource(ctx context.Context, req *palacev1.SaveTeamMetricDatasourceRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDatasourceService) UpdateTeamMetricDatasourceStatus(ctx context.Context, req *palacev1.UpdateTeamMetricDatasourceStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDatasourceService) DeleteTeamMetricDatasource(ctx context.Context, req *palacev1.DeleteTeamMetricDatasourceRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDatasourceService) GetTeamMetricDatasource(ctx context.Context, req *palacev1.GetTeamMetricDatasourceRequest) (*palacev1.GetTeamMetricDatasourceReply, error) {
	return &palacev1.GetTeamMetricDatasourceReply{}, nil
}
func (s *TeamDatasourceService) ListTeamMetricDatasource(ctx context.Context, req *palacev1.ListTeamMetricDatasourceRequest) (*palacev1.ListTeamMetricDatasourceReply, error) {
	return &palacev1.ListTeamMetricDatasourceReply{}, nil
}
