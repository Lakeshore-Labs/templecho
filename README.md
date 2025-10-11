# Templecho

A Go package for building modern web applications using Echo, templ, and Tailwind CSS. Templecho provides a batteries-included framework with 40+ UI components, HTMX support, and a flexible layout system.

## Features

- **Echo Framework**: High-performance Go web framework
- **Templ Templating**: Type-safe HTML templates in Go
- **40+ UI Components**: Pre-built, customizable components using Tailwind CSS
- **HTMX Support**: Build dynamic applications without JavaScript frameworks
- **Flexible Layout System**: Context-based layout configuration
- **Asset Management**: Built-in support for static files and embedded assets
- **Dark Mode Support**: Full dark mode theming out of the box

## Installation

```bash
go get github.com/Lakeshore-Labs/templecho
```

## Quick Start

Create a simple server with layout and components:

```go
package main

import (
    "context"
    "net/http"

    "github.com/Lakeshore-Labs/templecho"
    "github.com/Lakeshore-Labs/templecho/blocks/layout"
    "github.com/Lakeshore-Labs/templecho/blocks/navbar"
    "github.com/labstack/echo/v4"
)

func main() {
    // Create a new server
    srv := templecho.New(
        templecho.WithPort("3000"),
    )

    // Enable HTMX and TemplUI components
    srv.WithHTMX().
        WithTemplUI().
        WithLayout(createLayoutConfig)

    // Add routes
    srv.Echo.GET("/", homeHandler)

    // Start the server
    srv.Start()
}

func createLayoutConfig(ctx context.Context) *layout.LayoutConfig {
    return &layout.LayoutConfig{
        Title: "My App",
        Brand: navbar.BrandConfig{
            Name: "My Application",
        },
        NavbarConfig: navbar.NavbarConfig{
            MenuItems: []navbar.MenuItem{
                {Name: "Home", Path: "/"},
                {Name: "About", Path: "/about"},
            },
        },
    }
}

func homeHandler(c echo.Context) error {
    return templecho.Render(c, HomePage())
}
```

## Using Components

Templecho includes 40+ pre-built components. Here are some examples:

### Cards

```go
import "github.com/Lakeshore-Labs/templecho/components/card"

templ HomePage() {
    @card.Card(card.Props{
        Title: "Welcome",
        Description: "Get started with Templecho",
    }) {
        <p>Your content here</p>
    }
}
```

### Buttons

```go
import "github.com/Lakeshore-Labs/templecho/components/button"

templ ActionButtons() {
    @button.Button(button.Props{
        Text: "Primary Action",
        Variant: button.VariantPrimary,
    })

    @button.Button(button.Props{
        Text: "Secondary",
        Variant: button.VariantSecondary,
        Size: button.SizeSmall,
    })
}
```

### Forms

```go
import (
    "github.com/Lakeshore-Labs/templecho/components/input"
    "github.com/Lakeshore-Labs/templecho/components/label"
)

templ LoginForm() {
    <form>
        @label.Label(label.Props{
            Text: "Email",
            For: "email",
        })
        @input.Input(input.Props{
            ID: "email",
            Type: "email",
            Placeholder: "Enter your email",
        })
    </form>
}
```

## Available Components

### Layout Components
- `blocks/layout` - Main application layout with sidebar and navbar
- `blocks/navbar` - Navigation bar with user menu
- `blocks/sidebar` - Collapsible sidebar navigation

### UI Components
- `accordion` - Expandable content sections
- `alert` - Alert messages and notifications
- `avatar` - User avatars with fallback
- `badge` - Status badges and labels
- `breadcrumb` - Navigation breadcrumbs
- `button` - Various button styles and sizes
- `calendar` - Date picker calendar
- `card` - Content cards with headers and actions
- `checkbox` - Checkbox inputs
- `dialog` - Modal dialogs
- `dropdown` - Dropdown menus
- `input` - Text inputs with validation
- `label` - Form labels
- `pagination` - Page navigation
- `progress` - Progress bars
- `radio` - Radio button groups
- `select` - Select dropdowns
- `separator` - Visual separators
- `sheet` - Slide-out panels
- `skeleton` - Loading skeletons
- `slider` - Range sliders
- `switch` - Toggle switches
- `table` - Data tables
- `tabs` - Tab navigation
- `textarea` - Multi-line text inputs
- `tooltip` - Hover tooltips
- And more...

## Layout System

The layout system provides a flexible way to configure your application's structure:

```go
func createLayoutConfig(ctx context.Context) *layout.LayoutConfig {
    user := auth.GetUserFromContext(ctx)

    return &layout.LayoutConfig{
        Title: "My Application",
        Brand: navbar.BrandConfig{
            LogoIcon: icon.Layers(),
            Name: "My App",
            Subtitle: "Dashboard",
        },
        SidebarRoutes: []navbar.MenuItem{
            {Name: "Dashboard", Path: "/", Icon: icon.LayoutDashboard()},
            {Name: "Users", Path: "/users", Icon: icon.Users()},
            {Name: "Settings", Path: "/settings", Icon: icon.Settings()},
        },
        NavbarConfig: navbar.NavbarConfig{
            ShowSearch: true,
            ShowNotifications: true,
            User: user != nil ? &navbar.UserConfig{
                Name: user.Name,
                Email: user.Email,
                Avatar: user.AvatarURL,
            } : nil,
        },
        StaticCSS: []string{"/static/css/app.css"},
        StaticJS: []string{"/static/js/app.js"},
    }
}
```

## HTMX Integration

Enable HTMX for dynamic UI updates without full page refreshes:

```go
srv.WithHTMX()

// In your templates
templ Button() {
    <button
        hx-get="/api/data"
        hx-target="#result"
        hx-swap="innerHTML">
        Load Data
    </button>
    <div id="result"></div>
}
```

## Static Assets

Serve static files and embedded assets:

```go
//go:embed static/*
var staticFS embed.FS

func main() {
    srv := templecho.New()

    // Serve from filesystem
    srv.ServeStatic("/static", "path/to/static")

    // Serve from embedded FS
    srv.ServeStaticFS("/static", http.FS(staticFS))
}
```

## Middleware

Add middleware to your application:

```go
import "github.com/Lakeshore-Labs/templecho/auth"

srv.Echo.Use(auth.OptionalAuth)
srv.Echo.Use(middleware.Logger())
srv.Echo.Use(middleware.Recover())
```

## Configuration Options

```go
srv := templecho.New(
    templecho.WithPort("8080"),
    templecho.WithHTMX(),
    templecho.WithTemplUI(),
    templecho.WithLayout(layoutFunc),
)
```

## CSS Styling

Templecho uses Tailwind CSS for styling. For production, you should compile your CSS:

### Using Tailwind CDN (Development)

For quick development, Tailwind CDN is included by default in the layout.

### Building for Production

1. Install Tailwind CSS v4:
```bash
npm install tailwindcss@next @tailwindcss/cli@next
```

2. Create your CSS input file:
```css
@import "tailwindcss";
```

3. Build your CSS:
```bash
npx tailwindcss -i input.css -o output.css
```

4. Include the compiled CSS in your layout configuration:
```go
layoutConfig.StaticCSS = []string{"/static/css/output.css"}
```

## Examples

See the `examples/demo` directory for a complete example application demonstrating:
- Product catalog with cards
- User authentication flow
- Dashboard layouts
- Form handling
- HTMX interactions
- Dark mode theming

To run the demo:

```bash
cd examples/demo
go run .
```

Then visit http://localhost:3000

## Development

### Prerequisites

- Go 1.21+
- templ CLI for template generation
- Tailwind CSS (optional, for custom styling)

### Building Templates

Generate templ files:

```bash
go run github.com/a-h/templ/cmd/templ@latest generate
```

### Project Structure

```
templecho/
├── server.go              # Main server implementation
├── render.go              # Template rendering utilities
├── layout/                # Layout middleware and configuration
├── auth/                  # Authentication utilities
├── assets/                # Embedded static assets
├── blocks/                # Layout blocks (navbar, sidebar, etc.)
│   ├── layout/
│   ├── navbar/
│   └── sidebar/
├── components/            # UI components
│   ├── accordion/
│   ├── alert/
│   ├── avatar/
│   ├── badge/
│   ├── button/
│   ├── card/
│   └── ... (35+ more)
└── examples/
    └── demo/              # Example application
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Credits

Built with:
- [Echo](https://echo.labstack.com/) - High performance Go web framework
- [templ](https://templ.guide/) - Type-safe HTML templates for Go
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [HTMX](https://htmx.org/) - High power tools for HTML