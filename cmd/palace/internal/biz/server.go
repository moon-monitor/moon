package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewServerBiz(serverRepo repository.Server, logger log.Logger) *ServerBiz {
	return &ServerBiz{
		serverRepo: serverRepo,
		helper:     log.NewHelper(log.With(logger, "module", "biz.server")),
	}
}

type ServerBiz struct {
	serverRepo repository.Server
	helper     *log.Helper
}

func (b *ServerBiz) Register(ctx context.Context, req *bo.ServerRegisterReq) error {
	if req == nil {
		return merr.ErrorInvalidArgument("invalid request")
	}

	var err error
	// Store connection based on server type
	switch req.ServerType {
	case vobj.ServerTypeRabbit:
		err = b.serverRepo.RegisterRabbit(ctx, req)
	case vobj.ServerTypeHouyi:
		err = b.serverRepo.RegisterHouyi(ctx, req)
	default:
		return merr.ErrorInvalidArgument("unsupported server type: %v", req.ServerType)
	}

	if err != nil {
		return merr.ErrorInternalServerError("failed to register server: %v", err)
	}

	b.helper.Debugf("registered server type: %v, uuid: %s", req.ServerType, req.Uuid)
	return nil
}
