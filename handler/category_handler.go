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
	catGroup := v1Group.Group("/categories", middleware.AuthMiddleware)
	catGroup.GET("", getCategories)

	catGroup.GET("/:id/list-product", getListProductByCategory)
}

func getCategories(c echo.Context) error {
	data, statusCode, err := repository.GetCategories()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getListProductByCategory(c echo.Context) error {
	catID := utils.StrToUint64(c.Param("id"), 0)
	if catID == 0 {
		return utils.ResponseJSON(c, "Invalid category ID", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetListProductByCategory(catID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}
