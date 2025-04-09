package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
)

type Config interface {
	GetMetricDatasourceConfig(ctx context.Context, id string) (bo.MetricDatasourceConfig, bool)
	SetMetricDatasourceConfig(ctx context.Context, configs ...bo.MetricDatasourceConfig) error
}
