package biz

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewTeamDatasource(
	teamDatasourceMetricRepo repository.TeamDatasourceMetric,
) *TeamDatasource {
	return &TeamDatasource{
		TeamDatasourceMetricRepo: teamDatasourceMetricRepo,
	}
}

type TeamDatasource struct {
	TeamDatasourceMetricRepo repository.TeamDatasourceMetric
}

func (t *TeamDatasource) SaveMetricDatasource(ctx context.Context, req *bo.SaveTeamMetricDatasource) error {
	return t.TeamDatasourceMetricRepo.Create(ctx, req)
}

func (t *TeamDatasource) UpdateMetricDatasourceStatus(ctx context.Context, req *bo.UpdateTeamMetricDatasourceStatusRequest) error {
	return t.TeamDatasourceMetricRepo.UpdateStatus(ctx, req)
}

func (t *TeamDatasource) DeleteMetricDatasource(ctx context.Context, datasourceID uint32) error {
	return t.TeamDatasourceMetricRepo.Delete(ctx, datasourceID)
}

func (t *TeamDatasource) GetMetricDatasource(ctx context.Context, datasourceID uint32) (do.DatasourceMetric, error) {
	return t.TeamDatasourceMetricRepo.Get(ctx, datasourceID)
}

func (t *TeamDatasource) ListMetricDatasource(ctx context.Context, req *bo.ListTeamMetricDatasource) (*bo.ListTeamMetricDatasourceReply, error) {
	return t.TeamDatasourceMetricRepo.List(ctx, req)
}
