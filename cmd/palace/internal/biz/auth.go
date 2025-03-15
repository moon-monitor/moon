package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewAuthBiz(userRepo repository.User, captcha repository.Captcha, logger log.Logger) *AuthBiz {
	return &AuthBiz{
		userRepo: userRepo,
		captcha:  captcha,
		helper:   log.NewHelper(log.With(logger, "module", "biz.auth")),
	}
}

type AuthBiz struct {
	userRepo repository.User
	captcha  repository.Captcha

	helper *log.Helper
}

// GetCaptcha get image captcha
func (a *AuthBiz) GetCaptcha(ctx context.Context) (*bo.Captcha, error) {
	return a.captcha.Generate(ctx)
}

// VerifyCaptcha Captcha
func (a *AuthBiz) VerifyCaptcha(ctx context.Context, req *bo.CaptchaVerify) error {
	verify := a.captcha.Verify(ctx, req)
	if !verify {
		return merr.ErrorCaptchaError("captcha err")
	}
	return nil
}
