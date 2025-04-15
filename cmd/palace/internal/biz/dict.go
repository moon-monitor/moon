package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewDict(
	teamDictRepo repository.TeamDict,
	logger log.Logger,
) *Dict {
	return &Dict{
		helper:       log.NewHelper(log.With(logger, "module", "biz.dict")),
		teamDictRepo: teamDictRepo,
	}
}

type Dict struct {
	helper *log.Helper

	teamDictRepo repository.TeamDict
}

func (d *Dict) getTeamID(ctx context.Context) (uint32, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return 0, merr.ErrorPermissionDenied("team id is invalid")
	}
	return teamID, nil
}

func (d *Dict) SaveDict(ctx context.Context, req *bo.SaveDictReq) error {
	teamID, err := d.getTeamID(ctx)
	if err != nil {
		return err
	}
	if req.DictID == 0 {
		return d.teamDictRepo.Create(ctx, teamID, req)
	}
	dictItem, err := d.teamDictRepo.Get(ctx, teamID, req.DictID)
	if err != nil {
		return err
	}
	return d.teamDictRepo.Update(ctx, teamID, req.WithUpdateParams(dictItem))
}

func (d *Dict) GetDict(ctx context.Context, req *bo.OperateOneDictReq) (do.Dict, error) {
	teamID, err := d.getTeamID(ctx)
	if err != nil {
		return nil, err
	}
	return d.teamDictRepo.Get(ctx, teamID, req.DictID)
}

func (d *Dict) UpdateDictStatus(ctx context.Context, req *bo.UpdateDictStatusReq) error {
	teamID, err := d.getTeamID(ctx)
	if err != nil {
		return err
	}
	return d.teamDictRepo.UpdateStatus(ctx, teamID, req)
}

func (d *Dict) DeleteDict(ctx context.Context, req *bo.OperateOneDictReq) error {
	teamID, err := d.getTeamID(ctx)
	if err != nil {
		return err
	}
	return d.teamDictRepo.Delete(ctx, teamID, req.DictID)
}

func (d *Dict) ListDict(ctx context.Context, req *bo.ListDictReq) (*bo.ListDictReply, error) {
	teamID, err := d.getTeamID(ctx)
	if err != nil {
		return nil, err
	}
	return d.teamDictRepo.List(ctx, teamID, req)
}
