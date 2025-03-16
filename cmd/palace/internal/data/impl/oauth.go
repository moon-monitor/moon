package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
)

func NewOAuthRepo(data *data.Data, logger log.Logger) repository.OAuth {
	return &oauthRepoImpl{
		Data:     data,
		userRepo: NewUserRepo(data, logger),
		helper:   log.NewHelper(log.With(logger, "module", "data.repo.oauth")),
	}
}

type oauthRepoImpl struct {
	*data.Data

	userRepo repository.User

	helper *log.Helper
}

func (o *oauthRepoImpl) OAuthUserFirstOrCreate(ctx context.Context, oauthUser bo.IOAuthUser) (*system.User, error) {
	oauthUserQuery := systemQuery.Use(o.GetMainDB().GetTx(ctx)).OAuthUser
	oauthUserDo, err := oauthUserQuery.WithContext(ctx).
		Where(oauthUserQuery.OAuthID.Eq(oauthUser.GetOAuthID()), oauthUserQuery.APP.Eq(oauthUser.GetAPP().GetValue())).
		First()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		// create oauth user
		return o.createOAuthUser(ctx, oauthUser)
	}
	userDo, err := o.findUserByEmail(ctx, oauthUser.GetEmail())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		// create user
		return o.createUserWithOAuthUser(ctx, oauthUserDo)
	}
	return userDo, nil
}

func (o *oauthRepoImpl) createOAuthUser(ctx context.Context, oauthUser bo.IOAuthUser) (*system.User, error) {
	userDo, err := o.userRepo.FindByEmail(ctx, oauthUser.GetEmail())
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		// create user
		return o.createUserWithOAuthUser(ctx, oauthUserDo)
	}
	oauthUserQuery := systemQuery.Use(o.GetMainDB().GetTx(ctx)).OAuthUser
	oauthUserDo := &system.OAuthUser{
		OAuthID:   oauthUser.GetOAuthID(),
		SysUserID: userDo.ID,
		Row:       oauthUser.String(),
		APP:       oauthUser.GetAPP(),
		User:      userDo,
	}

	if err := oauthUserQuery.WithContext(ctx).Create(oauthUserDo); err != nil {
		return nil, err
	}
	return userDo, nil
}

func (o *oauthRepoImpl) SetEmail(ctx context.Context, userID uint32, email string) (*system.User, error) {
	//TODO implement me
	panic("implement me")
}

func (o *oauthRepoImpl) GetSysUserByOAuthID(ctx context.Context, oauthID uint32, app vobj.OAuthAPP) (*system.OAuthUser, error) {
	//TODO implement me
	panic("implement me")
}

func (o *oauthRepoImpl) SendVerifyEmail(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (o *oauthRepoImpl) CheckVerifyEmailCode(ctx context.Context, email, code string) error {
	//TODO implement me
	panic("implement me")
}
