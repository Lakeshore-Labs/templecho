package auth

import (
	"context"
)

// PathContextKey is the key for storing the current path in context
type PathContextKey string

const CurrentPathKey PathContextKey = "currentPath"

// GetUserFromContext extracts user from context (for use in templates)
func GetUserFromContext(ctx context.Context) *User {
	// In Echo, the context values are stored in the request context
	if val := ctx.Value(UserContextKey); val != nil {
		if user, ok := val.(*User); ok {
			return user
		}
	}
	return nil
}

// GetCurrentPath extracts the current request path from context
func GetCurrentPath(ctx context.Context) string {
	// Try the new context key first
	if path, ok := ctx.Value("currentPath").(string); ok {
		return path
	}
	// Fallback to the old key
	if path, ok := ctx.Value(CurrentPathKey).(string); ok {
		return path
	}
	return ""
}