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
	if !ok {
		return defaultConfig
	}
	return emailConfig
}
