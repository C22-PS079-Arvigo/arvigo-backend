package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterHomeRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")

	homeGroup := v1Group.Group("/homes", middleware.AuthMiddleware)
	homeGroup.GET("", getHome)
	homeGroup.GET("/merchant", getHomeMerchant)
	homeGroup.GET("/search/:search", getHomeSearch)
}

func getHome(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)

	data, statusCode, err := repository.GetHome(userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getHomeMerchant(c echo.Context) error {
	data, statusCode, err := repository.GetHomeMerchant()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getHomeSearch(c echo.Context) error {
	search := strings.TrimSpace(strings.ToLower(c.Param("search")))
	if search == "" {
		return utils.ResponseJSON(c, "Search must be filled", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetHomeSearch(search)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}
