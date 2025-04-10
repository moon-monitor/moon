package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data"
)

func NewConfigRepo(d *data.Data, logger log.Logger) repository.Config {
	return &configImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.config")),
	}
}

type configImpl struct {
	*data.Data
	helper *log.Helper
}

func (c *configImpl) GetMetricDatasourceConfig(ctx context.Context, id string) (bo.MetricDatasourceConfig, bool) {
	key := vobj.DatasourceCacheKey.Key()
	exist, err := c.Data.GetCache().Client().HExists(ctx, key, id).Result()
	if err != nil {
		c.helper.Errorw("method", "GetMetricDatasourceConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var metricDatasourceConfig do.DatasourceMetricConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, id).Scan(&metricDatasourceConfig); err != nil {
		c.helper.Errorw("method", "GetMetricDatasourceConfig", "err", err)
		return nil, false
	}
	return &metricDatasourceConfig, true
}

func (c *configImpl) SetMetricDatasourceConfig(ctx context.Context, configs ...bo.MetricDatasourceConfig) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		item := &do.DatasourceMetricConfig{
			ID:       v.GetId(),
			Driver:   v.GetDriver(),
			Endpoint: v.GetEndpoint(),
			Headers:  v.GetHeaders(),
			Method:   v.GetMethod(),
			CA:       v.GetCA(),
		}
		basicAuth := v.GetBasicAuth()
		if basicAuth != nil {
			item.BasicAuth = &do.BasicAuth{
				Username: basicAuth.GetUsername(),
				Password: basicAuth.GetPassword(),
			}
		}
		tls := v.GetTLS()
		if tls != nil {
			item.TLS = &do.TLS{
				ClientCertificate: tls.GetClientCertificate(),
				ClientKey:         tls.GetClientKey(),
				ServerName:        tls.GetServerName(),
			}
		}
		configDos[item.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.DatasourceCacheKey.Key(), configDos).Err()
}

func (c *configImpl) SetMetricRules(ctx context.Context, rules ...bo.MetricRule) error {
	configDos := make(map[string]any, len(rules))
	for _, v := range rules {
		item := &do.MetricRule{}
		configDos[v.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.MetricRuleCacheKey.Key(), configDos).Err()
}

func (c *configImpl) GetMetricRules(ctx context.Context) ([]bo.MetricRule, error) {
	key := vobj.MetricRuleCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.Errorw("method", "GetMetricRules", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}

	metricRulesMap, err := c.Data.GetCache().Client().HGetAll(ctx, key).Result()
	if err != nil {
		c.helper.Errorw("method", "GetMetricRules", "err", err)
		return nil, err
	}
	metricRules := make([]bo.MetricRule, 0, len(metricRulesMap))
	for _, v := range metricRulesMap {
		rule := new(do.MetricRule)
		if err := rule.UnmarshalBinary([]byte(v)); err != nil {
			continue
		}
		metricRules = append(metricRules, rule)
	}
	return metricRules, nil
}

func (c *configImpl) GetMetricRule(ctx context.Context, id string) (bo.MetricRule, bool) {
	key := vobj.MetricRuleCacheKey.Key()
	exist, err := c.Data.GetCache().Client().HExists(ctx, key, id).Result()
	if err != nil {
		c.helper.Errorw("method", "GetMetricRule", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var metricRule do.MetricRule
	if err := c.Data.GetCache().Client().HGet(ctx, key, id).Scan(&metricRule); err != nil {
		c.helper.Errorw("method", "GetMetricRule", "err", err)
		return nil, false
	}
	return &metricRule, true
}

func (c *configImpl) DeleteMetricRules(ctx context.Context, ids ...string) error {
	return c.Data.GetCache().Client().HDel(ctx, vobj.MetricRuleCacheKey.Key(), ids...).Err()
}
