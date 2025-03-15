package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewUserRepo(data *data.Data, logger log.Logger) repository.User {
	return &userRepoImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.user")),
	}
}

type userRepoImpl struct {
	*data.Data

	helper *log.Helper
}

func (u *userRepoImpl) FindByID(ctx context.Context, userID uint32) (*system.User, error) {
	userQuery := systemQuery.Use(u.GetMainDB().GetTx(ctx)).User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(userID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorUserNotFound("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (u *userRepoImpl) FindByEmail(ctx context.Context, email string) (*system.User, error) {
	userQuery := systemQuery.Use(u.GetMainDB().GetTx(ctx)).User
	user, err := userQuery.WithContext(ctx).Where(userQuery.Email.Eq(email)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorUserNotFound("user not found")
		}
		return nil, err
	}
	return user, nil
}
