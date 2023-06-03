package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterQuestionnaireRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	questionnaireGroup := v1Group.Group("/questionnaires", middleware.AuthMiddleware)
	questionnaireGroup.GET("", getQuestionnaire)
	questionnaireGroup.POST("", getQuestionnaire)
}

func getQuestionnaire(c echo.Context) error {
	data, statusCode, err := repository.GetQuestionnaire()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}
