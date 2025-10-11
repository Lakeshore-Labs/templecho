package urls

import "context"

// URL paths for static assets
const (
	HTMXPath = "/assets/htmx.min.js"
)

// Home returns the home page URL
func Home(ctx context.Context) string {
	return "/"
}

// Gallery returns the gallery page URL
func Gallery(ctx context.Context) string {
	return "/grid"
}

// About returns the about page URL
func About(ctx context.Context) string {
	return "/showcase"
}

// HTMX returns the HTMX script URL
func HTMX() string {
	return HTMXPath
}

// StaticCSS returns the CSS file path
func StaticCSS(ctx context.Context, file string) string {
	return "/static/css/" + file
}

// StaticJS returns the JS file path
func StaticJS(ctx context.Context, file string) string {
	return "/static/js/" + file
}