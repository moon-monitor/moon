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

	// 获取用户所属的所有团队
	GetTeamsByUserID(ctx context.Context, userID uint32) ([]*system.Team, error)
	// 获取用户在指定团队中的成员信息
	GetMemberByUserIDAndTeamID(ctx context.Context, userID, teamID uint32) (*system.TeamMember, error)
	// 获取用户所有团队的成员信息
	GetAllTeamMembers(ctx context.Context, userID uint32) ([]*system.TeamMember, error)
	// 根据团队ID列表获取团队
	GetTeamsByIDs(ctx context.Context, teamIDs []uint32) ([]*system.Team, error)
}
