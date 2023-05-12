package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterUserRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	userGroup := v1Group.Group("/users", middleware.AuthMiddleware)

	userGroup.GET("/:id", getUserbyIDHandler)
	userGroup.GET("/user-list", getAllUsersHandler)
	userGroup.GET("/partner-list", getAllPartnersHandler)
}

// move handler! rename repository to repositroy!
func getUserbyIDHandler(c echo.Context) error {
	userID := utils.StrToUint64(c.Param("id"), 0)
	if userID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, statusCode, err := repository.GetUserByID(userID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", user, statusCode)
}

func getAllUsersHandler(c echo.Context) error {
	user, statusCode, err := repository.GetUsers()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", user, statusCode)
}

func getAllPartnersHandler(c echo.Context) error {
	user, statusCode, err := repository.GetPartners()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", user, statusCode)
}
