package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamNoticeService struct {
	palace.UnimplementedTeamNoticeServer
	teamHookBiz *biz.TeamHook
}

func NewTeamNoticeService(teamHookBiz *biz.TeamHook) *TeamNoticeService {
	return &TeamNoticeService{
		teamHookBiz: teamHookBiz,
	}
}

// SaveTeamNoticeHook 保存团队通知钩子
func (s *TeamNoticeService) SaveTeamNoticeHook(ctx context.Context, req *palace.SaveTeamNoticeHookRequest) (*common.EmptyReply, error) {
	if err := s.teamHookBiz.SaveHook(ctx, build.ToSaveTeamNoticeHookRequest(req)); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

// UpdateTeamNoticeHookStatus 更新钩子状态
func (s *TeamNoticeService) UpdateTeamNoticeHookStatus(ctx context.Context, req *palace.UpdateTeamNoticeHookStatusRequest) (*common.EmptyReply, error) {
	params := &bo.UpdateTeamNoticeHookStatusRequest{
		HookID: req.GetHookID(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}
	if err := s.teamHookBiz.UpdateHookStatus(ctx, params); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

// DeleteTeamNoticeHook 删除钩子
func (s *TeamNoticeService) DeleteTeamNoticeHook(ctx context.Context, req *palace.DeleteTeamNoticeHookRequest) (*common.EmptyReply, error) {
	if err := s.teamHookBiz.DeleteHook(ctx, req.GetHookID()); err != nil {
		return nil, err
	}
	return &common.EmptyReply{}, nil
}

// GetTeamNoticeHook 获取钩子详情
func (s *TeamNoticeService) GetTeamNoticeHook(ctx context.Context, req *palace.GetTeamNoticeHookRequest) (*palace.GetTeamNoticeHookReply, error) {
	hook, err := s.teamHookBiz.GetHook(ctx, req.GetHookID())
	if err != nil {
		return nil, err
	}
	return &palace.GetTeamNoticeHookReply{
		Detail: build.ToNoticeHookItem(hook),
	}, nil
}

// ListTeamNoticeHook 获取钩子列表
func (s *TeamNoticeService) ListTeamNoticeHook(ctx context.Context, req *palace.ListTeamNoticeHookRequest) (*palace.ListTeamNoticeHookReply, error) {
	reply, err := s.teamHookBiz.ListHook(ctx, build.ToListTeamNoticeHookRequest(req))
	if err != nil {
		return nil, err
	}

	return &palace.ListTeamNoticeHookReply{
		Items:      build.ToNoticeHookItems(reply.Items),
		Pagination: build.ToPaginationReply(reply.PaginationReply),
	}, nil
}

func (s *TeamNoticeService) SaveTeamNoticeGroup(ctx context.Context, req *palace.SaveTeamNoticeGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamNoticeService) UpdateTeamNoticeGroupStatus(ctx context.Context, req *palace.UpdateTeamNoticeGroupStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamNoticeService) DeleteTeamNoticeGroup(ctx context.Context, req *palace.DeleteTeamNoticeGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}

func (s *TeamNoticeService) GetTeamNoticeGroup(ctx context.Context, req *palace.GetTeamNoticeGroupRequest) (*palace.GetTeamNoticeGroupReply, error) {
	return &palace.GetTeamNoticeGroupReply{}, nil
}

func (s *TeamNoticeService) ListTeamNoticeGroup(ctx context.Context, req *palace.ListTeamNoticeGroupRequest) (*palace.ListTeamNoticeGroupReply, error) {
	return &palace.ListTeamNoticeGroupReply{}, nil
}
