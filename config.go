package templecho

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Config holds server configuration
type Config struct {
	Port          string
	DatabaseURL   string
	SessionSecret string
	SessionName   string
	Environment   string
	Debug         bool
	HideBanner    bool
	LogLevel      string
	AppName       string // Application name for branding
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Port:        "8080",
		SessionName: "templecho-session",
		HideBanner:  true,
		Debug:       false,
		LogLevel:    "info",
		AppName:     "TemplEcho App",
		Environment: "development",
	}
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv() *Config {
	// Load .env file if it exists (ignore error as it's optional)
	godotenv.Load()

	cfg := DefaultConfig()

	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		cfg.DatabaseURL = dbURL
	}

	if secret := os.Getenv("SESSION_SECRET"); secret != "" {
		cfg.SessionSecret = secret
	}

	if name := os.Getenv("SESSION_NAME"); name != "" {
		cfg.SessionName = name
	}

	if env := os.Getenv("ENVIRONMENT"); env != "" {
		cfg.Environment = env
	} else if env := os.Getenv("ENV"); env != "" {
		cfg.Environment = env
	}

	if cfg.Environment == "development" || cfg.Environment == "dev" {
		cfg.Debug = true
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		cfg.LogLevel = logLevel
	}

	if appName := os.Getenv("APP_NAME"); appName != "" {
		cfg.AppName = appName
	}

	return cfg
}

// NewWithConfig creates a new server instance with a configuration
func NewWithConfig(cfg *Config) *Server {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	e := echo.New()
	e.HideBanner = cfg.HideBanner
	e.Debug = cfg.Debug

	s := &Server{
		Echo:   e,
		Port:   cfg.Port,
		Config: cfg,
	}

	// Setup default middleware
	s.setupDefaultMiddleware()

	return s
}
