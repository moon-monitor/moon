package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
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
	//TODO implement me
	panic("implement me")
}

func (u *userRepoImpl) FindByEmail(ctx context.Context, email string) (*system.User, error) {
	//TODO implement me
	panic("implement me")
}
