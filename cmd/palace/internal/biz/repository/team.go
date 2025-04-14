package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
)

type Team interface {
	FindByID(ctx context.Context, id uint32) (*system.Team, error)
	Create(ctx context.Context, team bo.Team) error
	Update(ctx context.Context, team bo.Team) error
	Delete(ctx context.Context, id uint32) error
	List(ctx context.Context, req *bo.TeamListRequest) (*bo.TeamListReply, error)
}
