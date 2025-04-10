package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
)

type MetricInit interface {
	Init(config bo.MetricDatasourceConfig) (Metric, error)
}

type Metric interface {
	Query(ctx context.Context, expr string, duration time.Duration) ([]*do.MetricQueryReply, error)

	QueryRange(ctx context.Context, expr string, start, end time.Time) ([]*do.MetricQueryRangeReply, error)

	Metadata(ctx context.Context) (<-chan []*do.MetricItem, error)
}
