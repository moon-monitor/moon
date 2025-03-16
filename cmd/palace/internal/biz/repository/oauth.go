package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type OAuth interface {
	OAuthUserFirstOrCreate(ctx context.Context, oauthUser bo.IOAuthUser) (*system.User, error)
	SetEmail(ctx context.Context, userID uint32, email string) (*system.User, error)
	GetSysUserByOAuthID(ctx context.Context, oauthID uint32, app vobj.OAuthAPP) (*system.OAuthUser, error)
	SendVerifyEmail(ctx context.Context, email string) error
	CheckVerifyEmailCode(ctx context.Context, email, code string) error
}
