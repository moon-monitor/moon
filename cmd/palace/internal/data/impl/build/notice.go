package build

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToStrategyNotice(ctx context.Context, route do.NoticeGroup) *team.NoticeGroup {
	if notice, ok := route.(*team.NoticeGroup); ok {
		notice.WithContext(ctx)
		return notice
	}
	item := &team.NoticeGroup{
		Name:          route.GetName(),
		Remark:        route.GetRemark(),
		Status:        route.GetStatus(),
		Members:       ToStrategyMembers(ctx, route.GetNoticeMembers()),
		Hooks:         ToStrategyHooks(ctx, route.GetHooks()),
		EmailConfigID: route.GetEmailConfig().GetID(),
		EmailConfig:   ToStrategyEmailConfig(ctx, route.GetEmailConfig()),
		SMSConfigID:   route.GetSMSConfig().GetID(),
		SMSConfig:     ToStrategySmsConfig(ctx, route.GetSMSConfig()),
		TeamModel:     ToTeamModel(route),
	}
	item.WithContext(ctx)
	return item
}

func ToStrategyNotices(ctx context.Context, routes []do.NoticeGroup) []*team.NoticeGroup {
	return slices.Map(routes, func(route do.NoticeGroup) *team.NoticeGroup {
		return ToStrategyNotice(ctx, route)
	})
}

func ToStrategyHook(ctx context.Context, hook do.NoticeHook) *team.NoticeHook {
	if hook, ok := hook.(*team.NoticeHook); ok {
		hook.WithContext(ctx)
		return hook
	}
	item := &team.NoticeHook{
		TeamModel:    ToTeamModel(hook),
		Name:         hook.GetName(),
		Remark:       hook.GetRemark(),
		Status:       hook.GetStatus(),
		URL:          hook.GetURL(),
		Method:       hook.GetMethod(),
		Secret:       hook.GetSecret(),
		Headers:      hook.GetHeaders(),
		NoticeGroups: ToStrategyNotices(ctx, hook.GetNoticeGroups()),
		APP:          hook.GetApp(),
	}
	item.WithContext(ctx)
	return item
}

func ToStrategyHooks(ctx context.Context, hooks []do.NoticeHook) []*team.NoticeHook {
	return slices.Map(hooks, func(hook do.NoticeHook) *team.NoticeHook {
		return ToStrategyHook(ctx, hook)
	})
}

func ToStrategyEmailConfig(ctx context.Context, config do.TeamEmailConfig) *team.EmailConfig {
	if config, ok := config.(*team.EmailConfig); ok {
		config.WithContext(ctx)
		return config
	}

	item := &team.EmailConfig{
		TeamModel: ToTeamModel(config),
		Name:      config.GetName(),
		Remark:    config.GetRemark(),
		Status:    config.GetStatus(),
		Email:     crypto.NewObject(config.GetEmailConfig()),
	}
	item.WithContext(ctx)
	return item
}

func ToStrategySmsConfig(ctx context.Context, config do.TeamSMSConfig) *team.SmsConfig {
	if config, ok := config.(*team.SmsConfig); ok {
		config.WithContext(ctx)
		return config
	}
	item := &team.SmsConfig{
		TeamModel: ToTeamModel(config),
		Name:      config.GetName(),
		Remark:    config.GetRemark(),
		Status:    config.GetStatus(),
		Sms:       crypto.NewObject(config.GetSMSConfig()),
		Provider:  config.GetProviderType(),
	}
	item.WithContext(ctx)
	return item
}
