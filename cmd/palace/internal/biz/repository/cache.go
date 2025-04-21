package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

type Cache interface {
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) error
	BanToken(ctx context.Context, token string) error
	VerifyToken(ctx context.Context, token string) error

	VerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error
	CacheVerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error
	SendVerifyEmailCode(ctx context.Context, email string) (*bo.SendEmailParams, error)
	VerifyEmailCode(ctx context.Context, email, code string) error

	CacheUsers(ctx context.Context, users ...do.User) error
	GetUser(ctx context.Context, userID uint32) (do.User, error)
	GetUsers(ctx context.Context, ids ...uint32) ([]do.User, error)
}

const (
	EmailCodeKey  cache.K = "palace:verify:email:code"
	BankTokenKey  cache.K = "palace:token:ban"
	OAuthTokenKey cache.K = "palace:token:oauth"
	UserCacheKey  cache.K = "palace:user:cache"
)
