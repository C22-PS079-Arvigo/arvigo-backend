package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type CustomLogger struct {
	Logger echo.Logger
}

func (c *CustomLogger) Write(p []byte) (n int, err error) {
	message := string(p)
	// Check if the log entry contains the message to be hidden
	if strings.Contains(message, "[error] unsupported data type: &[]") {
		// Skip writing this log entry
		return len(p), nil
	}
	return c.Logger.Output().Write(p)
}
