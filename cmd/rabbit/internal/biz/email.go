package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
)

func NewEmail(logger log.Logger) *Email {
	return &Email{helper: log.NewHelper(log.With(logger, "module", "biz.email"))}
}

type Email struct {
	helper *log.Helper
}

func (e *Email) Send(ctx context.Context, params bo.SendEmailParams) error {
	return nil
}
