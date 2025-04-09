package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
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

func (c *Config) GetMetricDatasourceConfig(ctx context.Context, id *string, defaultConfig bo.MetricDatasourceConfig) bo.MetricDatasourceConfig {
	if id == nil || *id == "" {
		return defaultConfig
	}
	metricDatasourceConfig, ok := c.configRepo.GetMetricDatasourceConfig(ctx, *id)
	if !ok || !metricDatasourceConfig.GetEnable() {
		return defaultConfig
	}
	return metricDatasourceConfig
}

func (c *Config) SetMetricDatasourceConfig(ctx context.Context, configs ...bo.MetricDatasourceConfig) error {
	if len(configs) == 0 {
		return nil
	}
	return c.configRepo.SetMetricDatasourceConfig(ctx, configs...)
}

func (c *Config) SetMetricRules(ctx context.Context, rules ...bo.MetricRule) error {
	if len(rules) == 0 {
		return nil
	}
	return c.configRepo.SetMetricRules(ctx, rules...)
}

func (c *Config) GetMetricRule(ctx context.Context, id string) (bo.MetricRule, bool) {
	return c.configRepo.GetMetricRule(ctx, id)
}

func (c *Config) RemoveMetricRules(ctx context.Context, ids ...string) error {
	if len(ids) == 0 {
		return nil
	}
	return c.configRepo.DeleteMetricRules(ctx, ids...)
}
