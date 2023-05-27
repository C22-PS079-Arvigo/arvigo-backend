package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ApiKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var apiKey = os.Getenv("X_API_KEY_SECRET")
	return func(c echo.Context) error {
		// Check if the API key is provided in the header
		providedAPIKey := c.Request().Header.Get("X-API-Key")
		if providedAPIKey == "" {
			// API key is not provided
			return echo.NewHTTPError(http.StatusUnauthorized, "API key is required")
		}

		// Verify the provided API key
		if providedAPIKey != apiKey {
			// Invalid API key
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
		}

		// API key is valid, call the next handler
		return next(c)
	}
}
