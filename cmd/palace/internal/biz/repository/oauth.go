package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type OAuth interface {
	Create(ctx context.Context, user bo.IOAuthUser) (*system.OAuthUser, error)
	FindByOAuthID(ctx context.Context, oauthID uint32, app vobj.OAuthAPP) (*system.OAuthUser, error)
	SetUser(ctx context.Context, user *system.OAuthUser) (*system.OAuthUser, error)
}
