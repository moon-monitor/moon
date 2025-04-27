package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

type SendMessageLog interface {
	Get(ctx context.Context, params *bo.GetSendMessageLogParams) (do.SendMessageLog, error)
	Create(ctx context.Context, params *bo.CreateSendMessageLogParams) error
	UpdateStatus(ctx context.Context, params *bo.UpdateSendMessageLogStatusParams) error
	List(ctx context.Context, params *bo.ListSendMessageLogParams) (*bo.ListSendMessageLogReply, error)
}
