package route

import (
	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/middleware"
)

func RegisterUserRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	_ = v1Group.Group("/users", middleware.AuthMiddleware)

	// userGroup.GET("", getUsersHandler)
	// userGroup.GET("/:id", getUserHandler)
	// userGroup.POST("", func(c echo.Context) error {
	// 	return
	// })
	// userGroup.PUT("/:id", updateUserHandler)
	// userGroup.DELETE("/:id", deleteUserHandler)
}

// move handler! rename repository to repositroy!
// func createUserHandler(c echo.Context) error {
// 	user, err := repository.Register(c)
// 	if err != nil {
// 		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
// 	}

// 	return utils.ResponseJSON(c, "Created", user, http.StatusCreated)
// }

// func getUsersHandler(c echo.Context) error {
// 	users, err := repository.GetUsers()
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, users)
// }

// func getUserHandler(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
// 	}

// 	user, err := repository.GetUser(userID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func updateUserHandler(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
// 	}

// 	var userData repository.UpdateUserInput
// 	err = c.Bind(&userData)
// 	if err != nil {
// 		return err
// 	}

// 	user, err := repository.UpdateUser(userID, userData)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func deleteUserHandler(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
// 	}

// 	err = repository.DeleteUser(userID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.NoContent(http.StatusOK)
// }
