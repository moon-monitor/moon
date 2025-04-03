package rpc

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
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

type sendClient struct {
	server *bo.Server
}

// Send implements repository.Rabbit.
func (r *rabbitServer) Send() (repository.SendClient, bool) {
	server, ok := r.FirstRabbitConn()
	if !ok {
		return nil, false
	}
	return &sendClient{server: server}, true
}

// Sync implements repository.Rabbit.
func (r *rabbitServer) Sync() (repository.SyncClient, bool) {
	server, ok := r.FirstRabbitConn()
	if !ok {
		return nil, false
	}
	return &syncClient{server: server}, true
}

func (s *sendClient) Email(ctx context.Context, in *rabbitv1.SendEmailRequest) (*common.EmptyReply, error) {
	switch s.server.Config.Server.GetNetwork() {
	case config.Network_GRPC:
		return rabbitv1.NewSendClient(s.server.Conn).Email(ctx, in)
	case config.Network_HTTP:
		return rabbitv1.NewSendHTTPClient(s.server.Client).Email(ctx, in)
	default:
		return nil, merr.ErrorInternalServerError("network is not supported")
	}
}

func (s *sendClient) Sms(ctx context.Context, in *rabbitv1.SendSmsRequest) (*common.EmptyReply, error) {
	switch s.server.Config.Server.GetNetwork() {
	case config.Network_GRPC:
		return rabbitv1.NewSendClient(s.server.Conn).Sms(ctx, in)
	case config.Network_HTTP:
		return rabbitv1.NewSendHTTPClient(s.server.Client).Sms(ctx, in)
	default:
		return nil, merr.ErrorInternalServerError("network is not supported")
	}
}

func (s *sendClient) Hook(ctx context.Context, in *rabbitv1.SendHookRequest) (*common.EmptyReply, error) {
	switch s.server.Config.Server.GetNetwork() {
	case config.Network_GRPC:
		return rabbitv1.NewSendClient(s.server.Conn).Hook(ctx, in)
	case config.Network_HTTP:
		return rabbitv1.NewSendHTTPClient(s.server.Client).Hook(ctx, in)
	default:
		return nil, merr.ErrorInternalServerError("network is not supported")
	}
}

type syncClient struct {
	server *bo.Server
}

func (s *syncClient) Sms(ctx context.Context, in *rabbitv1.SyncSmsRequest) (*common.EmptyReply, error) {
	switch s.server.Config.Server.GetNetwork() {
	case config.Network_GRPC:
		return rabbitv1.NewSyncClient(s.server.Conn).Sms(ctx, in)
	case config.Network_HTTP:
		return rabbitv1.NewSyncHTTPClient(s.server.Client).Sms(ctx, in)
	default:
		return nil, merr.ErrorInternalServerError("network is not supported")
	}
}

// Email implements repository.SyncClient.
func (s *syncClient) Email(ctx context.Context, in *rabbitv1.SyncEmailRequest) (*common.EmptyReply, error) {
	switch s.server.Config.Server.GetNetwork() {
	case config.Network_GRPC:
		return rabbitv1.NewSyncClient(s.server.Conn).Email(ctx, in)
	case config.Network_HTTP:
		return rabbitv1.NewSyncHTTPClient(s.server.Client).Email(ctx, in)
	default:
		return nil, merr.ErrorInternalServerError("network is not supported")
	}
}
