package routes

import (
	// "github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/initialize"
	"github.com/Zenk41/sipencari-rest-api/middlewares"

	// "github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteRegister(e *echo.Echo) {
	e.Use(initialize.LoggerHandler)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1 := e.Group("/api/v1")

	// Greets
	v1.GET("/", initialize.GreetHandler.Greeting)

	// Auth
	Auth := v1.Group("/auth")
	Auth.POST("/register", initialize.UserHandler.Register)

	Auth.POST("/login", initialize.UserHandler.Login)

	// user
	User := v1.Group("/user", middleware.JWTWithConfig(initialize.AuthHandler))
	User.GET("/profile", initialize.UserHandler.MyProfile)
	User.GET("/profile/:user_id", initialize.UserHandler.UserProfile)
	// User Setting
	Setting := User.Group("/setting")
	Setting.PUT("/update-data", initialize.UserHandler.Update)
	Setting.PUT("/update-picture", initialize.UserHandler.ChangePictureByUser)
	Setting.PUT("/update-password", initialize.UserHandler.ChangePassword)
	Setting.PUT("/update-address", initialize.UserHandler.ChangeAddress)

	// Discussion

	// Discussion Like

	// Discussion picture

	// Comment

	// Comment Like

	// Comment Picture

	// Comment Reaction

	// Feedback

	// Superadmin & Admin
	admin := v1.Group("/admin", middleware.JWTWithConfig(initialize.AuthHandler), middlewares.AuthorizedUserAs(constant.RoleAdmin.String(), constant.RoleSuperadmin.String()))
	// User
	admin.GET("/users", initialize.UserHandler.GetAll)
	admin.PUT("/users/:user_id", initialize.UserHandler.Update)
	admin.DELETE("/users/:user_id", initialize.UserHandler.DeleteByAdmin)

	
}
