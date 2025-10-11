package main

import (
	"context"
	"net/http"
	"os"

	"github.com/Lakeshore-Labs/templecho"
	"github.com/Lakeshore-Labs/templecho/auth"
	"github.com/Lakeshore-Labs/templecho/blocks/layout"
	"github.com/Lakeshore-Labs/templecho/blocks/navbar"
	"github.com/Lakeshore-Labs/templecho/components/badge"
	"github.com/Lakeshore-Labs/templecho/components/icon"
	"github.com/Lakeshore-Labs/templecho/examples/demo/handlers"
	"github.com/Lakeshore-Labs/templecho/examples/demo/static"
	"github.com/Lakeshore-Labs/templecho/examples/demo/urls"
	"github.com/labstack/echo/v4"
)

func main() {
	// Create server with default configuration
	srv := templecho.New(
		templecho.WithPort(getEnvOrDefault("PORT", "3000")),
	)

	// Setup routes
	setupRoutes(srv)

	// Start the server
	srv.Start()
}

// setupRoutes configures all routes for the demo app
func setupRoutes(srv *templecho.Server) {
	// Apply OptionalAuth middleware to all routes so templates can access user info
	srv.Echo.Use(auth.OptionalAuth)

	// Setup assets and layout BEFORE registering routes
	srv.WithHTMX().
		WithTemplUI().
		WithLayout(createLayoutConfig).
		ServeStaticFS("/static", http.FS(static.FS()))

	// Home page - redirect to product
	srv.Echo.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/product")
	})

	// Product pages
	srv.Echo.GET("/product", handlers.ProductHandler)
	srv.Echo.GET("/product/detail", handlers.ProductDetailHandler)

	// AI Prompt page
	srv.Echo.GET("/prompt", handlers.PromptHandler)

	// AI Prompt with Metadata page
	srv.Echo.GET("/prompt-metadata", handlers.PromptWithMetadataHandler)

	// Calendar page
	srv.Echo.GET("/calendar", handlers.CalendarHandler)

	// Marketing Split signup page
	srv.Echo.GET("/marketing-split", handlers.MarketingSplitHandler)

	// Portfolio Showcase page
	srv.Echo.GET("/showcase", handlers.ShowcaseHandler)

	// Gallery Grid page
	srv.Echo.GET("/grid", handlers.GridHandler)

	// Login page
	srv.Echo.GET("/login", handlers.LoginHandler)

	// Profile page
	srv.Echo.GET("/profile", handlers.ProfileHandler)

	// Integrations page
	srv.Echo.GET("/integrations", handlers.IntegrationsHandler)
}

// createLayoutConfig creates the layout configuration
func createLayoutConfig(ctx context.Context) *layout.LayoutConfig {
	return &layout.LayoutConfig{
		Title:         "Demo - Example Application",
		Brand:         getBrandConfig(),
		SidebarRoutes: getSidebarRoutes(ctx),
		NavbarConfig:  getNavbarConfig(ctx),
		StaticCSS: []string{
			urls.StaticCSS(ctx, "output.css"),
		},
		StaticJS: []string{
			urls.StaticJS(ctx, "app.js"),
		},
	}
}

// getBrandConfig returns the brand configuration
func getBrandConfig() navbar.BrandConfig {
	return navbar.BrandConfig{
		LogoIcon: icon.Layers(),
		Name:     "Demo App",
		Subtitle: "Example Dashboard",
	}
}

// getSidebarRoutes returns all navigation routes for the sidebar
func getSidebarRoutes(ctx context.Context) []navbar.MenuItem {
	currentPath := auth.GetCurrentPath(ctx)
	user := auth.GetUserFromContext(ctx)

	routes := []navbar.MenuItem{
		{Name: "Product", Path: "/product"},
		{Name: "Product Detail", Path: "/product/detail"},
		{Name: "AI Prompt", Path: "/prompt"},
		{Name: "AI Prompt with Metadata", Path: "/prompt-metadata"},
		{Name: "Calendar", Path: "/calendar"},
		{Name: "Marketing Split", Path: "/marketing-split"},
		{Name: "Showcase", Path: "/showcase"},
		{Name: "Gallery Grid", Path: "/grid"},
		{Name: "Login", Path: "/login"},
		{Name: "Profile", Path: "/profile"},
		{Name: "Integrations", Path: "/integrations"},
	}

	// Add auth-specific items
	if user != nil {
		routes = append(routes, navbar.MenuItem{Name: "Account", Path: "/account"})
	}

	routes = append(routes, navbar.MenuItem{Name: "About", Path: urls.About(ctx)})

	// Set active state
	for i := range routes {
		routes[i].IsActive = routes[i].Path == currentPath
	}

	return routes
}

// getNavbarConfig creates the navbar configuration based on the current context
func getNavbarConfig(ctx context.Context) navbar.NavbarConfig {
	user := auth.GetUserFromContext(ctx)
	currentPath := auth.GetCurrentPath(ctx)

	// Get navbar routes
	routes := []navbar.MenuItem{
		{Name: "Dashboard", Path: urls.Home(ctx), Icon: icon.LayoutDashboard(icon.Props{Size: 16})},
		{Name: "Gallery", Path: urls.Gallery(ctx), Icon: icon.FolderOpen(icon.Props{Size: 16})},
		{Name: "Products", Path: "/product", Icon: icon.Check(icon.Props{Size: 16})},
		{Name: "Analytics", Path: "/generated-images", Icon: icon.ChartBar(icon.Props{Size: 16})},
	}

	// Set active state
	for i := range routes {
		routes[i].IsActive = routes[i].Path == currentPath
	}

	config := navbar.NavbarConfig{
		Brand:             getBrandConfig(),
		MenuItems:         routes,
		ShowSearch:        true,
		ShowNotifications: true,
		ShowMessages:      true,
		NotificationCount: 1,
	}

	if user != nil {
		// Use proxy endpoint for avatar to avoid CORS issues
		avatarURL := ""
		if user.Avatar != "" {
			avatarURL = "/avatar"
		}

		// Determine display name using shared utility
		displayName := auth.GetUserDisplayName(user)

		config.User = &navbar.UserConfig{
			ID:     user.ID,
			Name:   displayName,
			Email:  user.Email,
			Avatar: avatarURL, // Use proxy endpoint
			Badges: []navbar.UserBadge{
				{Text: "Demo", Variant: string(badge.VariantDefault), Class: "bg-yellow-500/10 text-yellow-700 dark:text-yellow-400"},
				{Text: "Premium", Variant: string(badge.VariantSecondary)},
				{Text: "Online", Variant: string(badge.VariantSecondary), Class: "bg-primary/10 text-primary"},
			},
		}

		config.UserMenuItems = []navbar.UserMenuItem{
			{Name: "View Profile", Path: "/profile", Icon: icon.User(icon.Props{Size: 16})},
			{Name: "Help & Support", Path: "/help", Icon: icon.CircleHelp(icon.Props{Size: 16}), Separator: true},
			{Name: "Sign Out", Path: "/logout", Icon: icon.LogOut(icon.Props{Size: 16, Class: "text-destructive"}), Class: "hover:bg-destructive/10"},
		}
	} else {
		// User is not logged in - show Sign In button
		config.ShowSignIn = true
		config.SignInPath = "/login"
	}

	return config
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}