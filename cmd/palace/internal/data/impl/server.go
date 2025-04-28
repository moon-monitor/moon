package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

func NewServerRepo(data *data.Data, logger log.Logger) repository.Server {
	return &serverRepository{
		data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.server")),
	}
}

type serverRepository struct {
	data   *data.Data
	helper *log.Helper
}

func (s *serverRepository) DeregisterRabbit(ctx context.Context, req *bo.ServerRegisterReq) error {
	s.helper.WithContext(ctx).Debugf("deregister rabbit server: %v", req)
	rabbitConn, ok := s.data.GetRabbitConn(req.Uuid)
	if !ok {
		return nil
	}
	defer s.data.RemoveRabbitConn(req.Uuid)
	if err := rabbitConn.Close(); err != nil {
		return err
	}
	return nil
}

func (s *serverRepository) DeregisterHouyi(ctx context.Context, req *bo.ServerRegisterReq) error {
	s.helper.WithContext(ctx).Debugf("deregister houyi server: %v", req)
	houyiConn, ok := s.data.GetHouyiConn(req.Uuid)
	if !ok {
		return nil
	}
	defer s.data.RemoveHouyiConn(req.Uuid)
	if err := houyiConn.Close(); err != nil {
		return err
	}
	return nil
}

func (s *serverRepository) RegisterRabbit(ctx context.Context, req *bo.ServerRegisterReq) error {
	s.helper.WithContext(ctx).Debugf("register rabbit server: %v", req)
	initConfig := &server.InitConfig{
		MicroConfig: req.Server,
		Registry:    (*config.Registry)(req.Discovery),
	}
	serverBo := &bo.Server{Config: req}
	switch req.Server.GetNetwork() {
	case config.Network_GRPC:
		conn, err := server.InitGRPCClient(initConfig)
		if err != nil {
			return err
		}
		serverBo.Conn = conn
		s.data.SetRabbitConn(req.Uuid, serverBo)
	case config.Network_HTTP:
		client, err := server.InitHTTPClient(initConfig)
		if err != nil {
			return err
		}
		serverBo.Client = client
		s.data.SetRabbitConn(req.Uuid, serverBo)
	}
	return nil
}

func (s *serverRepository) RegisterHouyi(ctx context.Context, req *bo.ServerRegisterReq) error {
	s.helper.WithContext(ctx).Debugf("register houyi server: %v", req)
	initConfig := &server.InitConfig{
		MicroConfig: req.Server,
		Registry:    (*config.Registry)(req.Discovery),
	}
	serverBo := &bo.Server{Config: req}
	switch req.Server.GetNetwork() {
	case config.Network_GRPC:
		conn, err := server.InitGRPCClient(initConfig)
		if err != nil {
			return err
		}
		serverBo.Conn = conn
		s.data.SetHouyiConn(req.Uuid, serverBo)
	case config.Network_HTTP:
		client, err := server.InitHTTPClient(initConfig)
		if err != nil {
			return err
		}
		serverBo.Client = client
		s.data.SetHouyiConn(req.Uuid, serverBo)
	}
	return nil
}
