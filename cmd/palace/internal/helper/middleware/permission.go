package middleware

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type userIDContextKey struct{}

// GetUserIDByContext Retrieve the user ID from the context.
func GetUserIDByContext(ctx context.Context) (uint32, bool) {
	userID, ok := ctx.Value(userIDContextKey{}).(uint32)
	return userID, ok
}

// WithUserIDContext Set the user ID in the context.
func WithUserIDContext(ctx context.Context, userID uint32) context.Context {
	return context.WithValue(ctx, userIDContextKey{}, userID)
}

type teamIDContextKey struct{}

// GetTeamIDByContext Retrieve the team ID from the context.
func GetTeamIDByContext(ctx context.Context) (uint32, bool) {
	teamID, ok := ctx.Value(teamIDContextKey{}).(uint32)
	return teamID, ok
}

// WithTeamIDContext Set the team ID in the context.
func WithTeamIDContext(ctx context.Context, teamID uint32) context.Context {
	return context.WithValue(ctx, teamIDContextKey{}, teamID)
}

type teamMemberIDContextKey struct{}

// GetTeamMemberIDByContext Retrieve the team member ID from the context.
func GetTeamMemberIDByContext(ctx context.Context) (uint32, bool) {
	teamMemberID, ok := ctx.Value(teamMemberIDContextKey{}).(uint32)
	return teamMemberID, ok
}

// WithTeamMemberIDContext Set the team member ID in the context.
func WithTeamMemberIDContext(ctx context.Context, teamMemberID uint32) context.Context {
	return context.WithValue(ctx, teamMemberIDContextKey{}, teamMemberID)
}

type sysPositionContextKey struct{}

// GetSysPositionByContext Retrieve the system position from the context.
func GetSysPositionByContext(ctx context.Context) (vobj.Role, bool) {
	sysPosition, ok := ctx.Value(sysPositionContextKey{}).(vobj.Role)
	return sysPosition, ok
}

// WithSysPositionContext Set the system position in the context.
func WithSysPositionContext(ctx context.Context, sysPosition vobj.Role) context.Context {
	return context.WithValue(ctx, sysPositionContextKey{}, sysPosition)
}

type sysRoleIDContextKey struct{}

// GetSysRoleIDByContext Retrieve the system role from the context.
func GetSysRoleIDByContext(ctx context.Context) (uint32, bool) {
	sysRole, ok := ctx.Value(sysRoleIDContextKey{}).(uint32)
	return sysRole, ok
}

// WithSysRoleIDContext Set the system role in the context.
func WithSysRoleIDContext(ctx context.Context, sysRole uint32) context.Context {
	return context.WithValue(ctx, sysRoleIDContextKey{}, sysRole)
}

type teamPositionContextKey struct{}

// GetTeamPositionByContext Retrieve the team position from the context.
func GetTeamPositionByContext(ctx context.Context) (vobj.Role, bool) {
	teamPosition, ok := ctx.Value(teamPositionContextKey{}).(vobj.Role)
	return teamPosition, ok
}

// WithTeamPositionContext Set the team position in the context.
func WithTeamPositionContext(ctx context.Context, teamPosition vobj.Role) context.Context {
	return context.WithValue(ctx, teamPositionContextKey{}, teamPosition)
}

type teamRoleIDContextKey struct{}

// GetTeamRoleIDByContext Retrieve the team role from the context.
func GetTeamRoleIDByContext(ctx context.Context) (uint32, bool) {
	teamRole, ok := ctx.Value(teamRoleIDContextKey{}).(uint32)
	return teamRole, ok
}

// WithTeamRoleIDContext Set the team role in the context.
func WithTeamRoleIDContext(ctx context.Context, teamRole uint32) context.Context {
	return context.WithValue(ctx, teamRoleIDContextKey{}, teamRole)
}

type tokenContextKey struct{}

// GetTokenByContext Retrieve the token from the context.
func GetTokenByContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenContextKey{}).(string)
	return token, ok
}

// WithTokenContext Set the token in the context.
func WithTokenContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey{}, token)
}
