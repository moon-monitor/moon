package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
)

type Cache interface {
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) error

	StorageMetric(ctx context.Context, metrics ...bo.MetricVec) error
}
