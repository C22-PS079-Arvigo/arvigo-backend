package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterCronJobRoutes(e *echo.Echo) {
	e.POST("/v1/cron-job/subscription", subscriptionCronJob)
}

func subscriptionCronJob(c echo.Context) error {
	statusCode, err := repository.SubscriptionCronJob()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", nil, statusCode)
}
