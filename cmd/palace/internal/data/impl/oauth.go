package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewOAuthRepo(data *data.Data, logger log.Logger) repository.OAuth {
	return &oauthRepoImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.oauth")),
	}
}

type oauthRepoImpl struct {
	*data.Data

	helper *log.Helper
}

func toOAuthUserDo(u bo.IOAuthUser) *system.OAuthUser {
	if u == nil {
		return nil
	}
	return &system.OAuthUser{
		OAuthID:   u.GetOAuthID(),
		SysUserID: u.GetUserID(),
		Row:       u.String(),
		APP:       u.GetAPP(),
		User:      nil,
	}
}

func oauthUserNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorUserNotFound("oauth user not found").WithCause(err)
	}
	return err
}

func (o *oauthRepoImpl) Create(ctx context.Context, user bo.IOAuthUser) (*system.OAuthUser, error) {
	oauthUserDo := toOAuthUserDo(user)
	oauthUserMutation := systemQuery.Use(o.GetMainDB().GetDB()).OAuthUser
	if err := oauthUserMutation.WithContext(ctx).Create(oauthUserDo); err != nil {
		return nil, err
	}
	return oauthUserMutation.WithContext(ctx).
		Where(oauthUserMutation.OAuthID.Eq(oauthUserDo.OAuthID)).
		Where(oauthUserMutation.APP.Eq(oauthUserDo.APP.GetValue())).
		Preload(oauthUserMutation.User).
		First()
}

func (o *oauthRepoImpl) FindByOAuthID(ctx context.Context, oauthID uint32, app vobj.OAuthAPP) (*system.OAuthUser, error) {
	oauthUserMutation := systemQuery.Use(o.GetMainDB().GetDB()).OAuthUser
	oauthUserDo, err := oauthUserMutation.WithContext(ctx).
		Where(oauthUserMutation.OAuthID.Eq(oauthID)).
		Where(oauthUserMutation.APP.Eq(app.GetValue())).
		Preload(oauthUserMutation.User).
		First()
	if err != nil {
		return nil, oauthUserNotFound(err)
	}
	return oauthUserDo, nil
}

func (o *oauthRepoImpl) SetUser(ctx context.Context, user *system.OAuthUser) (*system.OAuthUser, error) {
	oauthUserMutation := systemQuery.Use(o.GetMainDB().GetDB()).OAuthUser
	wrapper := []gen.Condition{
		oauthUserMutation.ID.Eq(user.ID),
		oauthUserMutation.APP.Eq(user.APP.GetValue()),
		oauthUserMutation.OAuthID.Eq(user.OAuthID),
	}

	if _, err := oauthUserMutation.WithContext(ctx).Where(wrapper...).UpdateSimple(oauthUserMutation.SysUserID.Value(user.SysUserID)); err != nil {
		return nil, err
	}
	return oauthUserMutation.WithContext(ctx).Where(wrapper...).Preload(oauthUserMutation.User).First()
}
