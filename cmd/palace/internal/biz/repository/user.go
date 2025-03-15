package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
)

type User interface {
	FindByID(ctx context.Context, userID uint32) (*system.User, error)
	FindByEmail(ctx context.Context, email string) (*system.User, error)
}
