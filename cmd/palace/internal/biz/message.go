package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
	"github.com/moon-monitor/moon/pkg/plugin/email"
)

func NewMessage(bc *conf.Bootstrap, rabbitRepo repository.Rabbit, logger log.Logger) *Message {
	emailConfig := bc.GetEmail()
	return &Message{
		rabbitRepo: rabbitRepo,
		emailConfig: &common.EmailConfig{
			User:   emailConfig.GetUser(),
			Pass:   emailConfig.GetPass(),
			Host:   emailConfig.GetHost(),
			Port:   emailConfig.GetPort(),
			Enable: true,
			Name:   emailConfig.GetName(),
		},
		helper: log.NewHelper(log.With(logger, "module", "biz.message")),
	}
}

type Message struct {
	rabbitRepo  repository.Rabbit
	helper      *log.Helper
	emailConfig *common.EmailConfig
}

func (a *Message) SendEmail(ctx context.Context, sendEmailParams *bo.SendEmailParams) error {
	sendClient, ok := a.rabbitRepo.Send()
	if !ok {
		// call local send email
		return a.localSendEmail(ctx, sendEmailParams)
	}
	// call rabbit server send email
	return a.rabbitSendEmail(ctx, sendClient, sendEmailParams)
}

func (a *Message) localSendEmail(_ context.Context, params *bo.SendEmailParams) error {
	emailInstance := email.New(a.emailConfig)
	emailInstance.SetTo(params.Email).
		SetSubject(params.Subject).
		SetBody(params.Body, params.ContentType)
	if err := emailInstance.Send(); err != nil {
		a.helper.Warnw("method", "local send email error", "params", params, "error", err)
		return err
	}
	return nil
}

func (a *Message) rabbitSendEmail(ctx context.Context, client repository.SendClient, params *bo.SendEmailParams) error {
	reply, err := client.Email(ctx, &rabbitv1.SendEmailRequest{
		Emails:      []string{params.Email},
		Body:        params.Body,
		Subject:     params.Subject,
		ContentType: params.ContentType,
		EmailConfig: a.emailConfig,
	})
	if err != nil {
		a.helper.Warnw("method", "rabbit send email error", "params", params, "error", err)
		return err
	}
	a.helper.Debugf("send email reply: %v", reply)
	return nil
}
