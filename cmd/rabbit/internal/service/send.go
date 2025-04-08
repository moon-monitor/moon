package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type SendService struct {
	apiv1.UnimplementedSendServer
	configBiz *biz.Config
	emailBiz  *biz.Email

	helper *log.Helper
}

func NewSendService(configBiz *biz.Config, emailBiz *biz.Email, logger log.Logger) *SendService {
	return &SendService{
		configBiz: configBiz,
		emailBiz:  emailBiz,
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
	return &common.EmptyReply{}, nil
}

func (s *SendService) Hook(ctx context.Context, req *apiv1.SendHookRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
