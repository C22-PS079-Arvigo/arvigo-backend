package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterQuestionnaireRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	questionnaireGroup := v1Group.Group("/questionnaires", middleware.AuthMiddleware)
	questionnaireGroup.GET("", getQuestionnaire)
	questionnaireGroup.POST("", generateQuestionnaireResult)
}

func getQuestionnaire(c echo.Context) error {
	data, statusCode, err := repository.GetQuestionnaire()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func generateQuestionnaireResult(c echo.Context) error {
	var data datastruct.QuestionnaireRequest
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	result, statusCode, err := repository.GenerateQuestionnaireResult(data, userAuth.ID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success generate data", result, http.StatusOK)
}
