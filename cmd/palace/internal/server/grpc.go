package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bc *conf.Bootstrap) *grpc.Server {
	c := bc.GetGrpc()
	opts := []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
		),
	}
	if c.GetNetwork() != "" {
		opts = append(opts, grpc.Network(c.GetNetwork()))
	}
	if c.GetAddr() != "" {
		opts = append(opts, grpc.Address(c.GetAddr()))
	}
	if c.GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	return srv
}
