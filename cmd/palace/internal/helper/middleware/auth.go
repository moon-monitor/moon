package middleware

import (
	"context"
)

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
