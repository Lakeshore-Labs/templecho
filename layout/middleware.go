package layout

import (
	"context"
	"strings"

	"github.com/Lakeshore-Labs/templecho/blocks/layout"
	"github.com/labstack/echo/v4"
)

const (
	// IsHTMXKey is the context key for HTMX request detection
	IsHTMXKey = "isHTMX"
	// HTMXPath is the URL path where HTMX library is served
	HTMXPath = "/assets/htmx.min.js"
)

// LayoutFunc is a function that creates a layout configuration for a given context
type LayoutFunc func(ctx context.Context) *layout.LayoutConfig

var staticExtensions = []string{
	".css", ".js", ".png", ".jpg", ".jpeg", ".gif", ".ico",
	".svg", ".woff", ".woff2", ".ttf", ".eot", ".map",
}

var staticPrefixes = []string{
	"/api/",
	"/assets/",
	"/static/",
	"/health",
	"/favicon.ico",
}

// isStaticAssetPath determines if a request path is for a static asset
func isStaticAssetPath(path string) bool {
	for _, prefix := range staticPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}

	for _, ext := range staticExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	return false
}

// Middleware creates middleware that injects layout configuration
func Middleware(htmxEnabled bool, layoutFunc LayoutFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			path := c.Request().URL.Path

			// Skip layout entirely for static assets
			if isStaticAssetPath(path) {
				return next(c)
			}

			// Check if this is an HTMX request
			isHTMXRequest := c.Request().Header.Get("HX-Request") == "true"

			// Add the current path to context
			ctx = context.WithValue(ctx, "currentPath", path)

			if htmxEnabled && isHTMXRequest {
				// This is an HTMX request - just mark it and skip layout
				ctx = context.WithValue(ctx, IsHTMXKey, true)
				c.SetRequest(c.Request().WithContext(ctx))
			} else {
				// Regular request - create and inject layout config
				layoutConfig := layoutFunc(ctx)

				// If HTMX is enabled (but this isn't an HTMX request), add HTMX script
				if htmxEnabled && layoutConfig != nil {
					layoutConfig.StaticJS = append([]string{HTMXPath}, layoutConfig.StaticJS...)
				}

				ctx = layout.WithLayoutConfig(ctx, layoutConfig)
				c.SetRequest(c.Request().WithContext(ctx))
			}

			return next(c)
		}
	}
}

// IsHTMX checks if the current request is an HTMX request
func IsHTMX(ctx context.Context) bool {
	if val := ctx.Value(IsHTMXKey); val != nil {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}
