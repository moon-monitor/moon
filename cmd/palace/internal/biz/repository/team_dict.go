package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type TeamDict interface {
	Get(ctx context.Context, teamID uint32, dictID uint32) (bo.Dict, error)
	Delete(ctx context.Context, teamID, dictID uint32) error
	Create(ctx context.Context, teamID uint32, dict bo.Dict) error
	Update(ctx context.Context, teamID uint32, dict bo.Dict) error
	UpdateStatus(ctx context.Context, teamID uint32, dictIds []uint32, status vobj.GlobalStatus) error
	List(ctx context.Context, teamID uint32, req *bo.ListDictReq) (*bo.ListDictReply, error)
}
