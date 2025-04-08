package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/plugin/sms"
	"github.com/moon-monitor/moon/pkg/plugin/sms/ali"
)

func NewSendRepo(d *data.Data, logger log.Logger) repository.Send {
	return &sendImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.send")),
		logger: logger,
	}
}

type sendImpl struct {
	*data.Data
	helper *log.Helper
	logger log.Logger
}

func (s *sendImpl) Email(_ context.Context, params bo.SendEmailParams) error {
	emailInstance, ok := s.GetEmail(params.GetConfig().GetName())
	if !ok {
		emailInstance = email.New(params.GetConfig())
		s.SetEmail(params.GetConfig().GetName(), emailInstance)
	}

	emailInstance.SetTo(params.GetEmails()...).
		SetSubject(params.GetSubject()).
		SetBody(params.GetBody())
	if params.GetAttachment() != "" {
		emailInstance.SetAttach(params.GetAttachment())
	}
	if len(params.GetCc()) > 0 {
		emailInstance.SetCc(params.GetCc()...)
	}
	return emailInstance.Send()
}

func (s *sendImpl) SMS(ctx context.Context, params bo.SendSMSParams) error {
	var err error
	smsInstance, ok := s.GetSms(params.GetConfig().GetName())
	if !ok {
		smsInstance, err = s.newSms(params.GetConfig())
		if err != nil {
			return err
		}
		s.SetSms(params.GetConfig().GetName(), smsInstance)
	}
	message := sms.Message{
		TemplateCode:  params.GetTemplateCode(),
		TemplateParam: params.GetGetTemplateParam(),
	}
	if len(params.GetPhoneNumbers()) == 0 {
		return smsInstance.Send(ctx, params.GetPhoneNumbers()[0], message)
	}
	return smsInstance.SendBatch(ctx, params.GetPhoneNumbers(), message)
}

func (s *sendImpl) newSms(config bo.SMSConfig) (sms.Sender, error) {
	switch config.GetType() {
	case common.SMSConfig_ALIYUN:
		return ali.NewAliyun(config, ali.WithAliyunLogger(s.logger))
	default:
		return nil, merr.ErrorParamsError("No SMS configuration is available")
	}
}
