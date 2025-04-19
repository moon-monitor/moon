package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

func NewSystem(
	roleRepo repository.Role,
	transactionRepo repository.Transaction,
	logger log.Logger,
) *System {
	return &System{
		roleRepo:        roleRepo,
		transactionRepo: transactionRepo,
		helper:          log.NewHelper(log.With(logger, "module", "biz.system")),
	}
}

type System struct {
	roleRepo        repository.Role
	transactionRepo repository.Transaction
	helper          *log.Helper
}

func (s *System) GetRole(ctx context.Context, roleId uint32) (do.Role, error) {
	return s.roleRepo.Get(ctx, roleId)
}

func (s *System) GetRoles(ctx context.Context, req *bo.ListRoleReq) (*bo.ListRoleReply, error) {
	return s.roleRepo.List(ctx, req)
}

func (s *System) SaveRole(ctx context.Context, req *bo.SaveRoleReq) error {
	return s.transactionRepo.BizExec(ctx, func(ctx context.Context) error {
		if req.GetID() <= 0 {
			return s.roleRepo.Create(ctx, req)
		}
		roleDo, err := s.roleRepo.Get(ctx, req.GetID())
		if err != nil {
			return err
		}
		req.WithRole(roleDo)
		return s.roleRepo.Update(ctx, req)
	})
}

func (s *System) UpdateRoleStatus(ctx context.Context, req *bo.UpdateRoleStatusReq) error {
	return s.roleRepo.UpdateStatus(ctx, req)
}
