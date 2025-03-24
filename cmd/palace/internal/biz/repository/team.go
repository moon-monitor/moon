package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
)

type Team interface {
	FindByID(ctx context.Context, id uint32) (*system.Team, error)
}
