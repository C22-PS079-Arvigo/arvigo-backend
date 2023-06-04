package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterFaceShapeRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	locationGroup := v1Group.Group("/face-shape", middleware.AuthMiddleware)
	locationGroup.POST("/check", faceShapeRecognitionHandler)
}

func faceShapeRecognitionHandler(c echo.Context) error {
	// Parse the form data
	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	data, statusCode, err := repository.FaceShapeRecognition(form, userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success post data", data, http.StatusOK)
}
