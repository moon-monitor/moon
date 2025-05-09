package repository

import (
	"context"

	houyiv1 "github.com/moon-monitor/moon/pkg/api/houyi/v1"
)

type Houyi interface {
	Sync() (HouyiSyncClient, bool)
}

type HouyiSyncClient interface {
	MetricMetadata(ctx context.Context, req *houyiv1.MetricMetadataRequest) (*houyiv1.SyncReply, error)
}
