package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
)

type MetricInit interface {
	Init(config bo.MetricDatasourceConfig) (Metric, error)
}

type Metric interface {
	Query(ctx context.Context, expr string, duration time.Duration) ([]*bo.MetricQueryReply, error)

	QueryRange(ctx context.Context, expr string, start, end int64, step uint32) ([]*bo.MetricQueryRangeReply, error)

	Metadata(ctx context.Context) (<-chan []*bo.MetricItem, error)
}
