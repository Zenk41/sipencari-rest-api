package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Zenk41/sipencari-rest-api/initialize"
	"github.com/Zenk41/sipencari-rest-api/routes"
	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/labstack/echo/v4"

	_sql "github.com/Zenk41/sipencari-rest-api/db/sql"
)

func main() {

	dbSQL := initialize.Init()
	e := echo.New()

	routes.RouteRegister(e)

	go func() {
		if err := e.Start(":" + util.GetEnv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down server")
		}
	}()
	wait := util.GracefulShutdown(context.Background(), 2*time.Second, map[string]util.Operation{
		"database": func(ctx context.Context) error {
			return _sql.CloseDB(dbSQL) // Closing Database
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background()) // Closing Server
		},
	})

	<-wait
}
