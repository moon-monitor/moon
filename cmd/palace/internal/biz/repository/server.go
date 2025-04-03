package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
)

type Server interface {
	RegisterRabbit(ctx context.Context, req *bo.ServerRegisterReq) error
	RegisterHouyi(ctx context.Context, req *bo.ServerRegisterReq) error
}
