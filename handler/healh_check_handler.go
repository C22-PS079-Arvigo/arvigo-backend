package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/pkg/database"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterHealthCheckRoutes(e *echo.Echo) {
	healthCheckGroup := e.Group("/health_check")
	healthCheckGroup.GET("/ping", healthCheckHandler)
}

func healthCheckHandler(c echo.Context) error {
	status := "healthy"
	_, err := database.ConnectDB()
	if err != nil {
		status = "unhealthy"
		log.Fatal(err)
	}

	return utils.ResponseJSON(c, status, map[string]interface{}{
		"dependencies": map[string]string{
			"mysql":                "OK",
			"google_cloud_storage": "OK",
		},
	}, http.StatusOK)
}
