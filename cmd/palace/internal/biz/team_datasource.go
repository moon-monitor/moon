package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewTeamDatasource(
	teamDatasourceMetricRepo repository.TeamDatasourceMetric,
) *TeamDatasource {
	return &TeamDatasource{
		teamDatasourceMetricRepo: teamDatasourceMetricRepo,
	}
}

type TeamDatasource struct {
	teamDatasourceMetricRepo repository.TeamDatasourceMetric
}

func (t *TeamDatasource) SaveMetricDatasource(ctx context.Context, req *bo.SaveTeamMetricDatasource) error {
	if req.ID <= 0 {
		return t.teamDatasourceMetricRepo.Create(ctx, req)
	}
	metricDatasourceDo, err := t.teamDatasourceMetricRepo.Get(ctx, req.ID)
	if err != nil {
		return err
	}
	if metricDatasourceDo.GetStatus().IsEnable() {
		return merr.ErrorBadRequest("数据源已启用，不能修改")
	}

	return t.teamDatasourceMetricRepo.Update(ctx, req)
}

func (t *TeamDatasource) UpdateMetricDatasourceStatus(ctx context.Context, req *bo.UpdateTeamMetricDatasourceStatusRequest) error {
	return t.teamDatasourceMetricRepo.UpdateStatus(ctx, req)
}

func (t *TeamDatasource) DeleteMetricDatasource(ctx context.Context, datasourceID uint32) error {
	return t.teamDatasourceMetricRepo.Delete(ctx, datasourceID)
}

func (t *TeamDatasource) GetMetricDatasource(ctx context.Context, datasourceID uint32) (do.DatasourceMetric, error) {
	return t.teamDatasourceMetricRepo.Get(ctx, datasourceID)
}

func (t *TeamDatasource) ListMetricDatasource(ctx context.Context, req *bo.ListTeamMetricDatasource) (*bo.ListTeamMetricDatasourceReply, error) {
	return t.teamDatasourceMetricRepo.List(ctx, req)
}
