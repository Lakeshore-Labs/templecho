package navbar

import (
	"context"

	"github.com/Lakeshore-Labs/templecho/auth"
)

// RequiresAuth returns a visibility function that checks if user is authenticated
func RequiresAuth() func(context.Context) bool {
	return func(ctx context.Context) bool {
		return auth.GetUserFromContext(ctx) != nil
	}
}

// RequiresNoAuth returns a visibility function that checks if user is NOT authenticated
func RequiresNoAuth() func(context.Context) bool {
	return func(ctx context.Context) bool {
		return auth.GetUserFromContext(ctx) == nil
	}
}

// AlwaysVisible returns a visibility function that always returns true
func AlwaysVisible() func(context.Context) bool {
	return func(ctx context.Context) bool {
		return true
	}
}

// RequiresRole returns a visibility function that checks if user has a specific role
// This is a placeholder for future RBAC implementation
func RequiresRole(role string) func(context.Context) bool {
	return func(ctx context.Context) bool {
		user := auth.GetUserFromContext(ctx)
		if user == nil {
			return false
		}
		// TODO: Implement role checking when RBAC is added
		// For now, just check if user is authenticated
		return true
	}
}

// RequiresFeature returns a visibility function that checks if a feature flag is enabled
// This is a placeholder for future feature flag implementation
func RequiresFeature(feature string) func(context.Context) bool {
	return func(ctx context.Context) bool {
		// TODO: Implement feature flag checking
		// For now, always return true
		return true
	}
}

// RequiresAuthAndFeature combines auth check with feature flag check
func RequiresAuthAndFeature(feature string) func(context.Context) bool {
	return func(ctx context.Context) bool {
		return RequiresAuth()(ctx) && RequiresFeature(feature)(ctx)
	}
}
