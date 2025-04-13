package service

import (
	"context"

	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamDictService struct {
	palacev1.UnimplementedTeamDictServer
}

func NewTeamDictService() *TeamDictService {
	return &TeamDictService{}
}

func (s *TeamDictService) SaveTeamDict(ctx context.Context, req *palacev1.SaveTeamDictRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDictService) UpdateTeamDictStatus(ctx context.Context, req *palacev1.UpdateTeamDictStatusRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDictService) DeleteTeamDict(ctx context.Context, req *palacev1.DeleteTeamDictRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
func (s *TeamDictService) GetTeamDict(ctx context.Context, req *palacev1.GetTeamDictRequest) (*palacev1.GetTeamDictReply, error) {
	return &palacev1.GetTeamDictReply{}, nil
}
func (s *TeamDictService) ListTeamDict(ctx context.Context, req *palacev1.ListTeamDictRequest) (*palacev1.ListTeamDictReply, error) {
	return &palacev1.ListTeamDictReply{}, nil
}
