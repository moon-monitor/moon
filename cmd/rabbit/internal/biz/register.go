package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/conf"
	"github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/config"
)

func NewRegisterBiz(bc *conf.Bootstrap, serverRegisterRepo repository.ServerRegister, logger log.Logger) *RegisterBiz {
	return &RegisterBiz{
		serverRegisterRepo: serverRegisterRepo,
		bc:                 bc,
		uuid:               uuid.New().String(),
		helper:             log.NewHelper(log.With(logger, "module", "biz.register")),
	}
}

type RegisterBiz struct {
	uuid               string
	bc                 *conf.Bootstrap
	serverRegisterRepo repository.ServerRegister
	helper             *log.Helper
}

func (b *RegisterBiz) register() *common.ServerRegisterRequest {
	params := &common.ServerRegisterRequest{
		Server:    b.bc.GetMicroServer(),
		Discovery: nil,
		TeamIds:   b.bc.GetTeamIds(),
		IsOnline:  true,
		Uuid:      b.uuid,
	}
	register := b.bc.GetRegistry()
	if register != nil {
		params.Discovery = &config.Discovery{
			Driver: register.GetDriver(),
			Enable: register.GetEnable(),
			Etcd:   register.GetEtcd(),
		}
	}
	return params
}

func (b *RegisterBiz) Online(ctx context.Context) error {
	return b.serverRegisterRepo.Register(ctx, b.register())
}

func (b *RegisterBiz) Offline(ctx context.Context) error {
	return b.serverRegisterRepo.Register(ctx, b.register())
}
