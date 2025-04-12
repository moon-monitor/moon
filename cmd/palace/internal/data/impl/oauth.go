package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
)

func NewOAuthRepo(data *data.Data, logger log.Logger) repository.OAuth {
	return &oauthRepoImpl{
		Data:   data,
		Query:  systemgen.Use(data.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.oauth")),
	}
}

type oauthRepoImpl struct {
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func toOAuthUserDo(u bo.IOAuthUser) *system.UserOAuth {
	if u == nil {
		return nil
	}
	return &system.UserOAuth{
		OAuthID: u.GetOAuthID(),
		UserID:  u.GetUserID(),
		Row:     u.String(),
		APP:     u.GetAPP(),
		User:    nil,
	}
}

func (o *oauthRepoImpl) Create(ctx context.Context, user bo.IOAuthUser) (*system.UserOAuth, error) {
	oauthUserDo := toOAuthUserDo(user)
	oauthUserMutation := o.UserOAuth
	if err := oauthUserMutation.WithContext(ctx).Create(oauthUserDo); err != nil {
		return nil, err
	}
	return oauthUserMutation.WithContext(ctx).
		Where(oauthUserMutation.OAuthID.Eq(oauthUserDo.OAuthID)).
		Where(oauthUserMutation.APP.Eq(oauthUserDo.APP.GetValue())).
		Preload(oauthUserMutation.User).
		First()
}

func (o *oauthRepoImpl) FindByOAuthID(ctx context.Context, oauthID uint32, app vobj.OAuthAPP) (*system.UserOAuth, error) {
	oauthUserMutation := o.UserOAuth
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

func (o *oauthRepoImpl) SetUser(ctx context.Context, user *system.UserOAuth) (*system.UserOAuth, error) {
	oauthUserMutation := o.UserOAuth
	wrapper := []gen.Condition{
		oauthUserMutation.ID.Eq(user.ID),
		oauthUserMutation.APP.Eq(user.APP.GetValue()),
		oauthUserMutation.OAuthID.Eq(user.OAuthID),
	}

	if _, err := oauthUserMutation.WithContext(ctx).Where(wrapper...).UpdateSimple(oauthUserMutation.UserID.Value(user.UserID)); err != nil {
		return nil, err
	}
	return oauthUserMutation.WithContext(ctx).Where(wrapper...).Preload(oauthUserMutation.User).First()
}
