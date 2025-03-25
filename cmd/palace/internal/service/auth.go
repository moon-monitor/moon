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
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type AuthService struct {
	palacev1.UnimplementedAuthServer
	authBiz       *biz.AuthBiz
	permissionBiz *biz.PermissionBiz
	oauth2List    []*palacev1.OAuth2ListReply_OAuthItem
	helper        *log.Helper
}

func builderOAuth2List(oauth2 *conf.Auth_OAuth2) []*palacev1.OAuth2ListReply_OAuthItem {
	if !oauth2.GetEnable() {
		return nil
	}
	list := oauth2.GetConfigs()
	oauthList := make([]*palacev1.OAuth2ListReply_OAuthItem, 0, len(list))
	for _, oauth := range list {
		oauthList = append(oauthList, &palacev1.OAuth2ListReply_OAuthItem{
			Icon:     oauth.GetApp().String(),
			Label:    oauth.GetApp().String() + " login",
			Redirect: oauth.GetAuthUrl(),
		})
	}
	return oauthList
}

func login(loginSign *bo.LoginSign, err error) (*palacev1.LoginReply, error) {
	if err != nil {
		return nil, err
	}
	return loginSign.LoginReply(), nil
}

func NewAuthService(
	bc *conf.Bootstrap,
	authBiz *biz.AuthBiz,
	permissionBiz *biz.PermissionBiz,
	logger log.Logger,
) *AuthService {
	return &AuthService{
		authBiz:       authBiz,
		permissionBiz: permissionBiz,
		oauth2List:    builderOAuth2List(bc.GetAuth().GetOauth2()),
		helper:        log.NewHelper(log.With(logger, "module", "service.auth")),
	}
}

func (s *AuthService) GetCaptcha(ctx context.Context, _ *common.EmptyRequest) (*palacev1.GetCaptchaReply, error) {
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
	return login(s.authBiz.LoginByPassword(ctx, loginReq))
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
	if err := s.authBiz.VerifyEmail(ctx, req.GetEmail()); err != nil {
		return nil, err
	}
	return &palacev1.VerifyEmailReply{ExpiredSeconds: int64(5 * time.Minute.Seconds())}, nil
}

func (s *AuthService) LoginByEmail(ctx context.Context, req *palacev1.LoginByEmailRequest) (*palacev1.LoginReply, error) {
	userDo := &system.User{
		BaseModel: do.BaseModel{},
		Username:  req.GetUsername(),
		Nickname:  req.GetNickname(),
		Email:     crypto.String(req.GetEmail()),
		Remark:    req.GetRemark(),
		Gender:    vobj.Gender(req.GetGender()),
		Position:  vobj.RoleUser,
		Status:    vobj.UserStatusNormal,
	}
	return login(s.authBiz.LoginWithEmail(ctx, req.GetCode(), userDo))
}

func (s *AuthService) OAuthLoginByEmail(ctx context.Context, req *palacev1.OAuthLoginByEmailRequest) (*palacev1.LoginReply, error) {
	oauthParams := &bo.OAuthLoginParams{
		APP:     vobj.OAuthAPP(req.GetApp()),
		Code:    req.GetCode(),
		Email:   req.GetEmail(),
		OAuthID: req.GetOauthID(),
		Token:   req.GetToken(),
	}
	return login(s.authBiz.OAuthLoginWithEmail(ctx, oauthParams))
}

func (s *AuthService) VerifyToken(ctx context.Context, token string) error {
	return s.authBiz.VerifyToken(ctx, token)
}

func (s *AuthService) VerifyPermission(ctx context.Context) error {
	return s.permissionBiz.VerifyPermission(ctx)
}

func (s *AuthService) RefreshToken(ctx context.Context, _ *common.EmptyRequest) (*palacev1.LoginReply, error) {
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
	return login(s.authBiz.RefreshToken(ctx, refreshReq))
}

func (s *AuthService) OAuth2List(_ context.Context, _ *common.EmptyRequest) (*palacev1.OAuth2ListReply, error) {
	return &palacev1.OAuth2ListReply{Items: s.oauth2List}, nil
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
