package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamDatasourceService struct {
	palacev1.UnimplementedTeamDatasourceServer
	teamDatasourceBiz *biz.TeamDatasource
	helper            *log.Helper
}

func NewTeamDatasourceService(
	teamDatasourceBiz *biz.TeamDatasource,
	logger log.Logger,
) *TeamDatasourceService {
	return &TeamDatasourceService{
		teamDatasourceBiz: teamDatasourceBiz,
		helper:            log.NewHelper(log.With(logger, "module", "service.datasource")),
	}
}

func (s *TeamDatasourceService) SaveTeamMetricDatasource(ctx context.Context, req *palacev1.SaveTeamMetricDatasourceRequest) (*common.EmptyReply, error) {
	params := build.ToSaveTeamMetricDatasourceRequest(req)
	if err := s.teamDatasourceBiz.SaveMetricDatasource(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "保存团队数据源成功"}, nil
}

func (s *TeamDatasourceService) UpdateTeamMetricDatasourceStatus(ctx context.Context, req *palacev1.UpdateTeamMetricDatasourceStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateTeamMetricDatasourceStatusRequest{
		DatasourceID: req.GetDatasourceId(),
		Status:       vobj.GlobalStatus(req.GetStatus()),
	}
	if err := s.teamDatasourceBiz.UpdateMetricDatasourceStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "更新团队数据源状态成功"}, nil
}

func (s *TeamDatasourceService) DeleteTeamMetricDatasource(ctx context.Context, req *palacev1.DeleteTeamMetricDatasourceRequest) (*common.EmptyReply, error) {
	if err := s.teamDatasourceBiz.DeleteMetricDatasource(ctx, req.GetDatasourceId()); err != nil {
		return nil, err
	}
	return &common.EmptyReply{Message: "删除团队数据源成功"}, nil
}

func (s *TeamDatasourceService) GetTeamMetricDatasource(ctx context.Context, req *palacev1.GetTeamMetricDatasourceRequest) (*common.TeamMetricDatasourceItem, error) {
	datasource, err := s.teamDatasourceBiz.GetMetricDatasource(ctx, req.GetDatasourceId())
	if err != nil {
		return nil, err
	}

	return build.ToTeamMetricDatasourceItem(datasource), nil
}

func (s *TeamDatasourceService) ListTeamMetricDatasource(ctx context.Context, req *palacev1.ListTeamMetricDatasourceRequest) (*palacev1.ListTeamMetricDatasourceReply, error) {
	params := build.ToListTeamMetricDatasourceRequest(req)
	datasourceReply, err := s.teamDatasourceBiz.ListMetricDatasource(ctx, params)
	if err != nil {
		return nil, err
	}

	return &palacev1.ListTeamMetricDatasourceReply{
		Pagination: build.ToPaginationReply(datasourceReply.PaginationReply),
		Items:      build.ToTeamMetricDatasourceItems(datasourceReply.Items),
	}, nil
}
