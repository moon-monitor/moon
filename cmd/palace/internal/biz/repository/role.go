package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

type TeamRole interface {
	Get(ctx context.Context, id uint32) (do.TeamRole, error)
	List(ctx context.Context, req *bo.ListRoleReq) (*bo.ListTeamRoleReply, error)
	Create(ctx context.Context, role bo.Role) error
	Update(ctx context.Context, role bo.Role) error
	Delete(ctx context.Context, id uint32) error
	UpdateStatus(ctx context.Context, req *bo.UpdateRoleStatusReq) error
}

type Role interface {
	Get(ctx context.Context, id uint32) (do.Role, error)
	List(ctx context.Context, req *bo.ListRoleReq) (*bo.ListRoleReply, error)
	Create(ctx context.Context, role bo.Role) error
	Update(ctx context.Context, role bo.Role) error
	Delete(ctx context.Context, id uint32) error
	UpdateStatus(ctx context.Context, req *bo.UpdateRoleStatusReq) error
}
