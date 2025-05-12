package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
	"github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/merr"
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
		c.helper.WithContext(ctx).Errorw("method", "GetEmailConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var emailConfig do.EmailConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&emailConfig); err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetEmailConfig", "err", err)
		return nil, false
	}

	return &emailConfig, true
}

func (c *configImpl) GetEmailConfigs(ctx context.Context, names ...string) ([]bo.EmailConfig, error) {
	key := vobj.EmailCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetEmailConfig", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}
	all, err := c.Data.GetCache().Client().HMGet(ctx, key, names...).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetEmailConfig", "err", err)
		return nil, err
	}

	emailConfigs := make([]bo.EmailConfig, 0, len(all))
	for _, v := range all {
		if v == nil {
			continue
		}
		var emailConfig do.EmailConfig
		switch val := v.(type) {
		case []byte:
			if err := emailConfig.UnmarshalBinary(val); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetEmailConfig", "err", err)
				return nil, err
			}
		case string:
			if err := emailConfig.UnmarshalBinary([]byte(val)); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetEmailConfig", "err", err)
				return nil, err
			}
		default:
			c.helper.WithContext(ctx).Warnw("method", "GetEmailConfig", "err", merr.ErrorParamsError("invalid email config"))
			continue
		}
		emailConfigs = append(emailConfigs, &emailConfig)
	}
	return emailConfigs, nil
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
		c.helper.WithContext(ctx).Errorw("method", "GetSMSConfig", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var smsConfig do.SMSConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&smsConfig); err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetSMSConfig", "err", err)
		return nil, false
	}
	return &smsConfig, true
}

func (c *configImpl) GetSMSConfigs(ctx context.Context, names ...string) ([]bo.SMSConfig, error) {
	key := vobj.SmsCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetSMSConfig", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}
	all, err := c.Data.GetCache().Client().HMGet(ctx, key, names...).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetSMSConfig", "err", err)
		return nil, err
	}

	smsConfigs := make([]bo.SMSConfig, 0, len(all))
	for _, v := range all {
		if v == nil {
			continue
		}
		var smsConfig do.SMSConfig
		switch val := v.(type) {
		case []byte:
			if err := smsConfig.UnmarshalBinary(val); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetSMSConfig", "err", err)
				return nil, err
			}
		case string:
			if err := smsConfig.UnmarshalBinary([]byte(val)); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetSMSConfig", "err", err)
				return nil, err
			}
		default:
			c.helper.WithContext(ctx).Warnw("method", "GetSMSConfig", "err", merr.ErrorParamsError("invalid sms config"))
			continue
		}
		smsConfigs = append(smsConfigs, &smsConfig)
	}
	return smsConfigs, nil
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
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetHookConfig", "err", err)
		return nil, false
	}
	if exist == 0 {
		return nil, false
	}
	var hookConfig do.HookConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&hookConfig); err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetHookConfig", "err", err)
		return nil, false
	}
	return &hookConfig, true
}

func (c *configImpl) GetHookConfigs(ctx context.Context, names ...string) ([]bo.HookConfig, error) {
	key := vobj.HookCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetHookConfig", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}
	all, err := c.Data.GetCache().Client().HMGet(ctx, key, names...).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetHookConfig", "err", err)
		return nil, err
	}

	hookConfigs := make([]bo.HookConfig, 0, len(all))
	for _, v := range all {
		if v == nil {
			continue
		}
		var hookConfig do.HookConfig
		switch val := v.(type) {
		case []byte:
			if err := hookConfig.UnmarshalBinary(val); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetHookConfig", "err", err)
				return nil, err
			}
		case string:
			if err := hookConfig.UnmarshalBinary([]byte(val)); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetHookConfig", "err", err)
				return nil, err
			}
		default:
			c.helper.WithContext(ctx).Warnw("method", "GetHookConfig", "err", merr.ErrorParamsError("invalid hook config"))
			continue
		}
		hookConfigs = append(hookConfigs, &hookConfig)
	}
	return hookConfigs, nil
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

func (c *configImpl) GetNoticeGroupConfig(ctx context.Context, name string) (bo.NoticeGroup, bool) {
	key := vobj.NoticeGroupCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeGroupConfig", "err", err)
		return nil, false
	}
	if exist == 0 {
		return nil, false
	}
	var noticeGroupConfig do.NoticeGroupConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&noticeGroupConfig); err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeGroupConfig", "err", err)
		return nil, false
	}
	return &noticeGroupConfig, true
}

func (c *configImpl) GetNoticeGroupConfigs(ctx context.Context, names ...string) ([]bo.NoticeGroup, error) {
	key := vobj.NoticeGroupCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeGroupConfig", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}
	all, err := c.Data.GetCache().Client().HMGet(ctx, key, names...).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeGroupConfig", "err", err)
		return nil, err
	}

	noticeGroupConfigs := make([]bo.NoticeGroup, 0, len(all))
	for _, v := range all {
		if v == nil {
			continue
		}
		var noticeGroupConfig do.NoticeGroupConfig
		switch val := v.(type) {
		case []byte:
			if err := noticeGroupConfig.UnmarshalBinary(val); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetNoticeGroupConfig", "err", err)
				return nil, err
			}
		case string:
			if err := noticeGroupConfig.UnmarshalBinary([]byte(val)); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetNoticeGroupConfig", "err", err)
				return nil, err
			}
		default:
			c.helper.WithContext(ctx).Warnw("method", "GetNoticeGroupConfig", "err", merr.ErrorParamsError("invalid notice group config"))
			continue
		}
		noticeGroupConfigs = append(noticeGroupConfigs, &noticeGroupConfig)
	}
	return noticeGroupConfigs, nil
}

func (c *configImpl) SetNoticeGroupConfig(ctx context.Context, configs ...bo.NoticeGroup) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		templateMap := make(map[common.NoticeType]*do.Template, len(v.GetTemplates()))
		for _, t := range v.GetTemplates() {
			templateMap[t.GetType()] = &do.Template{
				Type:           t.GetType(),
				Template:       t.GetTemplate(),
				TemplateParams: t.GetTemplateParameters(),
			}
		}
		item := &do.NoticeGroupConfig{
			Name:            v.GetName(),
			SMSConfigName:   v.GetSmsConfigName(),
			EmailConfigName: v.GetEmailConfigName(),
			HookConfigNames: v.GetHookConfigNames(),
			SMSUserNames:    v.GetSmsUserNames(),
			EmailUserNames:  v.GetEmailUserNames(),
			Templates:       templateMap,
		}
		configDos[item.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.NoticeGroupCacheKey.Key(), configDos).Err()
}

func (c *configImpl) GetNoticeUserConfig(ctx context.Context, name string) (bo.NoticeUser, bool) {
	key := vobj.NoticeUserCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeUserConfig", "err", err)
		return nil, false
	}
	if exist == 0 {
		return nil, false
	}
	var noticeUserConfig do.NoticeUserConfig
	if err := c.Data.GetCache().Client().HGet(ctx, key, name).Scan(&noticeUserConfig); err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeUserConfig", "err", err)
		return nil, false
	}
	return &noticeUserConfig, true
}

func (c *configImpl) GetNoticeUserConfigs(ctx context.Context, names ...string) ([]bo.NoticeUser, error) {
	key := vobj.NoticeUserCacheKey.Key()
	exist, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeUserConfig", "err", err)
		return nil, err
	}
	if exist == 0 {
		return nil, nil
	}
	all, err := c.Data.GetCache().Client().HMGet(ctx, key, names...).Result()
	if err != nil {
		c.helper.WithContext(ctx).Errorw("method", "GetNoticeUserConfig", "err", err)
		return nil, err
	}

	noticeUserConfigs := make([]bo.NoticeUser, 0, len(all))
	for _, v := range all {
		if v == nil {
			continue
		}
		var noticeUserConfig do.NoticeUserConfig
		switch val := v.(type) {
		case []byte:
			if err := noticeUserConfig.UnmarshalBinary(val); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetNoticeUserConfig", "err", err)
				return nil, err
			}
		case string:
			if err := noticeUserConfig.UnmarshalBinary([]byte(val)); err != nil {
				c.helper.WithContext(ctx).Warnw("method", "GetNoticeUserConfig", "err", err)
				return nil, err
			}
		default:
			c.helper.WithContext(ctx).Warnw("method", "GetNoticeUserConfig", "err", merr.ErrorParamsError("invalid notice user config"))
			continue
		}
		noticeUserConfigs = append(noticeUserConfigs, &noticeUserConfig)
	}
	return noticeUserConfigs, nil
}

func (c *configImpl) SetNoticeUserConfig(ctx context.Context, configs ...bo.NoticeUser) error {
	configDos := make(map[string]any, len(configs))
	for _, v := range configs {
		item := &do.NoticeUserConfig{
			Name:  v.GetName(),
			Email: v.GetEmail(),
			Phone: v.GetPhone(),
		}
		configDos[item.UniqueKey()] = item
	}
	return c.Data.GetCache().Client().HSet(ctx, vobj.NoticeUserCacheKey.Key(), configDos).Err()
}
