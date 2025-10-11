package layout

import (
	"context"

	"github.com/Lakeshore-Labs/templecho/blocks/navbar"
	"github.com/a-h/templ"
)

// LayoutType defines the type of layout to use
type LayoutType string

const (
	// LayoutTypeAuto automatically selects layout based on configuration
	LayoutTypeAuto LayoutType = "auto"
	// LayoutTypeSidebar forces sidebar layout even if no sidebar routes
	LayoutTypeSidebar LayoutType = "sidebar"
	// LayoutTypeHeader uses header-only layout (no sidebar)
	LayoutTypeHeader LayoutType = "header"
	// LayoutTypeMinimal uses minimal layout (no navigation)
	LayoutTypeMinimal LayoutType = "minimal"
	// LayoutTypeFullPage uses full-page layout (no navigation at all)
	LayoutTypeFullPage LayoutType = "fullpage"
)

// LayoutConfig holds the configuration for the layout
type LayoutConfig struct {
	Title              string
	Brand              navbar.BrandConfig
	SidebarRoutes      []navbar.MenuItem
	MobileFooterRoutes []navbar.MenuItem // Mobile footer navigation items
	NavbarConfig       navbar.NavbarConfig
	LayoutType         LayoutType      // Layout selector
	LayoutComponent    templ.Component // Direct component override
	// Asset paths
	StaticCSS []string
	StaticJS  []string
}

// GetLayout returns the appropriate layout component based on configuration
func (lc *LayoutConfig) GetLayout() templ.Component {
	// Direct override takes precedence
	if lc.LayoutComponent != nil {
		return lc.LayoutComponent
	}

	// Select based on LayoutType
	switch lc.LayoutType {
	case LayoutTypeFullPage:
		return Layout005()
	case LayoutTypeHeader:
		return Layout004()
	case LayoutTypeSidebar:
		return Layout001()
	case LayoutTypeMinimal:
		// TODO: Add minimal layout when needed
		return Layout004() // Default to header for now
	case LayoutTypeAuto:
		fallthrough
	default:
		// Auto-select based on configuration
		if len(lc.SidebarRoutes) > 0 {
			return Layout001() // Has sidebar routes, use sidebar layout
		}
		return Layout004() // No sidebar routes, use header-only layout
	}
}

// ContextKey is the type for context keys
type ContextKey string

const (
	// LayoutConfigKey is the context key for layout configuration - must match layout package
	LayoutConfigKey ContextKey = "layoutConfig"
)

// GetLayoutConfig retrieves the layout configuration from context
func GetLayoutConfig(ctx context.Context) *LayoutConfig {
	if config, ok := ctx.Value(LayoutConfigKey).(*LayoutConfig); ok {
		if config.LayoutComponent == nil {
			config.LayoutComponent = config.GetLayout()
		}
		return config
	}
	return nil
}

// WithLayoutConfig adds layout configuration to context
func WithLayoutConfig(ctx context.Context, config *LayoutConfig) context.Context {
	return context.WithValue(ctx, LayoutConfigKey, config)
}
