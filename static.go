package templecho

import (
	"net/http"

	"github.com/Lakeshore-Labs/templecho/assets"
	"github.com/Lakeshore-Labs/templecho/htmx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// HTMXPath is the URL path where HTMX library is served
	HTMXPath = "/assets/htmx.min.js"
)

// WithHTMX configures the server to serve HTMX library at the default path
func (s *Server) WithHTMX() *Server {
	s.Echo.GET(HTMXPath, echo.WrapHandler(htmx.Handler()))
	s.htmxEnabled = true
	return s
}

// WithTemplUI configures the server to serve TemplUI assets
func (s *Server) WithTemplUI() *Server {
	s.Echo.GET("/assets/*", echo.WrapHandler(
		http.StripPrefix("/assets/", http.FileServer(assets.GetAssetsFS())),
	))
	return s
}

// WithHealthCheck adds a health check endpoint at /health
func (s *Server) WithHealthCheck() *Server {
	s.Echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return s
}

// ServeStaticFS configures the server to serve static files from an embedded filesystem
func (s *Server) ServeStaticFS(prefix string, filesystem http.FileSystem) *Server {
	staticGroup := s.Echo.Group(prefix)
	staticGroup.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "/",
		Browse:     false,
		Filesystem: filesystem,
	}))
	return s
}
