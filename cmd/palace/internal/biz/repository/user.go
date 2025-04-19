package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type User interface {
	FindByID(ctx context.Context, userID uint32) (do.User, error)
	FindByEmail(ctx context.Context, email crypto.String) (do.User, error)
	CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser, sendEmailFunc bo.SendEmailFun) (do.User, error)
	SetEmail(ctx context.Context, user do.User, sendEmailFunc bo.SendEmailFun) (do.User, error)
	Create(ctx context.Context, user do.User, sendEmailFunc bo.SendEmailFun) (do.User, error)
	UpdateUserInfo(ctx context.Context, user do.User) error
	UpdatePassword(ctx context.Context, updateUserPasswordInfo *bo.UpdateUserPasswordInfo) error
	GetTeamsByUserID(ctx context.Context, userID uint32) ([]do.Team, error)
	AppendTeam(ctx context.Context, team do.Team) error
	UpdateUserStatus(ctx context.Context, req *bo.UpdateUserStatusRequest) error
	UpdateUserPosition(ctx context.Context, req *bo.UpdateUserPositionRequest) error
	List(ctx context.Context, req *bo.UserListRequest) (*bo.UserListReply, error)
}
