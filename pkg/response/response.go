package response

import (
	"github.com/labstack/echo/v4"
)

func WithSuccess(c echo.Context, status int, data any) error {
	return c.JSON(status, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func WithError(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}
