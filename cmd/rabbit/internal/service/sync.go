package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	apiv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

type SyncService struct {
	apiv1.UnimplementedSyncServer
	configBiz *biz.Config
	helper    *log.Helper
}

func NewSyncService(configBiz *biz.Config, logger log.Logger) *SyncService {
	return &SyncService{
		configBiz: configBiz,
		helper:    log.NewHelper(log.With(logger, "module", "service.sync")),
	}
}
func (s *SyncService) Sms(ctx context.Context, req *apiv1.SyncSmsRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *SyncService) Email(ctx context.Context, req *apiv1.SyncEmailRequest) (*common.EmptyReply, error) {
	emails := slices.Map(req.GetEmails(), func(emailItem *common.EmailConfig) bo.EmailConfig {
		return emailItem
	})
	if err := s.configBiz.SetEmailConfig(ctx, emails...); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}
