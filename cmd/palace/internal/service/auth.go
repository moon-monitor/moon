package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
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

func (s *AuthService) LoginByEmail(ctx context.Context, req *palacev1.LoginByEmailRequest) (*palacev1.LoginByEmailReply, error) {
	return &palacev1.LoginByEmailReply{}, nil
}
