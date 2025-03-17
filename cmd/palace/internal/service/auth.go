package service

import (
	"context"
	nhttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/oauth2"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/merr"
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

func (s *AuthService) VerifyEmail(ctx context.Context, req *palacev1.VerifyEmailRequest) (*palacev1.VerifyEmailReply, error) {
	captchaReq := req.GetCaptcha()
	captchaVerify := &bo.CaptchaVerify{
		Id:     captchaReq.GetCaptchaId(),
		Answer: captchaReq.GetAnswer(),
		Clear:  true,
	}

	if err := s.authBiz.VerifyCaptcha(ctx, captchaVerify); err != nil {
		return nil, err
	}
	oauthParams := &bo.OAuthLoginParams{
		APP:     vobj.OAuthAPP(req.GetApp()),
		Code:    "",
		Email:   req.GetEmail(),
		OAuthID: req.GetOauthID(),
		Token:   req.GetToken(),
	}
	if err := s.authBiz.VerifyOAuthLoginEmail(ctx, oauthParams); err != nil {
		return nil, err
	}
	return &palacev1.VerifyEmailReply{ExpiredSeconds: int64(5 * time.Minute.Seconds())}, nil
}

func (s *AuthService) LoginByEmail(ctx context.Context, req *palacev1.LoginByEmailRequest) (*palacev1.LoginReply, error) {
	oauthParams := &bo.OAuthLoginParams{
		APP:     vobj.OAuthAPP(req.GetApp()),
		Code:    req.GetCode(),
		Email:   req.GetEmail(),
		OAuthID: req.GetOauthID(),
		Token:   req.GetToken(),
	}
	loginSign, err := s.authBiz.OAuthLoginWithEmail(ctx, oauthParams)
	if err != nil {
		return nil, err
	}
	return loginSign.LoginReply(), nil
}

func (s *AuthService) VerifyToken(ctx context.Context, token string) error {
	return s.authBiz.VerifyToken(ctx, token)
}

func (s *AuthService) RefreshToken(ctx context.Context, req *palacev1.RefreshTokenRequest) (*palacev1.LoginReply, error) {
	token, ok := permission.GetTokenByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("token error")
	}
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("token error")
	}
	refreshReq := &bo.RefreshToken{
		Token:  token,
		UserID: userID,
	}
	loginSign, err := s.authBiz.RefreshToken(ctx, refreshReq)
	if err != nil {
		return nil, err
	}
	return loginSign.LoginReply(), nil
}

// OAuthLogin oauth login
func (s *AuthService) OAuthLogin(app vobj.OAuthAPP) http.HandlerFunc {
	return func(ctx http.Context) error {
		oauthConf, err := s.authBiz.GetOAuthConf(app)
		if err != nil {
			return err
		}
		// 重定向到指定地址
		url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOnline)
		req := ctx.Request()
		resp := ctx.Response()
		resp.Header().Set("Location", url)
		resp.WriteHeader(nhttp.StatusTemporaryRedirect)
		ctx.Reset(resp, req)
		return nil
	}
}

// OAuthLoginCallback oauth callback
func (s *AuthService) OAuthLoginCallback(app vobj.OAuthAPP) http.HandlerFunc {
	return func(ctx http.Context) error {
		code := ctx.Query().Get("code")
		loginRedirect, err := s.authBiz.OAuthLogin(ctx, app, code)
		if err != nil {
			return err
		}
		// 重定向到指定地址
		req := ctx.Request()
		resp := ctx.Response()

		resp.Header().Set("Location", loginRedirect)
		resp.WriteHeader(nhttp.StatusTemporaryRedirect)
		ctx.Reset(resp, req)
		return nil
	}
}
