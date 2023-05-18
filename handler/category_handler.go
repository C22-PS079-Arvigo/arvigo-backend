package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterCategoryRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	locationGroup := v1Group.Group("/categories", middleware.AuthMiddleware)
	locationGroup.GET("", getCategories)
}

func getCategories(c echo.Context) error {
	data, statusCode, err := repository.GetCategories()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}
