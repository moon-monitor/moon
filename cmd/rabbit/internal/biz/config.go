package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
)

func NewConfig(configRepo repository.Config, logger log.Logger) *Config {
	return &Config{
		helper:     log.NewHelper(log.With(logger, "module", "biz.config")),
		configRepo: configRepo,
	}
}

type Config struct {
	helper     *log.Helper
	configRepo repository.Config
}

func (c *Config) GetEmailConfig(ctx context.Context, name *string, defaultConfig bo.EmailConfig) bo.EmailConfig {
	if name == nil || *name == "" {
		return defaultConfig
	}
	emailConfig, ok := c.configRepo.GetEmailConfig(ctx, *name)
	if !ok || !emailConfig.GetEnable() {
		return defaultConfig
	}
	return emailConfig
}

func (c *Config) SetEmailConfig(ctx context.Context, configs ...bo.EmailConfig) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetEmailConfig(ctx, configs...)
}

func (c *Config) GetSMSConfig(ctx context.Context, name *string, defaultConfig bo.SMSConfig) bo.SMSConfig {
	if name == nil || *name == "" {
		return defaultConfig
	}
	smsConfig, ok := c.configRepo.GetSMSConfig(ctx, *name)
	if !ok || !smsConfig.GetEnable() {
		return defaultConfig
	}
	return smsConfig
}

func (c *Config) SetSMSConfig(ctx context.Context, configs ...bo.SMSConfig) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetSMSConfig(ctx, configs...)
}

func (c *Config) GetHookConfig(ctx context.Context, name *string, defaultConfig bo.HookConfig) bo.HookConfig {
	if name == nil || *name == "" {
		return defaultConfig
	}
	hookConfig, ok := c.configRepo.GetHookConfig(ctx, *name)
	if !ok || !hookConfig.GetEnable() {
		return defaultConfig
	}
	return hookConfig
}

func (c *Config) SetHookConfig(ctx context.Context, configs ...bo.HookConfig) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetHookConfig(ctx, configs...)
}

func (c *Config) GetNoticeGroupConfig(ctx context.Context, name *string, defaultConfig bo.NoticeGroup) bo.NoticeGroup {
	if name == nil || *name == "" {
		return defaultConfig
	}
	noticeGroupConfig, ok := c.configRepo.GetNoticeGroupConfig(ctx, *name)
	if !ok {
		return defaultConfig
	}
	return noticeGroupConfig
}

func (c *Config) SetNoticeGroupConfig(ctx context.Context, configs ...bo.NoticeGroup) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetNoticeGroupConfig(ctx, configs...)
}

func (c *Config) GetNoticeUserConfig(ctx context.Context, name *string, defaultConfig bo.NoticeUser) bo.NoticeUser {
	if name == nil || *name == "" {
		return defaultConfig
	}
	noticeUserConfig, ok := c.configRepo.GetNoticeUserConfig(ctx, *name)
	if !ok {
		return defaultConfig
	}
	return noticeUserConfig
}

func (c *Config) SetNoticeUserConfig(ctx context.Context, configs ...bo.NoticeUser) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetNoticeUserConfig(ctx, configs...)
}
