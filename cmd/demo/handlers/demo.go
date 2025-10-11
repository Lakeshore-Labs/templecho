package handlers

import (
	"fmt"
	"net/http"

	"github.com/Lakeshore-Labs/templecho/examples/demo/templates"
	"github.com/labstack/echo/v4"
)

// ProductHandler handles the product page
func ProductHandler(c echo.Context) error {

	component := templates.ProductPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// ProductDetailHandler handles the product detail page
func ProductDetailHandler(c echo.Context) error {
	component := templates.ProductDetailPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// PromptHandler handles the AI prompt page
func PromptHandler(c echo.Context) error {
	component := templates.PromptPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// CalendarHandler handles the calendar page
func CalendarHandler(c echo.Context) error {
	component := templates.CalendarPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// MarketingSplitHandler handles the marketing split signup page
func MarketingSplitHandler(c echo.Context) error {
	component := templates.MarketingSplitPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// LoginHandler handles the social login page
func LoginHandler(c echo.Context) error {
	component := templates.LoginPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// ShowcaseHandler handles the portfolio showcase page
func ShowcaseHandler(c echo.Context) error {
	component := templates.ShowcasePage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// GridHandler handles the gallery grid page
func GridHandler(c echo.Context) error {
	component := templates.GridPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// ProfileHandler handles the profile page
func ProfileHandler(c echo.Context) error {
	component := templates.ProfilePage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// IntegrationsHandler handles the integrations page
func IntegrationsHandler(c echo.Context) error {
	component := templates.IntegrationsPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}

// PromptWithMetadataHandler handles the AI prompt with metadata page
func PromptWithMetadataHandler(c echo.Context) error {
	component := templates.PromptWithMetadataPage()

	err := component.Render(c.Request().Context(), c.Response())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
	return nil
}