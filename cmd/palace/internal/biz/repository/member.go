package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

type Member interface {
	FindByUserID(ctx context.Context, userID uint32) (do.TeamMember, error)
}
