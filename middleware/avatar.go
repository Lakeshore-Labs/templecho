package middleware

import "context"

type contextKey string

const avatarDataURIKey contextKey = "avatarDataURI"

// WithAvatarDataURI adds an avatar data URI to the context
func WithAvatarDataURI(ctx context.Context, dataURI string) context.Context {
	return context.WithValue(ctx, avatarDataURIKey, dataURI)
}

// GetAvatarDataURI retrieves the avatar data URI from context
// Returns empty string if not set
func GetAvatarDataURI(ctx context.Context) string {
	if dataURI, ok := ctx.Value(avatarDataURIKey).(string); ok {
		return dataURI
	}
	return ""
}
