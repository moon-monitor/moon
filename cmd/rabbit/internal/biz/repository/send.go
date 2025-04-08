package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
)

type Send interface {
	Email(ctx context.Context, params bo.SendEmailParams) error
}
