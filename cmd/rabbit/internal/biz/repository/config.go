package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
)

type Config interface {
	GetEmailConfig(ctx context.Context, name string) (bo.EmailConfig, bool)
	GetEmailConfigs(ctx context.Context, names ...string) ([]bo.EmailConfig, error)
	SetEmailConfig(ctx context.Context, configs ...bo.EmailConfig) error

	GetSMSConfig(ctx context.Context, name string) (bo.SMSConfig, bool)
	GetSMSConfigs(ctx context.Context, names ...string) ([]bo.SMSConfig, error)
	SetSMSConfig(ctx context.Context, configs ...bo.SMSConfig) error

	GetHookConfig(ctx context.Context, name string) (bo.HookConfig, bool)
	GetHookConfigs(ctx context.Context, names ...string) ([]bo.HookConfig, error)
	SetHookConfig(ctx context.Context, configs ...bo.HookConfig) error

	GetNoticeGroupConfig(ctx context.Context, name string) (bo.NoticeGroup, bool)
	GetNoticeGroupConfigs(ctx context.Context, names ...string) ([]bo.NoticeGroup, error)
	SetNoticeGroupConfig(ctx context.Context, configs ...bo.NoticeGroup) error

	GetNoticeUserConfig(ctx context.Context, name string) (bo.NoticeUser, bool)
	GetNoticeUserConfigs(ctx context.Context, names ...string) ([]bo.NoticeUser, error)
	SetNoticeUserConfig(ctx context.Context, configs ...bo.NoticeUser) error
}
