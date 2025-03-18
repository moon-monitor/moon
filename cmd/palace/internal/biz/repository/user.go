package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type User interface {
	FindByID(ctx context.Context, userID uint32) (*system.User, error)
	FindByEmail(ctx context.Context, email crypto.String) (*system.User, error)
	CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser) (*system.User, error)
	SetEmail(ctx context.Context, user *system.User) (*system.User, error)
	Create(ctx context.Context, user *system.User) (*system.User, error)
}
