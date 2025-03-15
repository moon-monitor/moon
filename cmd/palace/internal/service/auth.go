package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
)

type AuthService struct {
	palacev1.UnimplementedAuthServer

	authBiz *biz.AuthBiz

	helper *log.Helper
}

func NewAuthService(authBiz *biz.AuthBiz, logger log.Logger) *AuthService {
	return &AuthService{
		authBiz: authBiz,
		helper:  log.NewHelper(log.With(logger, "module", "service.auth")),
	}
}

func (s *AuthService) GetCaptcha(ctx context.Context, _ *palacev1.GetCaptchaRequest) (*palacev1.GetCaptchaReply, error) {
	captchaBo, err := s.authBiz.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetCaptchaReply{
		CaptchaId:      captchaBo.Id,
		CaptchaImg:     captchaBo.B64s,
		ExpiredSeconds: captchaBo.ExpiredSeconds,
	}, nil
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *palacev1.LoginByPasswordRequest) (*palacev1.LoginReply, error) {
	captchaReq := req.GetCaptcha()
	captchaVerify := &bo.CaptchaVerify{
		Id:     captchaReq.GetCaptchaId(),
		Answer: captchaReq.GetAnswer(),
		Clear:  true,
	}

	if err := s.authBiz.VerifyCaptcha(ctx, captchaVerify); err != nil {
		return nil, err
	}
	loginReq := &bo.LoginByPassword{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	loginSign, err := s.authBiz.LoginByPassword(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	return loginSign.LoginReply(), nil
}

func (s *AuthService) Logout(ctx context.Context, req *palacev1.LogoutRequest) (*palacev1.LogoutReply, error) {
	token, ok := permission.GetTokenByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("token error")
	}
	if err := s.authBiz.Logout(ctx, token); err != nil {
		return nil, err
	}
	return &palacev1.LogoutReply{Redirect: req.GetRedirect()}, nil
}
