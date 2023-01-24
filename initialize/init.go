package initialize

import (
	"github.com/Zenk41/sipencari-rest-api/config"
	"github.com/Zenk41/sipencari-rest-api/db/sql"

	"github.com/Zenk41/sipencari-rest-api/middlewares"

	"github.com/Zenk41/sipencari-rest-api/handlers/greets"

	usersH "github.com/Zenk41/sipencari-rest-api/handlers/users"
	usersR "github.com/Zenk41/sipencari-rest-api/repositories/users"
	usersS "github.com/Zenk41/sipencari-rest-api/services/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/gorm"
)

// Logger
var LoggerHandler echo.MiddlewareFunc

// Auth
var AuthService middlewares.ConfigJWT
var AuthHandler middleware.JWTConfig

// Greets
var GreetHandler greets.GreetingHandler

// User
var userRepository usersR.UserRepository
var userService usersS.UserService
var UserHandler usersH.UserHandler

func Init() *gorm.DB {
	dbSQL := sql.InitDB() // initialize sql database

	initRepositories(dbSQL) // initialize repostories

	initServices() // initialize services

	initHandlers() // initialize handlers

	sql.MigrationDB(dbSQL) // Migrating Table

	return dbSQL // returning db
}

func initRepositories(db *gorm.DB) {
	userRepository = usersR.NewUserRepository(db)
}

func initServices() {
	AuthService = middlewares.ConfigJWT{
		SecretJWT:      config.LoadJWTConfig().JWT_SECRET_KEY,
		ExpireDuration: config.LoadJWTConfig().JWT_EXP_DURATION,
	}
	userService = usersS.NewUserService(userRepository, &AuthService)
}

func initHandlers() {
	// Logger
	LoggerHandler = middlewares.Logger()

	// JWT
	AuthHandler = AuthService.Init()

	GreetHandler = greets.NewGreetingHandler()

	UserHandler = usersH.NewUserHandler(userService)

}
