package middlewares

import (
	"github.com/Zenk41/sipencari-rest-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: config.LoadLoggerConfig().Format,
	})
}
