package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToStrategyNotice(route do.NoticeGroup) *team.NoticeGroup {
	if validate.IsNil(route) {
		return nil
	}
	if notice, ok := route.(*team.NoticeGroup); ok {
		return notice
	}
	item := &team.NoticeGroup{
		Name:          route.GetName(),
		Remark:        route.GetRemark(),
		Status:        route.GetStatus(),
		Members:       ToStrategyMembers(route.GetNoticeMembers()),
		Hooks:         ToStrategyHooks(route.GetHooks()),
		EmailConfigID: route.GetEmailConfig().GetID(),
		EmailConfig:   ToStrategyEmailConfig(route.GetEmailConfig()),
		SMSConfigID:   route.GetSMSConfig().GetID(),
		SMSConfig:     ToStrategySmsConfig(route.GetSMSConfig()),
		TeamModel:     ToTeamModel(route),
	}
	return item
}

func ToStrategyNotices(routes []do.NoticeGroup) []*team.NoticeGroup {
	return slices.Map(routes, ToStrategyNotice)
}

func ToStrategyHook(hook do.NoticeHook) *team.NoticeHook {
	if validate.IsNil(hook) {
		return nil
	}
	if hook, ok := hook.(*team.NoticeHook); ok {
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
		NoticeGroups: ToStrategyNotices(hook.GetNoticeGroups()),
		APP:          hook.GetApp(),
	}
	return item
}

func ToStrategyHooks(hooks []do.NoticeHook) []*team.NoticeHook {
	return slices.Map(hooks, ToStrategyHook)
}

func ToStrategyEmailConfig(config do.TeamEmailConfig) *team.EmailConfig {
	if validate.IsNil(config) {
		return nil
	}
	if config, ok := config.(*team.EmailConfig); ok {
		return config
	}

	item := &team.EmailConfig{
		TeamModel: ToTeamModel(config),
		Name:      config.GetName(),
		Remark:    config.GetRemark(),
		Status:    config.GetStatus(),
		Email:     crypto.NewObject(config.GetEmailConfig()),
	}
	return item
}

func ToStrategySmsConfig(config do.TeamSMSConfig) *team.SmsConfig {
	if validate.IsNil(config) {
		return nil
	}
	if config, ok := config.(*team.SmsConfig); ok {
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
	return item
}

func ToStrategyMetricRuleLabelNotice(notice do.StrategyMetricRuleLabelNotice) *team.StrategyMetricRuleLabelNotice {
	if validate.IsNil(notice) {
		return nil
	}
	if notice, ok := notice.(*team.StrategyMetricRuleLabelNotice); ok {
		return notice
	}
	item := &team.StrategyMetricRuleLabelNotice{
		TeamModel:            ToTeamModel(notice),
		StrategyMetricRuleID: notice.GetStrategyMetricRuleID(),
		LabelKey:             notice.GetLabelKey(),
		LabelValue:           notice.GetLabelValue(),
		Notices:              ToStrategyNotices(notice.GetNotices()),
	}
	return item
}

func ToStrategyMetricRuleLabelNotices(notices []do.StrategyMetricRuleLabelNotice) []*team.StrategyMetricRuleLabelNotice {
	return slices.Map(notices, ToStrategyMetricRuleLabelNotice)
}
