package rpc

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

func NewRabbitServer(data *data.Data, logger log.Logger) repository.Rabbit {
	return &rabbitServer{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.rabbit")),
	}
}

type rabbitServer struct {
	*data.Data
	helper *log.Helper
}

// Send implements repository.Rabbit.
func (r *rabbitServer) Send() (rabbitv1.SendClient, bool) {
	conn, ok := r.FirstRabbitConn()
	if !ok {
		return nil, false
	}
	return rabbitv1.NewSendClient(conn), true
}

// Sync implements repository.Rabbit.
func (r *rabbitServer) Sync() (rabbitv1.SyncClient, bool) {
	conn, ok := r.FirstRabbitConn()
	if !ok {
		return nil, false
	}
	return rabbitv1.NewSyncClient(conn), true
}
