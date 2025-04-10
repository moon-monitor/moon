package impl

import (
	"context"
	_ "embed"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/hash"
	"github.com/moon-monitor/moon/pkg/util/template"
	"github.com/moon-monitor/moon/pkg/util/validate"
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
	return c.GetCache().Client().Set(ctx, repository.BankTokenKey.Key(hash.MD5(token)), 1, expiration).Err()
}

func (c *cacheReoImpl) VerifyToken(ctx context.Context, token string) error {
	exist, err := c.GetCache().Client().Exists(ctx, repository.BankTokenKey.Key(hash.MD5(token))).Result()
	if err != nil {
		return err
	}
	if exist > 0 {
		return merr.ErrorInvalidToken("token is ban")
	}
	return nil
}

func (c *cacheReoImpl) VerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error {
	key := repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token)
	exist, err := c.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		return merr.ErrorInternalServerError("cache err").WithCause(err)
	}
	if exist == 0 {
		return merr.ErrorUnauthorized("oauth unauthorized").WithMetadata(map[string]string{
			"exist": "false",
		})
	}
	return c.GetCache().Client().Del(ctx, key).Err()
}

func (c *cacheReoImpl) CacheVerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error {
	return c.GetCache().Client().Set(ctx, repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token), "##code##", 10*time.Minute).Err()
}

func (c *cacheReoImpl) VerifyEmailCode(ctx context.Context, email, code string) error {
	key := repository.EmailCodeKey.Key(email)
	cacheCode, err := c.GetCache().Client().Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return merr.ErrorCaptchaError("captcha is expire").WithMetadata(map[string]string{
				"code": "captcha is expire",
			})
		}
		return merr.ErrorInternalServerError("cache err").WithCause(err)
	}
	defer c.GetCache().Client().Del(ctx, key).Val()
	if strings.EqualFold(cacheCode, code) {
		return nil
	}
	return merr.ErrorCaptchaError("captcha err").WithMetadata(map[string]string{
		"code": "The verification code is incorrect. Please retrieve a new one and try again.",
	})
}

//go:embed template/verify_email.html
var verifyEmailTemplate string

func (c *cacheReoImpl) SendVerifyEmailCode(ctx context.Context, email string) (*bo.SendEmailParams, error) {
	if err := validate.CheckEmail(email); err != nil {
		return nil, err
	}
	code := strings.ToUpper(hash.MD5(time.Now().String())[:6])
	err := c.GetCache().Client().Set(ctx, repository.EmailCodeKey.Key(email), code, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	bodyParams := map[string]string{
		"Email":       email,
		"Code":        code,
		"RedirectURI": c.bc.GetAuth().GetOauth2().GetRedirectUri(),
	}
	emailBody, err := template.HtmlFormatter(verifyEmailTemplate, bodyParams)
	if err != nil {
		return nil, err
	}
	params := &bo.SendEmailParams{
		Email:       email,
		Body:        emailBody,
		Subject:     "Email verification code.",
		ContentType: "text/html",
	}
	return params, nil
}
