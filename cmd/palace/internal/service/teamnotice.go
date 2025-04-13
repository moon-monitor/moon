package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamNoticeService struct {
	palacev1.UnimplementedTeamNoticeServer
}

func NewTeamNoticeService() *TeamNoticeService {
	return &TeamNoticeService{}
}

func (s *TeamNoticeService) SaveTeamNoticeHook(ctx context.Context, req *palacev1.SaveTeamNoticeHookRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) UpdateTeamNoticeHookStatus(ctx context.Context, req *palacev1.UpdateTeamNoticeHookStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) DeleteTeamNoticeHook(ctx context.Context, req *palacev1.DeleteTeamNoticeHookRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) GetTeamNoticeHook(ctx context.Context, req *palacev1.GetTeamNoticeHookRequest) (*palacev1.GetTeamNoticeHookReply, error) {
	return &palacev1.GetTeamNoticeHookReply{}, nil
}
func (s *TeamNoticeService) ListTeamNoticeHook(ctx context.Context, req *palacev1.ListTeamNoticeHookRequest) (*palacev1.ListTeamNoticeHookReply, error) {
	return &palacev1.ListTeamNoticeHookReply{}, nil
}
func (s *TeamNoticeService) SaveTeamNoticeGroup(ctx context.Context, req *palacev1.SaveTeamNoticeGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) UpdateTeamNoticeGroupStatus(ctx context.Context, req *palacev1.UpdateTeamNoticeGroupStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) DeleteTeamNoticeGroup(ctx context.Context, req *palacev1.DeleteTeamNoticeGroupRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamNoticeService) GetTeamNoticeGroup(ctx context.Context, req *palacev1.GetTeamNoticeGroupRequest) (*palacev1.GetTeamNoticeGroupReply, error) {
	return &palacev1.GetTeamNoticeGroupReply{}, nil
}
func (s *TeamNoticeService) ListTeamNoticeGroup(ctx context.Context, req *palacev1.ListTeamNoticeGroupRequest) (*palacev1.ListTeamNoticeGroupReply, error) {
	return &palacev1.ListTeamNoticeGroupReply{}, nil
}
