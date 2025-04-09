package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/pointer"
	"github.com/moon-monitor/moon/pkg/util/slices"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/service/build"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SendService struct {
	apiv1.UnimplementedSendServer
	configBiz *biz.Config
	emailBiz  *biz.Email
	smsBiz    *biz.SMS
	hookBiz   *biz.Hook

	helper *log.Helper
}

func NewSendService(
	configBiz *biz.Config,
	emailBiz *biz.Email,
	smsBiz *biz.SMS,
	hookBiz *biz.Hook,
	logger log.Logger,
) *SendService {
	return &SendService{
		configBiz: configBiz,
		emailBiz:  emailBiz,
		smsBiz:    smsBiz,
		hookBiz:   hookBiz,
		helper:    log.NewHelper(log.With(logger, "module", "service.send")),
	}
}

func (s *SendService) Email(ctx context.Context, req *apiv1.SendEmailRequest) (*common.EmptyReply, error) {
	emailConfig := s.configBiz.GetEmailConfig(ctx, req.ConfigName, req.GetEmailConfig())
	opts := []bo.SendEmailParamsOption{
		bo.WithSendEmailParamsOptionEmail(req.GetEmails()...),
		bo.WithSendEmailParamsOptionBody(req.GetBody()),
		bo.WithSendEmailParamsOptionSubject(req.GetSubject()),
		bo.WithSendEmailParamsOptionContentType(req.GetContentType()),
		bo.WithSendEmailParamsOptionAttachment(req.GetAttachment()),
		bo.WithSendEmailParamsOptionCc(req.GetCc()...),
	}
	sendEmailParams, err := bo.NewSendEmailParams(emailConfig, opts...)
	if err != nil {
		return nil, err
	}
	if err := s.emailBiz.Send(ctx, sendEmailParams); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

func (s *SendService) Sms(ctx context.Context, req *apiv1.SendSmsRequest) (*common.EmptyReply, error) {
	smsConfig, _ := build.ToSMSConfig(req.GetSmsConfig())
	smsConfig = s.configBiz.GetSMSConfig(ctx, req.ConfigName, smsConfig)
	opts := []bo.SendSMSParamsOption{
		bo.WithSendSMSParamsOptionPhoneNumbers(req.GetPhones()...),
		bo.WithSendSMSParamsOptionTemplateParam(req.GetTemplateParameters()),
		bo.WithSendSMSParamsOptionTemplateCode(req.GetTemplateCode()),
	}
	sendSMSParams, err := bo.NewSendSMSParams(smsConfig, opts...)
	if err != nil {
		return nil, err
	}
	if err := s.smsBiz.Send(ctx, sendSMSParams); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

func (s *SendService) Hook(ctx context.Context, req *apiv1.SendHookRequest) (*common.EmptyReply, error) {
	hookConfigs := slices.MapFilter(req.GetHooks(), func(hookItem *common.HookConfig) (bo.HookConfig, bool) {
		opts := []do.HookConfigOption{
			do.WithHookConfigOptionApp(hookItem.App),
			do.WithHookConfigOptionEnable(hookItem.Enable),
			do.WithHookConfigOptionHeaders(hookItem.Headers),
			do.WithHookConfigOptionName(hookItem.Name),
			do.WithHookConfigOptionPassword(hookItem.Password),
			do.WithHookConfigOptionSecret(hookItem.Secret),
			do.WithHookConfigOptionToken(hookItem.Token),
			do.WithHookConfigOptionUsername(hookItem.Username),
		}
		var hookConfig bo.HookConfig
		hookConfigDo, err := do.NewHookConfig(hookItem.Url, opts...)
		if err == nil {
			hookConfig = hookConfigDo
		}
		return s.configBiz.GetHookConfig(ctx, pointer.Of(hookItem.Name), hookConfig), true
	})
	opts := []bo.SendHookParamsOption{
		bo.WithSendHookParamsOptionBody([]byte(req.GetBody())),
	}
	sendHookParams, err := bo.NewSendHookParams(hookConfigs, opts...)
	if err != nil {
		return nil, err
	}
	if err := s.hookBiz.Send(ctx, sendHookParams); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}
