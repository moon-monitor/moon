package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type SendEmailFunc func(ctx context.Context, params *bo.SendEmailParams) error

type User interface {
	FindByID(ctx context.Context, userID uint32) (do.User, error)
	FindByEmail(ctx context.Context, email crypto.String) (do.User, error)
	CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser, sendEmailFunc SendEmailFunc) (do.User, error)
	SetEmail(ctx context.Context, user do.User, sendEmailFunc SendEmailFunc) (do.User, error)
	Create(ctx context.Context, user do.User, sendEmailFunc SendEmailFunc) (do.User, error)
	UpdateSelfInfo(ctx context.Context, user do.User) error
	UpdatePassword(ctx context.Context, updateUserPasswordInfo *bo.UpdateUserPasswordInfo) error
	GetTeamsByUserID(ctx context.Context, userID uint32) ([]do.Team, error)
	GetMemberByUserIDAndTeamID(ctx context.Context, userID, teamID uint32) (do.TeamMember, error)
	AppendTeam(ctx context.Context, team do.Team) error
}
