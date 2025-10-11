package navbar

import (
	"context"

	"github.com/a-h/templ"
)

// NavbarConfig contains all configuration for the navbar component
type NavbarConfig struct {
	Brand             BrandConfig
	User              *UserConfig
	MenuItems         []MenuItem
	UserMenuItems     []UserMenuItem
	ShowSearch        bool
	ShowNotifications bool
	ShowMessages      bool
	NotificationCount int
	ShowSignIn        bool   // Show sign in button when user is not authenticated
	SignInPath        string // Path for sign in button

}

// BrandConfig contains branding information
type BrandConfig struct {
	LogoIcon templ.Component // Icon component for the logo
	Name     string          // Company/App name
	Subtitle string          // Optional subtitle
}

// UserConfig contains user information
type UserConfig struct {
	ID     string
	Name   string
	Email  string
	Avatar string // URL to avatar image
	Badges []UserBadge
}

// UserBadge represents a badge shown next to user info
type UserBadge struct {
	Text    string
	Variant string // "secondary", "primary", etc.
	Class   string // Additional CSS classes
}

// MenuItem represents a navigation menu item
type MenuItem struct {
	Name      string
	Path      string
	Icon      templ.Component                // Icon component
	IsActive  bool                           // Whether this item is currently active
	IsVisible func(ctx context.Context) bool // Optional visibility predicate
}

// UserMenuItem represents an item in the user dropdown menu
type UserMenuItem struct {
	Name      string
	Path      string
	Icon      templ.Component // Icon component
	Class     string          // Additional CSS classes
	Separator bool            // Whether to show separator before this item
}
