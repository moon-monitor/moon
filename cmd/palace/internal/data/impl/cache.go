package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/password"
)

func NewCacheRepo(bc *conf.Bootstrap, d *data.Data, logger log.Logger) repository.Cache {
	return &cacheReoImpl{
		bc:      bc,
		signKey: bc.GetAuth().GetJwt().GetSignKey(),
		Data:    d,
		helper:  log.NewHelper(log.With(logger, "module", "data.repo.cache")),
	}
}

type cacheReoImpl struct {
	bc      *conf.Bootstrap
	signKey string
	*data.Data

	helper *log.Helper
}

func (c *cacheReoImpl) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.GetCache().Client().SetNX(ctx, key, 1, expiration).Result()
}

func (c *cacheReoImpl) Unlock(ctx context.Context, key string) error {
	return c.GetCache().Client().Del(ctx, key).Err()
}

func (c *cacheReoImpl) BanToken(ctx context.Context, token string) error {
	jwtClaims, err := middleware.ParseJwtClaimsFromToken(token, c.signKey)
	if err != nil {
		return err
	}
	expiration := jwtClaims.ExpiresAt.Sub(time.Now())
	if expiration <= 0 {
		return merr.ErrorInvalidToken("token is invalid")
	}
	return c.GetCache().Client().Set(ctx, repository.BankTokenKey.Key(password.MD5(token)), 1, expiration).Err()
}

func (c *cacheReoImpl) VerifyToken(ctx context.Context, token string) error {
	exist, err := c.GetCache().Client().Exists(ctx, repository.BankTokenKey.Key(password.MD5(token))).Result()
	if err != nil {
		return err
	}
	if exist > 0 {
		return merr.ErrorInvalidToken("token is ban")
	}
	return nil
}
