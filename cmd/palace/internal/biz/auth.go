package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewAuthBiz(userRepo repository.User, logger log.Logger) *AuthBiz {
	return &AuthBiz{
		userRepo: userRepo,
		helper:   log.NewHelper(log.With(logger, "module", "biz.auth")),
	}
}

type AuthBiz struct {
	userRepo repository.User

	helper *log.Helper
}
