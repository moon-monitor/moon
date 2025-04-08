package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
	"github.com/moon-monitor/moon/pkg/plugin/email"
)

func NewSendRepo(d *data.Data, logger log.Logger) repository.Send {
	return &sendImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.send")),
	}
}

type sendImpl struct {
	*data.Data
	helper *log.Helper
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
