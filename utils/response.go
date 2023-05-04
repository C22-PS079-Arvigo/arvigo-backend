package utils

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c echo.Context, message string, data interface{}, statusCode int) error {
	response := Response{
		Message: capitalizeSentences(message),
		Data:    data,
	}

	return c.JSON(statusCode, response)
}
