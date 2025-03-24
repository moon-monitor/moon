package biz

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/oauth2"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/password"
	"github.com/moon-monitor/moon/pkg/util/safety"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func buildOAuthConf(c *conf.Auth_OAuth2) *safety.Map[vobj.OAuthAPP, *oauth2.Config] {
	oauth2Map := safety.NewMap[vobj.OAuthAPP, *oauth2.Config]()
	for _, config := range c.GetConfigs() {
		oauth2Map.Set(vobj.OAuthAPP(config.GetApp()), &oauth2.Config{
			ClientID:     config.GetClientId(),
			ClientSecret: config.GetClientSecret(),
			Endpoint: oauth2.Endpoint{
				AuthURL:  config.GetAuthUrl(),
				TokenURL: config.GetTokenUrl(),
			},
			RedirectURL: config.GetCallbackUri(),
			Scopes:      config.GetScopes(),
		})
	}
	return oauth2Map
}

func NewAuthBiz(
	bc *conf.Bootstrap,
	userRepo repository.User,
	memberRepo repository.Member,
	captchaRepo repository.Captcha,
	cacheRepo repository.Cache,
	oauthRepo repository.OAuth,
	resourceRepo repository.Resource,
	teamRepo repository.Team,
	transaction repository.Transaction,
	logger log.Logger,
) *AuthBiz {
	return &AuthBiz{
		bc:           bc,
		redirectURL:  bc.GetAuth().GetOauth2().GetRedirectUri(),
		oauthConfigs: buildOAuthConf(bc.GetAuth().GetOauth2()),
		userRepo:     userRepo,
		memberRepo:   memberRepo,
		captchaRepo:  captchaRepo,
		cacheRepo:    cacheRepo,
		oauthRepo:    oauthRepo,
		resourceRepo: resourceRepo,
		teamRepo:     teamRepo,
		transaction:  transaction,
		helper:       log.NewHelper(log.With(logger, "module", "biz.auth")),
	}
}

type AuthBiz struct {
	bc           *conf.Bootstrap
	redirectURL  string
	oauthConfigs *safety.Map[vobj.OAuthAPP, *oauth2.Config]

	userRepo     repository.User
	memberRepo   repository.Member
	captchaRepo  repository.Captcha
	cacheRepo    repository.Cache
	oauthRepo    repository.OAuth
	resourceRepo repository.Resource
	teamRepo     repository.Team
	transaction  repository.Transaction

	helper *log.Helper
}

// GetCaptcha get image captchaRepo
func (a *AuthBiz) GetCaptcha(ctx context.Context) (*bo.Captcha, error) {
	return a.captchaRepo.Generate(ctx)
}

// VerifyCaptcha Captcha
func (a *AuthBiz) VerifyCaptcha(ctx context.Context, req *bo.CaptchaVerify) error {
	verify := a.captchaRepo.Verify(ctx, req)
	if !verify {
		return merr.ErrorCaptchaError("captcha err").WithMetadata(map[string]string{
			"captcha.answer": "The verification code is incorrect. Please retrieve a new one and try again.",
		})
	}
	return nil
}

// Logout token logout
func (a *AuthBiz) Logout(ctx context.Context, token string) error {
	return a.cacheRepo.BanToken(ctx, token)
}

// VerifyToken verify token
func (a *AuthBiz) VerifyToken(ctx context.Context, token string) error {
	if err := a.cacheRepo.VerifyToken(ctx, token); err != nil {
		return err
	}
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorInvalidToken("token is invalid")
	}
	userDo, err := a.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	if !userDo.Status.IsNormal() {
		return merr.ErrorUserForbidden("user forbidden")
	}
	return nil
}

// VerifyPermission verify permission
func (a *AuthBiz) VerifyPermission(ctx context.Context) error {
	operation, ok := permission.GetOperationByContext(ctx)
	if !ok {
		return merr.ErrorBadRequest("operation is invalid")
	}
	resourceDo, err := a.resourceRepo.GetResourceByOperation(ctx, operation)
	if err != nil {
		return err
	}
	if !resourceDo.Status.IsEnabled() {
		return merr.ErrorPermissionDenied("permission denied")
	}

	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorBadRequest("user id is invalid")
	}
	userDo, err := a.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	if !userDo.Status.IsNormal() {
		return merr.ErrorUserForbidden("user forbidden")
	}

	if resourceDo.Allow.IsNone() || resourceDo.Allow.IsUser() {
		return nil
	}

	systemPosition, err := a.verifyPermissionWithSystemPosition(ctx, userDo)
	if err != nil {
		return err
	}
	if systemPosition.IsAdminOrSuperAdmin() {
		return nil
	}
	if err := a.verifyPermissionWithSystemRBAC(ctx, userDo, resourceDo); err != nil {
		return err
	}
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return merr.ErrorPermissionDenied("team id is invalid")
	}
	teamDo, err := a.teamRepo.FindByID(ctx, teamID)
	if err != nil {
		return err
	}
	if !teamDo.Status.IsNormal() {
		return merr.ErrorPermissionDenied("team is invalid")
	}
	teamPosition, memberDo, err := a.verifyPermissionWithTeamPosition(ctx, userDo)
	if err != nil {
		return err
	}
	if teamDo.ID != memberDo.TeamID {
		return merr.ErrorPermissionDenied("team id is invalid")
	}
	if teamPosition.IsAdminOrSuperAdmin() {
		return nil
	}

	if err := a.verifyPermissionWithTeamRBAC(ctx, memberDo, resourceDo); err != nil {
		return err
	}

	return nil
}

func (a *AuthBiz) verifyPermissionWithSystemPosition(ctx context.Context, userDo *system.User) (vobj.Role, error) {
	sysPosition, ok := permission.GetSysPositionByContext(ctx)
	if !ok {
		return userDo.Position, nil
	}
	if userDo.Position.GTE(sysPosition) {
		return sysPosition, nil
	}
	return 0, merr.ErrorPermissionDenied("Your current role [%s] is not allowed to access this resource", sysPosition)
}

func (a *AuthBiz) verifyPermissionWithSystemRBAC(ctx context.Context, userDo *system.User, resourceDo *system.Resource) error {
	if !resourceDo.Allow.IsSystemRBAC() {
		return nil
	}
	sysRoleID, ok := permission.GetSysRoleIDByContext(ctx)
	if ok {
		// 判断角色是否存在，且该角色具备次API权限
		systemRoleDo, ok := validate.SliceFindByValue(userDo.Roles, sysRoleID, func(role *system.SysRole) uint32 {
			return role.ID
		})
		if !ok {
			return merr.ErrorPermissionDenied("User role is invalid.")
		}
		if !systemRoleDo.Status.IsNormal() {
			return merr.ErrorPermissionDenied("role is invalid [%s]", systemRoleDo.Status)
		}
		_, ok = validate.SliceFindByValue(systemRoleDo.Resources, resourceDo.ID, func(role *system.Resource) uint32 {
			return role.ID
		})
		if ok {
			return nil
		}
		return merr.ErrorPermissionDenied("User role resource is invalid.")
	}
	resources := make([]*system.Resource, 0, len(userDo.Roles)*10)
	for _, role := range userDo.Roles {
		if role.Status.IsNormal() {
			resources = append(resources, role.Resources...)
		}
	}
	_, ok = validate.SliceFindByValue(resources, resourceDo.ID, func(role *system.Resource) uint32 {
		return role.ID
	})
	if ok {
		return nil
	}
	return merr.ErrorPermissionDenied("User role resource is invalid.")
}

func (a *AuthBiz) verifyPermissionWithTeamPosition(ctx context.Context, userDo *system.User) (vobj.Role, *system.TeamMember, error) {
	memberDo, err := a.memberRepo.FindByUserID(ctx, userDo.ID)
	if err != nil {
		return 0, nil, err
	}
	if !memberDo.Status.IsNormal() {
		return 0, nil, merr.ErrorPermissionDenied("team member is invalid [%s]", memberDo.Status)
	}
	teamPosition, ok := permission.GetTeamPositionByContext(ctx)
	if !ok {
		return memberDo.Position, memberDo, nil
	}
	if memberDo.Position.GTE(teamPosition) {
		return teamPosition, memberDo, nil
	}
	return 0, nil, merr.ErrorPermissionDenied("Your current team role [%s] is not allowed to access this resource", teamPosition)
}

func (a *AuthBiz) verifyPermissionWithTeamRBAC(ctx context.Context, memberDo *system.TeamMember, resourceDo *system.Resource) error {
	if !resourceDo.Allow.IsTeamRBAC() {
		return nil
	}
	teamRoleID, ok := permission.GetTeamRoleIDByContext(ctx)
	if ok {
		teamRoleDo, ok := validate.SliceFindByValue(memberDo.Roles, teamRoleID, func(role *system.TeamRole) uint32 {
			return role.ID
		})
		if !ok {
			return merr.ErrorPermissionDenied("team role is invalid")
		}
		if !teamRoleDo.Status.IsNormal() {
			return merr.ErrorPermissionDenied("team role is invalid [%s]", teamRoleDo.Status)
		}
		_, ok = validate.SliceFindByValue(teamRoleDo.Resources, resourceDo.ID, func(role *system.Resource) uint32 {
			return role.ID
		})
		if ok {
			return nil
		}
		return merr.ErrorPermissionDenied("team role resource is invalid.")
	}
	resources := make([]*system.Resource, 0, len(memberDo.Roles)*10)
	for _, role := range memberDo.Roles {
		if role.Status.IsNormal() {
			resources = append(resources, role.Resources...)
		}
	}
	_, ok = validate.SliceFindByValue(resources, resourceDo.ID, func(role *system.Resource) uint32 {
		return role.ID
	})
	if ok {
		return nil
	}
	return merr.ErrorPermissionDenied("team role resource is invalid.")
}

// LoginByPassword login by password
func (a *AuthBiz) LoginByPassword(ctx context.Context, req *bo.LoginByPassword) (*bo.LoginSign, error) {
	user, err := a.userRepo.FindByEmail(ctx, crypto.String(req.Email))
	if err != nil {
		return nil, merr.ErrorPasswordError("password error").WithCause(err)
	}
	if !user.ValidatePassword(req.Password) {
		return nil, merr.ErrorPasswordError("password error")
	}
	return a.login(user)
}

// RefreshToken refresh token
func (a *AuthBiz) RefreshToken(ctx context.Context, req *bo.RefreshToken) (*bo.LoginSign, error) {
	if err := a.VerifyToken(ctx, req.Token); err != nil {
		return nil, err
	}
	userDo, err := a.userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := a.cacheRepo.BanToken(ctx, req.Token); err != nil {
			a.helper.Errorf("refresh token unban err: %v", err)
		}
	}()
	return a.login(userDo)
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

// GetOAuthConf 获取oauth配置
func (a *AuthBiz) GetOAuthConf(provider vobj.OAuthAPP) (*oauth2.Config, error) {
	config, ok := a.oauthConfigs.Get(provider)
	if !ok {
		return nil, merr.ErrorInternalServerError("not support oauth provider")
	}
	return config, nil
}

func (a *AuthBiz) OAuthLogin(ctx context.Context, provider vobj.OAuthAPP, code string) (string, error) {
	switch provider {
	case vobj.OAuthAPPGithub:
		return a.githubLogin(ctx, code)
	case vobj.OAuthAPPGitee:
		return a.giteeLogin(ctx, code)
	default:
		return "", merr.ErrorInternalServerError("not support oauth provider")
	}
}

func (a *AuthBiz) githubLogin(ctx context.Context, code string) (string, error) {
	githubOAuthConf, err := a.GetOAuthConf(vobj.OAuthAPPGithub)
	if err != nil {
		return "", err
	}

	token, err := githubOAuthConf.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	// 使用token来获取用户信息
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	userResp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return "", err
	}
	body := userResp.Body
	defer body.Close()
	var userInfo bo.GithubUser
	if err := json.NewDecoder(body).Decode(&userInfo); err != nil {
		return "", err
	}

	return a.oauthLogin(ctx, &userInfo)
}

func (a *AuthBiz) giteeLogin(ctx context.Context, code string) (string, error) {
	giteeOAuthConf, err := a.GetOAuthConf(vobj.OAuthAPPGitee)
	if err != nil {
		return "", err
	}
	opts := []oauth2.AuthCodeOption{
		// https://gitee.com/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&redirect_uri={redirect_uri}&client_secret={client_secret}
		oauth2.SetAuthURLParam("grant_type", "authorization_code"),
		oauth2.SetAuthURLParam("client_secret", giteeOAuthConf.ClientSecret),
		oauth2.SetAuthURLParam("client_id", giteeOAuthConf.ClientID),
		oauth2.SetAuthURLParam("redirect_uri", giteeOAuthConf.RedirectURL),
		oauth2.SetAuthURLParam("code", code),
	}
	token, err := giteeOAuthConf.Exchange(context.Background(), code, opts...)
	if err != nil {
		return "", err
	}
	// 使用token来获取用户信息
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	resp, err := client.Get("https://gitee.com/api/v5/user")
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	var userInfo bo.GiteeUser
	if err := json.NewDecoder(body).Decode(&userInfo); err != nil {
		return "", err
	}

	return a.oauthLogin(ctx, &userInfo)
}

func (a *AuthBiz) oauthUserFirstOrCreate(ctx context.Context, userInfo bo.IOAuthUser) (*system.OAuthUser, error) {
	oauthUserDoExist := true
	oauthUserDo, err := a.oauthRepo.FindByOAuthID(ctx, userInfo.GetOAuthID(), userInfo.GetAPP())
	if err != nil {
		if !merr.IsUserNotFound(err) {
			return nil, err
		}
		oauthUserDoExist = false
	}
	userDo, err := a.userRepo.FindByEmail(ctx, crypto.String(userInfo.GetEmail()))
	if err != nil {
		if !merr.IsUserNotFound(err) {
			return nil, err
		}
	}
	if userDo != nil {
		userInfo.WithUserID(userDo.ID)
	}
	err = a.transaction.MainExec(ctx, func(ctx context.Context) error {
		if !oauthUserDoExist {
			oauthUserDo, err = a.oauthRepo.Create(ctx, userInfo)
			if err != nil {
				return err
			}
		}
		if oauthUserDo.User == nil {
			// 创建用户
			userDo, err = a.userRepo.CreateUserWithOAuthUser(ctx, userInfo)
			if err != nil {
				return err
			}
			oauthUserDo.SysUserID = userDo.ID
			oauthUserDo.User = userDo
			oauthUserDo, err = a.oauthRepo.SetUser(ctx, oauthUserDo)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return oauthUserDo, nil
}

func (a *AuthBiz) oauthLogin(ctx context.Context, userInfo bo.IOAuthUser) (string, error) {
	oauthUserDo, err := a.oauthUserFirstOrCreate(ctx, userInfo)
	if err != nil {
		return "", err
	}

	if oauthUserDo.User == nil || validate.CheckEmail(userInfo.GetEmail()) != nil {
		oauthParams := &bo.OAuthLoginParams{
			OAuthID: oauthUserDo.ID,
			Token:   password.MD5(password.GenerateRandomPassword(64)),
		}
		if err := a.cacheRepo.CacheVerifyOAuthToken(ctx, oauthParams); err != nil {
			return "", err
		}
		redirect := fmt.Sprintf("%s?oauth_id=%d&app=%d&token=%s#/oauth/register/email", a.redirectURL, oauthParams.OAuthID, userInfo.GetAPP(), oauthParams.Token)
		return redirect, nil
	}

	loginSign, err := a.login(oauthUserDo.User)
	if err != nil {
		return "", err
	}
	redirect := fmt.Sprintf("%s?token=%s", a.redirectURL, loginSign.Token)
	return redirect, nil
}

// OAuthLoginWithEmail oauth2 set email login
func (a *AuthBiz) OAuthLoginWithEmail(ctx context.Context, oauthParams *bo.OAuthLoginParams) (*bo.LoginSign, error) {
	if err := a.cacheRepo.VerifyEmailCode(ctx, oauthParams.Email, oauthParams.Code); err != nil {
		return nil, err
	}

	if err := a.cacheRepo.VerifyOAuthToken(ctx, oauthParams); err != nil {
		return nil, err
	}

	oauthUserDo, err := a.oauthRepo.FindByOAuthID(ctx, oauthParams.OAuthID, oauthParams.APP)
	if err != nil {
		return nil, err
	}
	userDo := oauthUserDo.User
	if userDo == nil {
		return nil, merr.ErrorUnauthorized("oauth unauthorized").WithMetadata(map[string]string{
			"exist": "false",
		})
	}
	if userDo.Email.EQ(crypto.String(oauthParams.Email)) {
		return a.login(oauthUserDo.User)
	}

	userDo.Email = crypto.String(oauthParams.Email)
	userDo, err = a.userRepo.SetEmail(ctx, userDo)
	if err != nil {
		return nil, err
	}
	return a.login(userDo)
}

// VerifyEmail verify email
func (a *AuthBiz) VerifyEmail(ctx context.Context, email string) error {
	return a.cacheRepo.SendVerifyEmailCode(ctx, email)
}

// LoginWithEmail 邮箱登录
func (a *AuthBiz) LoginWithEmail(ctx context.Context, code string, user *system.User) (*bo.LoginSign, error) {
	if err := a.cacheRepo.VerifyEmailCode(ctx, string(user.Email), code); err != nil {
		return nil, err
	}
	userDo, err := a.userRepo.FindByEmail(ctx, user.Email)
	if err == nil {
		return a.login(userDo)
	}
	userDo = user
	userDo, err = a.userRepo.Create(ctx, userDo)
	if err != nil {
		return nil, err
	}
	return a.login(userDo)
}
