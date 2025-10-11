package templecho

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Server represents the HTTP server with templ and HTMX support
type Server struct {
	Echo         *echo.Echo
	Port         string
	Config       *Config
	htmxEnabled  bool
	cleanupFuncs []func()
}

// Option is a functional option for configuring the server
type Option func(*Server)

// WithPort sets the server port
func WithPort(port string) Option {
	return func(s *Server) {
		s.Port = port
	}
}

// WithDebug enables debug mode
func WithDebug(debug bool) Option {
	return func(s *Server) {
		s.Echo.Debug = debug
	}
}

// New creates a new server instance with functional options
func New(opts ...Option) *Server {
	cfg := DefaultConfig()

	e := echo.New()
	e.HideBanner = cfg.HideBanner
	e.Debug = cfg.Debug

	s := &Server{
		Echo:   e,
		Port:   cfg.Port,
		Config: cfg,
	}

	// Apply functional options
	for _, opt := range opts {
		opt(s)
	}

	// Setup default middleware
	s.setupDefaultMiddleware()

	return s
}

// setupDefaultMiddleware configures standard middleware
func (s *Server) setupDefaultMiddleware() {
	// Request ID for tracing
	s.Echo.Use(echoMiddleware.RequestID())

	// Logger with reasonable format
	s.Echo.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${status} ${method} ${uri} (${latency_human})\n",
		Output: os.Stdout,
	}))

	// Recoverer - recover from panics
	s.Echo.Use(echoMiddleware.Recover())

	// Compress responses
	s.Echo.Use(echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
		Level: 5,
	}))

	// CORS for development
	s.Echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
	}))
}

// WithCSRF enables CSRF protection (recommended for production)
func (s *Server) WithCSRF(secret string) *Server {
	s.Echo.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		Skipper: func(c echo.Context) bool {
			// Skip CSRF for API routes
			return strings.HasPrefix(c.Request().URL.Path, "/api/")
		},
		TokenLookup:    "form:csrf,header:X-CSRF-Token",
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: http.SameSiteLaxMode,
	}))
	return s
}

// WithTimeout sets request timeout
func (s *Server) WithTimeout(duration time.Duration) *Server {
	s.Echo.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Timeout: duration,
	}))
	return s
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.Echo.Start(":" + s.Port)
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Echo.Shutdown(ctx)
}

// OnShutdown registers a cleanup function to be called during shutdown
func (s *Server) OnShutdown(fn func()) {
	s.cleanupFuncs = append(s.cleanupFuncs, fn)
}

// RunWithGracefulShutdown starts the server and handles graceful shutdown
func (s *Server) RunWithGracefulShutdown(appName string) {
	// Setup graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Clean startup message
	fmt.Printf("\n🚀 %s\n", appName)
	fmt.Printf("⚡ Starting on http://localhost:%s\n", s.Port)
	fmt.Printf("📅 %s\n\n", time.Now().Format("Mon Jan 2 15:04:05 2006"))

	// Start server in goroutine
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("❌ Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()

	fmt.Println("\n🛑 Shutting down server...")

	// Run cleanup functions
	for _, cleanup := range s.cleanupFuncs {
		cleanup()
	}

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Perform graceful shutdown
	if err := s.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("❌ Server shutdown failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Server shutdown complete")
}
