package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

// UserBiz is a user business logic implementation.
type UserBiz struct {
	user repository.User
	log  *log.Helper
}

// NewUserBiz creates a new UserBiz instance.
func NewUserBiz(user repository.User, logger log.Logger) *UserBiz {
	return &UserBiz{
		user: user,
		log:  log.NewHelper(log.With(logger, "module", "biz/user")),
	}
}

// GetSelfInfo retrieves the current user's information from the context.
func (b *UserBiz) GetSelfInfo(ctx context.Context) (*system.User, error) {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("user not found in context")
	}

	user, err := b.user.FindByID(ctx, userID)
	if err != nil {
		return nil, merr.ErrorInternalServerError("failed to find user").WithCause(err)
	}

	if user == nil {
		return nil, merr.ErrorUserNotFound("user not found")
	}

	return user, nil
}
