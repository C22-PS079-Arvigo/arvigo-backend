package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterWishlistRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	wishlistGroup := v1Group.Group("/wishlists", middleware.AuthMiddleware)

	wishlistGroup.POST("", addToWishlist)
	wishlistGroup.GET("", getUserWishlist)
}

func addToWishlist(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	var data datastruct.AddWhislistProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.AddWhislistProduct(userAuth.ID, data)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", nil, statusCode)
}

func getUserWishlist(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)

	data, statusCode, err := repository.GetUserWishlist(userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}
