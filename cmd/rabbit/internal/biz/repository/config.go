package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
)

type Config interface {
	GetEmailConfig(ctx context.Context, name string) (bo.EmailConfig, bool)
	SetEmailConfig(ctx context.Context, configs ...bo.EmailConfig) error
	GetSMSConfig(ctx context.Context, name string) (bo.SMSConfig, bool)
	SetSMSConfig(ctx context.Context, configs ...bo.SMSConfig) error
}
