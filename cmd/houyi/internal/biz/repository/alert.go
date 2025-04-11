package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
)

type Alert interface {
	Save(ctx context.Context, alerts ...bo.Alert) error
	Get(ctx context.Context, fingerprint string) (bo.Alert, bool)
	Delete(ctx context.Context, fingerprint string) error
}
