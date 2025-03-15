package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewAuthBiz(bc *conf.Bootstrap, userRepo repository.User, captcha repository.Captcha, logger log.Logger) *AuthBiz {
	return &AuthBiz{
		bc:       bc,
		userRepo: userRepo,
		captcha:  captcha,
		helper:   log.NewHelper(log.With(logger, "module", "biz.auth")),
	}
}

type AuthBiz struct {
	bc       *conf.Bootstrap
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

// LoginByPassword login by password
func (a *AuthBiz) LoginByPassword(ctx context.Context, req *bo.LoginByPassword) (*bo.LoginSign, error) {
	user, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, merr.ErrorUserNotFound("user not found")
	}
	if !user.ValidatePassword(req.Password) {
		return nil, merr.ErrorPasswordError("password error")
	}
	return a.login(user)
}

func (a *AuthBiz) login(userDo *system.User) (*bo.LoginSign, error) {
	base := &middleware.JwtBaseInfo{
		UserID:   userDo.ID,
		Username: userDo.Username,
		Nickname: userDo.Nickname,
		Avatar:   userDo.Avatar,
		Gender:   userDo.Gender,
	}
	token, err := middleware.NewJwtClaims(a.bc.GetAuth().GetJwt(), base).GetToken()
	if err != nil {
		return nil, err
	}
	return &bo.LoginSign{
		Base:           base,
		ExpiredSeconds: int64(a.bc.GetAuth().GetJwt().GetExpire().AsDuration().Seconds()),
		Token:          token,
	}, nil
}
