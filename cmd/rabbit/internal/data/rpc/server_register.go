package rpc

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
	common "github.com/moon-monitor/moon/pkg/api/common"
)

func NewServerRegisterRepo(data *data.Data, logger log.Logger) repository.ServerRegister {
	return &serverRegisterRepo{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.server_register")),
	}
}

type serverRegisterRepo struct {
	*data.Data

	helper *log.Helper
}

func (r *serverRegisterRepo) Register(ctx context.Context, server *common.ServerRegisterRequest) error {
	return nil
}
