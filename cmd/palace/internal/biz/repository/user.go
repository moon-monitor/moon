package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type SendEmailFunc func(ctx context.Context, params *bo.SendEmailParams) error

type User interface {
	FindByID(ctx context.Context, userID uint32) (*system.User, error)
	FindByEmail(ctx context.Context, email crypto.String) (*system.User, error)
	CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser, sendEmailFunc SendEmailFunc) (*system.User, error)
	SetEmail(ctx context.Context, user *system.User, sendEmailFunc SendEmailFunc) (*system.User, error)
	Create(ctx context.Context, user *system.User, sendEmailFunc SendEmailFunc) (*system.User, error)
	UpdateSelfInfo(ctx context.Context, user *system.User) error
	UpdatePassword(ctx context.Context, updateUserPasswordInfo *bo.UpdateUserPasswordInfo) error
	GetTeamsByUserID(ctx context.Context, userID uint32) ([]*system.Team, error)
	GetMemberByUserIDAndTeamID(ctx context.Context, userID, teamID uint32) (*system.TeamMember, error)
	GetAllTeamMembers(ctx context.Context, userID uint32) ([]*system.TeamMember, error)
	GetTeamsByIDs(ctx context.Context, teamIDs []uint32) ([]*system.Team, error)
}
