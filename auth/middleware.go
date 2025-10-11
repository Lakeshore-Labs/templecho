package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ContextKey is a custom type for context keys
type ContextKey string

const (
	// UserContextKey is the key for storing user in context
	UserContextKey ContextKey = "user"
)

// RequireAuth middleware ensures user is authenticated
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := GetUser(c)
		if err != nil {
			// Session error
			return c.Redirect(http.StatusFound, "/login")
		}

		if user == nil {
			// Not authenticated
			return c.Redirect(http.StatusFound, "/login")
		}

		// Store user in context for handlers
		c.Set(string(UserContextKey), user)

		return next(c)
	}
}

// OptionalAuth middleware adds user to context if authenticated
func OptionalAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := GetUser(c)
		if err == nil && user != nil {
			// Store user in Echo context for handlers
			c.Set(string(UserContextKey), user)

			// Also store in request context for templates
			ctx := context.WithValue(c.Request().Context(), UserContextKey, user)
			c.SetRequest(c.Request().WithContext(ctx))
		}

		return next(c)
	}
}

// GetCurrentUser retrieves user from context
func GetCurrentUser(c echo.Context) *User {
	if user, ok := c.Get(string(UserContextKey)).(*User); ok {
		return user
	}
	return nil
}