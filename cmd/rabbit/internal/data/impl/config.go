package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
)

func NewConfigRepo(d *data.Data, logger log.Logger) repository.Config {
	return &configImpl{
		helper: log.NewHelper(log.With(logger, "module", "data.repo.config")),
		Data:   d,
	}
}

type configImpl struct {
	helper *log.Helper
	*data.Data
}

func (c *configImpl) GetEmailConfig(ctx context.Context, name string) (bo.EmailConfig, bool) {
	key := vobj.EmailCacheKey.Key()
	exist, err := c.Data.GetCache().Client().HExists(ctx, key, name).Result()
	if err != nil {
		c.helper.Errorw("method", "GetEmailConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var emailConfig do.EmailConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&emailConfig); err != nil {
		c.helper.Errorw("method", "GetEmailConfig", "err", err)
		return nil, false
	}

	return &emailConfig, true
}

func (c *configImpl) SetEmailConfig(ctx context.Context, configs ...bo.EmailConfig) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		item := &do.EmailConfig{
			User:   v.GetUser(),
			Pass:   v.GetPass(),
			Host:   v.GetHost(),
			Port:   v.GetPort(),
			Enable: v.GetEnable(),
			Name:   v.GetName(),
		}
		configDos[item.UniqueKey()] = item
	}

	return c.Data.GetCache().Client().HSet(ctx, vobj.EmailCacheKey.Key(), configDos).Err()
}

func (c *configImpl) GetSMSConfig(ctx context.Context, name string) (bo.SMSConfig, bool) {
	key := vobj.SmsCacheKey.Key()
	exist, err := c.Data.GetCache().Client().HExists(ctx, key, name).Result()
	if err != nil {
		c.helper.Errorw("method", "GetSMSConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var smsConfig do.SMSConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&smsConfig); err != nil {
		c.helper.Errorw("method", "GetSMSConfig", "err", err)
		return nil, false
	}
	return &smsConfig, true
}

func (c *configImpl) SetSMSConfig(ctx context.Context, configs ...bo.SMSConfig) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		item := &do.SMSConfig{
			AccessKeyId:     v.GetAccessKeyId(),
			AccessKeySecret: v.GetAccessKeySecret(),
			Endpoint:        v.GetEndpoint(),
			Name:            v.GetName(),
			SignName:        v.GetSignName(),
			Type:            v.GetType(),
			Enable:          v.GetEnable(),
		}
		configDos[item.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.SmsCacheKey.Key(), configDos).Err()
}

func (c *configImpl) GetHookConfig(ctx context.Context, name string) (bo.HookConfig, bool) {
	key := vobj.HookCacheKey.Key()
	exist, err := c.Data.GetCache().Client().HExists(ctx, key, name).Result()
	if err != nil {
		c.helper.Errorw("method", "GetHookConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var hookConfig do.HookConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&hookConfig); err != nil {
		c.helper.Errorw("method", "GetHookConfig", "err", err)
		return nil, false
	}
	return &hookConfig, true
}

func (c *configImpl) SetHookConfig(ctx context.Context, configs ...bo.HookConfig) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		item := &do.HookConfig{
			Name:     v.GetName(),
			App:      v.GetApp(),
			Url:      v.GetUrl(),
			Secret:   v.GetSecret(),
			Token:    v.GetToken(),
			Username: v.GetUsername(),
			Password: v.GetPassword(),
			Headers:  v.GetHeaders(),
			Enable:   v.GetEnable(),
		}
		configDos[item.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.HookCacheKey.Key(), configDos).Err()
}
