package templecho

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders a templ component to the response
func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

// RenderWithStatus renders a templ component with a specific HTTP status code
func RenderWithStatus(c echo.Context, status int, component templ.Component) error {
	c.Response().Status = status
	c.Response().Writer.WriteHeader(status)
	return component.Render(c.Request().Context(), c.Response())
}

// RenderError renders an error with the given status code
func RenderError(c echo.Context, status int, message string) error {
	return c.String(status, message)
}

// Handler wraps a templ component in an echo.HandlerFunc
// Usage: e.GET("/", templecho.Handler(myComponent()))
func Handler(component templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return Render(c, component)
	}
}

// HandlerFunc wraps a function that returns a templ component in an echo.HandlerFunc
// Usage: e.GET("/user/:id", templecho.HandlerFunc(func(c echo.Context) templ.Component {
//     return userPage(c.Param("id"))
// }))
func HandlerFunc(fn func(echo.Context) templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		component := fn(c)
		return Render(c, component)
	}
}

// HandlerFuncWithError wraps a function that returns a component and error
func HandlerFuncWithError(fn func(echo.Context) (templ.Component, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		component, err := fn(c)
		if err != nil {
			return err
		}
		return Render(c, component)
	}
}

// HTMLString returns a templ component from an HTML string
// Useful for quick testing or simple content
func HTMLString(html string) templ.Component {
	return templ.Raw(html)
}

// IsHTMXRequest checks if the request is an HTMX request
func IsHTMXRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

// HTMXTrigger sets the HX-Trigger response header
func HTMXTrigger(c echo.Context, event string) {
	c.Response().Header().Set("HX-Trigger", event)
}

// HTMXRedirect performs an HTMX redirect
func HTMXRedirect(c echo.Context, url string) error {
	c.Response().Header().Set("HX-Redirect", url)
	return c.NoContent(http.StatusOK)
}

// HTMXRefresh triggers a page refresh
func HTMXRefresh(c echo.Context) error {
	c.Response().Header().Set("HX-Refresh", "true")
	return c.NoContent(http.StatusOK)
}
