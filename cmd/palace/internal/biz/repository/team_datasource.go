package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

type TeamDatasourceMetric interface {
	Create(ctx context.Context, req *bo.SaveTeamMetricDatasource) error
	Update(ctx context.Context, req *bo.SaveTeamMetricDatasource) error
	UpdateStatus(ctx context.Context, req *bo.UpdateTeamMetricDatasourceStatusRequest) error
	Delete(ctx context.Context, datasourceID uint32) error
	Get(ctx context.Context, datasourceID uint32) (do.DatasourceMetric, error)
	List(ctx context.Context, req *bo.ListTeamMetricDatasource) (*bo.ListTeamMetricDatasourceReply, error)
}
