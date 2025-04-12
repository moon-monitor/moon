package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
)

type Member interface {
	FindByUserID(ctx context.Context, userID uint32) (*team.Member, error)
}
