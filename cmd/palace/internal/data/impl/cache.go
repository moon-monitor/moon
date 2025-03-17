package impl

import (
	"context"
	_ "embed"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/password"
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

func (c *cacheReoImpl) VerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error {
	exist, err := c.GetCache().Client().Exists(ctx, repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token)).Result()
	if err != nil {
		return merr.ErrorInternalServerError("cache err").WithCause(err)
	}
	if exist == 0 {
		return merr.ErrorUnauthorized("oauth unauthorized").WithMetadata(map[string]string{
			"exist": "false",
		})
	}
	if validate.TextIsNull(oauthParams.Code) {
		return nil
	}
	code, err := c.GetCache().Client().Get(ctx, repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token)).Result()
	if err != nil {
		return merr.ErrorInternalServerError("cache err").WithCause(err)
	}
	if code != oauthParams.Code {
		return merr.ErrorUnauthorized("oauth unauthorized").WithMetadata(map[string]string{
			"code": "err",
		})
	}
	return nil
}

func (c *cacheReoImpl) CacheVerifyOAuthToken(ctx context.Context, oauthParams *bo.OAuthLoginParams) error {
	return c.GetCache().Client().Set(ctx, repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token), "##code##", 10*time.Minute).Err()
}

//go:embed template/verify_email.html
var verifyEmailTemplate string

func (c *cacheReoImpl) SendVerifyEmailCode(ctx context.Context, oauthParams *bo.OAuthLoginParams) error {
	if err := validate.CheckEmail(oauthParams.Email); err != nil {
		return err
	}
	err := c.GetCache().Client().Set(ctx, repository.OAuthTokenKey.Key(oauthParams.OAuthID, oauthParams.Token), oauthParams.Code, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	bodyParams := map[string]string{
		"Email":       oauthParams.Email,
		"Code":        oauthParams.Code,
		"RedirectURI": c.bc.GetAuth().GetOauth2().GetRedirectUri(),
	}
	emailBody, err := template.HtmlFormatter(verifyEmailTemplate, bodyParams)
	if err != nil {
		return err
	}
	// 发送用户密码到用户邮箱
	return c.GetEmail().SetSubject("Email verification code.").SetTo(oauthParams.Email).SetBody(emailBody, "text/html").Send()
}
