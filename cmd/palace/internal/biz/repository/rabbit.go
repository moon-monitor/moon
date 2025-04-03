package repository

import (
	"context"

	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type Rabbit interface {
	Send() (SendClient, bool)
	Sync() (SyncClient, bool)
}

type SendClient interface {
	Email(ctx context.Context, in *rabbitv1.SendEmailRequest) (*common.EmptyReply, error)
	Sms(ctx context.Context, in *rabbitv1.SendSmsRequest) (*common.EmptyReply, error)
	Hook(ctx context.Context, in *rabbitv1.SendHookRequest) (*common.EmptyReply, error)
}

type SyncClient interface {
	Sms(ctx context.Context, in *rabbitv1.SyncSmsRequest) (*common.EmptyReply, error)
	Email(ctx context.Context, in *rabbitv1.SyncEmailRequest) (*common.EmptyReply, error)
}
