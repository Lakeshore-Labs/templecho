package templecho

import (
	"github.com/Lakeshore-Labs/templecho/layout"
)

// WithLayout configures the server with a layout function that will be called
// for each request to generate the layout configuration
func (s *Server) WithLayout(layoutFunc layout.LayoutFunc) *Server {
	s.Echo.Use(layout.Middleware(s.htmxEnabled, layoutFunc))
	return s
}
