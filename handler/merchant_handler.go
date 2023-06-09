package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterMerchantRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")

	merchantGroup := v1Group.Group("/merchant-app", middleware.AuthMiddleware)
	merchantGroup.GET("/home", getMerchantAppHome)
	// merchantGroup.GET("/product/:id", getHomeMerchant)
}

func getMerchantAppHome(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)

	data, statusCode, err := repository.GetMerchantAppHome(userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

// func getHomeMerchant(c echo.Context) error {
// 	data, statusCode, err := repository.GetHomeMerchant()
// 	if err != nil {
// 		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
// 	}

// 	return utils.ResponseJSON(c, "Success", data, statusCode)
// }

// func getHomeSearch(c echo.Context) error {
// 	search := strings.TrimSpace(strings.ToLower(c.Param("search")))
// 	if search == "" {
// 		return utils.ResponseJSON(c, "Search must be filled", nil, http.StatusBadRequest)
// 	}

// 	data, statusCode, err := repository.GetHomeSearch(search)
// 	if err != nil {
// 		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
// 	}

// 	return utils.ResponseJSON(c, "Success", data, statusCode)
// }
