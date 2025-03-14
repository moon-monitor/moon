package middleware

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
