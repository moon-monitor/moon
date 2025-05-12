package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/service/build"
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
	smss := build.ToSMSConfigs(req.GetSmss())
	if err := s.configBiz.SetSMSConfig(ctx, smss...); err != nil {
		return nil, err
	}
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

func (s *SyncService) Hook(ctx context.Context, req *apiv1.SyncHookRequest) (*common.EmptyReply, error) {
	hooks := slices.Map(req.GetHooks(), func(hookItem *common.HookConfig) bo.HookConfig {
		return hookItem
	})
	if err := s.configBiz.SetHookConfig(ctx, hooks...); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

func (s *SyncService) NoticeGroup(ctx context.Context, req *apiv1.SyncNoticeGroupRequest) (*common.EmptyReply, error) {
	noticeGroups := slices.Map(req.GetNoticeGroups(), func(noticeGroupItem *common.NoticeGroup) bo.NoticeGroup {
		templates := slices.Map(noticeGroupItem.GetTemplates(), func(templateItem *common.NoticeGroup_Template) bo.Template {
			return templateItem
		})
		return bo.NewNoticeGroup(
			bo.WithNoticeGroupOptionName(noticeGroupItem.GetName()),
			bo.WithNoticeGroupOptionSmsConfigName(noticeGroupItem.GetSmsConfigName()),
			bo.WithNoticeGroupOptionEmailConfigName(noticeGroupItem.GetEmailConfigName()),
			bo.WithNoticeGroupOptionHookConfigNames(noticeGroupItem.GetHookConfigNames()),
			bo.WithNoticeGroupOptionSmsUserNames(noticeGroupItem.GetSmsUserNames()),
			bo.WithNoticeGroupOptionEmailUserNames(noticeGroupItem.GetEmailUserNames()),
			bo.WithNoticeGroupOptionTemplates(templates),
		)
	})
	if err := s.configBiz.SetNoticeGroupConfig(ctx, noticeGroups...); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}
