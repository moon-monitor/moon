package permission

import (
	"context"
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

type operationContextKey struct{}

// GetOperationByContext Retrieve the operation from the context.
func GetOperationByContext(ctx context.Context) (string, bool) {
	operation, ok := ctx.Value(operationContextKey{}).(string)
	return operation, ok
}

// WithOperationContext Set the operation in the context.
func WithOperationContext(ctx context.Context, operation string) context.Context {
	return context.WithValue(ctx, operationContextKey{}, operation)
}
