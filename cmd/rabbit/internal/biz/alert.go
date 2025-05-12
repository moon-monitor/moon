package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/template"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewAlert(
	configRepo repository.Config,
	sendRepo repository.Send,
	logger log.Logger,
) *Alert {
	return &Alert{
		configRepo: configRepo,
		sendRepo:   sendRepo,
		helper:     log.NewHelper(log.With(logger, "module", "biz.alert")),
	}
}

type Alert struct {
	configRepo repository.Config
	sendRepo   repository.Send

	helper *log.Helper
}

func (a *Alert) SendAlert(ctx context.Context, alert *bo.AlertsItem) error {
	if validate.IsNil(alert) {
		return merr.ErrorParamsError("No alert is available")
	}
	receivers := alert.GetReceiver()
	if len(receivers) == 0 {
		return merr.ErrorParamsError("No receiver is available")
	}
	for _, receiver := range receivers {
		noticeGroupConfig, ok := a.configRepo.GetNoticeGroupConfig(ctx, receiver)
		if !ok || validate.IsNil(noticeGroupConfig) {
			continue
		}
		a.sendEmail(ctx, noticeGroupConfig, alert)
		a.sendSms(ctx, noticeGroupConfig, alert)
		a.sendHook(ctx, noticeGroupConfig, alert)
	}
	return nil
}

func (a *Alert) sendEmail(ctx context.Context, noticeGroupConfig bo.NoticeGroup, alert *bo.AlertsItem) {
	emailNames := noticeGroupConfig.GetEmailUserNames()
	if len(emailNames) == 0 {
		return
	}
}

func (a *Alert) sendSms(ctx context.Context, noticeGroupConfig bo.NoticeGroup, alert *bo.AlertsItem) {
	smsNames := noticeGroupConfig.GetSmsUserNames()
	if len(smsNames) == 0 {
		return
	}
}

func (a *Alert) sendHook(ctx context.Context, noticeGroupConfig bo.NoticeGroup, alert *bo.AlertsItem) {
	hookNames := noticeGroupConfig.GetHookConfigNames()
	if len(hookNames) == 0 {
		return
	}
	hookConfigs := make([]bo.HookConfig, 0, len(hookNames))
	body := make([]*bo.HookBody, 0, len(hookNames)*len(alert.Alerts))

	for _, hookName := range hookNames {
		hookConfig, ok := a.configRepo.GetHookConfig(ctx, hookName)
		if !ok || validate.IsNil(hookConfig) {
			continue
		}
		hookConfigs = append(hookConfigs, hookConfig)
		for _, alertItem := range alert.Alerts {
			body = append(body, &bo.HookBody{
				AppName: hookConfig.GetName(),
				Body:    a.getHookBody(noticeGroupConfig.GetHookTemplate(hookConfig.GetApp()), alertItem),
			})
		}
	}
	sendParamsOpts := []bo.SendHookParamsOption{
		bo.WithSendHookParamsOptionBody(body),
	}
	sendHookParams, err := bo.NewSendHookParams(hookConfigs, sendParamsOpts...)
	if err != nil {
		a.helper.WithContext(ctx).Warnw("method", "NewSendHookParams", "err", err)
		return
	}
	if err := a.sendRepo.Hook(ctx, sendHookParams); err != nil {
		a.helper.WithContext(ctx).Warnw("method", "sendRepo.Hook", "err", err)
	}
}

func (a *Alert) getHookBody(temp string, alert *bo.AlertItem) []byte {
	return []byte(template.TextFormatterX(temp, alert))
}
