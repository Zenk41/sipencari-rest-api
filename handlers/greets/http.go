package greets

import (
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/labstack/echo/v4"
)

type GreetingHandler interface {
	Greeting(c echo.Context) error
}

type greetingHandler struct{}

func NewGreetingHandler() GreetingHandler {
	return &greetingHandler{}
}

func (gh *greetingHandler) Greeting(c echo.Context) error {
	return response.NewResponseSuccess(c, http.StatusOK, "success", "Greetings", echo.Map{"Message": "Welcome to Sipencari API"})
}
