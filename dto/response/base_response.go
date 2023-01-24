package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error"`
}

func NewResponseSuccess(c echo.Context, statusCode int, statusMessage string, message string, data any) error {
	return c.JSON(statusCode, Response{
		Status:  statusMessage,
		Message: message,
		Data:    data,
	})
}

func NewResponseFailed(c echo.Context, statusCode int, statusMessage string, message string, data any, err string) error {
	return c.JSON(statusCode, Response{
		Status:  statusMessage,
		Message: message,
		Data:    nil,
		Error: err,
	})
}
