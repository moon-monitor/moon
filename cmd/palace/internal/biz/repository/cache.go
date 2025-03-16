package repository

import (
	"context"
	"time"

	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

type Cache interface {
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) error
	BanToken(ctx context.Context, token string) error
	VerifyToken(ctx context.Context, token string) error
	VerifyOAuthToken(ctx context.Context, oauthID uint32, token string) error
	WaitVerifyOAuthToken(ctx context.Context, oauthID uint32, token string) error
}

const (
	BankTokenKey  cache.K = "palace:token:ban"
	OAuthTokenKey cache.K = "palace:token:oauth"
)
