package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterAuthRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/login", loginHandler)
	authGroup.POST("/register", registerHandler)
}

func loginHandler(c echo.Context) error {
	var loginData datastruct.LoginUserInput
	if err := c.Bind(&loginData); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(loginData)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	token, statusCode, err := repository.Login(loginData)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), token, statusCode)
	}

	return utils.ResponseJSON(c, "Login success", token, http.StatusOK)
}

func registerHandler(c echo.Context) error {
	var userData datastruct.User
	err := c.Bind(&userData)
	if err != nil {
		return err
	}

	user, statusCode, err := repository.Register(userData)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Created", user, statusCode)
}
