package initialize

import (
	"github.com/Zenk41/sipencari-rest-api/config"
	"github.com/Zenk41/sipencari-rest-api/db/sql"

	"github.com/Zenk41/sipencari-rest-api/middlewares"

	"github.com/Zenk41/sipencari-rest-api/handlers/greets"

	userH "github.com/Zenk41/sipencari-rest-api/handlers/users"
	userR "github.com/Zenk41/sipencari-rest-api/repositories/users"
	userS "github.com/Zenk41/sipencari-rest-api/services/users"

	disH "github.com/Zenk41/sipencari-rest-api/handlers/discussions"
	disR "github.com/Zenk41/sipencari-rest-api/repositories/discussions"
	disS "github.com/Zenk41/sipencari-rest-api/services/discussions"

	disLikeH "github.com/Zenk41/sipencari-rest-api/handlers/discussion_likes"
	disLikeR "github.com/Zenk41/sipencari-rest-api/repositories/discussion_likes"
	disLikeS "github.com/Zenk41/sipencari-rest-api/services/discussion_likes"

	disLocationH "github.com/Zenk41/sipencari-rest-api/handlers/discussion_locations"
	disLocationR "github.com/Zenk41/sipencari-rest-api/repositories/discussion_locations"
	disLocationS "github.com/Zenk41/sipencari-rest-api/services/discussion_locations"

	disPictureH "github.com/Zenk41/sipencari-rest-api/handlers/discussion_pictures"
	disPictureR "github.com/Zenk41/sipencari-rest-api/repositories/discussion_pictures"
	disPictureS "github.com/Zenk41/sipencari-rest-api/services/discussion_pictures"

	commentH "github.com/Zenk41/sipencari-rest-api/handlers/comments"
	commentR "github.com/Zenk41/sipencari-rest-api/repositories/comments"
	commentS "github.com/Zenk41/sipencari-rest-api/services/comments"

	comLikeH "github.com/Zenk41/sipencari-rest-api/handlers/comment_likes"
	comLikeR "github.com/Zenk41/sipencari-rest-api/repositories/comment_likes"
	comLikeS "github.com/Zenk41/sipencari-rest-api/services/comment_likes"

	comPictureH "github.com/Zenk41/sipencari-rest-api/handlers/comment_pictures"
	comPictureR "github.com/Zenk41/sipencari-rest-api/repositories/comment_pictures"
	comPictureS "github.com/Zenk41/sipencari-rest-api/services/comment_pictures"

	comReactionH "github.com/Zenk41/sipencari-rest-api/handlers/comment_reactions"
	comReactionR "github.com/Zenk41/sipencari-rest-api/repositories/comment_reactions"
	comReactionS "github.com/Zenk41/sipencari-rest-api/services/comment_reactions"

	feedbackH "github.com/Zenk41/sipencari-rest-api/handlers/feedbacks"
	feedbackR "github.com/Zenk41/sipencari-rest-api/repositories/feedbacks"
	feedbackS "github.com/Zenk41/sipencari-rest-api/services/feedbacks"

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
var userRepository userR.UserRepository
var userService userS.UserService
var UserHandler userH.UserHandler

// Discussion
var disRepository disR.DiscussionRepository
var disService disS.DiscussionService
var DisHandler disH.DiscussionHandler

// Discussion Like
var disLikeRepository disLikeR.DisLikeRepository
var disLikeService disLikeS.DisLikeService
var DisLikeHandler disLikeH.DisLikeHandler

// Discussion Location
var disLocationRepository disLocationR.DisLocationRepository
var disLocationService disLocationS.DisLocationService
var DisLocationHandler disLocationH.DisLocationHandler

// Discussion Picture
var disPictureRepository disPictureR.DisPictureRepository
var disPictureService disPictureS.DisPictureService
var DisPictureHandler disPictureH.DisPictureHandler

// Comment
var commentRepository commentR.CommentRepository
var commentService commentS.CommentService
var CommentHandler commentH.CommentHandler

// Comment Like
var comLikeRepository comLikeR.ComLikeRepository
var comLikeService comLikeS.ComLikeService
var ComLikeHandler comLikeH.ComLikeHandler

// Comment Picture
var comPictureRepository comPictureR.ComPictureRepository
var comPictureService comPictureS.ComPictureService
var ComPictureHandler comPictureH.ComPictureHandler

// Comment Reaction
var comReactionRepository comReactionR.ComReactionRepository
var comReactionService comReactionS.ComReactionService
var ComReactionHandler comReactionH.ComReactionHandler

// Feedback
var feedbackRepository feedbackR.FeedbackRepository
var feedbackService feedbackS.FeedbackService
var FeedbackHandler feedbackH.FeedbackHandler

func Init() *gorm.DB {
	dbSQL := sql.InitDB() // initialize sql database

	initRepositories(dbSQL) // initialize repostories

	initServices() // initialize services

	initHandlers() // initialize handlers

	sql.MigrationDB(dbSQL) // Migrating Table

	return dbSQL // returning db
}

func initRepositories(db *gorm.DB) {
	userRepository = userR.NewUserRepository(db)
	disRepository = disR.NewDiscussionRepository(db)
	disLikeRepository = disLikeR.NewDisLikeRepository(db)
	disLocationRepository = disLocationR.NewDisLocationRepository(db)
	disPictureRepository = disPictureR.NewDisPictureRepository(db)
	commentRepository = commentR.NewCommentRepository(db)
	comLikeRepository = comLikeR.NewComLikeRepository(db)
	comPictureRepository = comPictureR.NewComPictureRepository(db)
	comReactionRepository = comReactionR.NewComReactionRepository(db)
	feedbackRepository = feedbackR.NewFeedbackRepository(db)
}

func initServices() {
	AuthService = middlewares.ConfigJWT{
		SecretJWT:      config.LoadJWTConfig().JWT_SECRET_KEY,
		ExpireDuration: config.LoadJWTConfig().JWT_EXP_DURATION,
	}

	userService = userS.NewUserService(userRepository, &AuthService)

	disService = disS.NewDiscussionService(disRepository)

	disLikeService = disLikeS.NewDisLikeService(disLikeRepository)

	disLocationService = disLocationS.NewDisLocationService(disLocationRepository)

	disPictureService = disPictureS.NewDisPictureService(disPictureRepository)

	commentService = commentS.NewCommentService(commentRepository)

	comLikeService = comLikeS.NewComLikeService(comLikeRepository)

	comPictureService = comPictureS.NewComPictureService(comPictureRepository)

	comReactionService = comReactionS.NewComReactionService(comReactionRepository)

	feedbackService = feedbackS.NewFeedbackService(feedbackRepository)

}

func initHandlers() {
	// Logger
	LoggerHandler = middlewares.Logger()

	// JWT
	AuthHandler = AuthService.Init()

	GreetHandler = greets.NewGreetingHandler()

	UserHandler = userH.NewUserHandler(userService)

	DisHandler = disH.NewDiscussionHandler(disService)

	DisLikeHandler = disLikeH.NewDisLikeHandler(disLikeService)

	DisLocationHandler = disLocationH.NewDisLocationHandler(disLocationService)

	DisPictureHandler = disPictureH.NewDisPictureHandler(disPictureService)

	CommentHandler = commentH.NewCommentHandler(commentService)

	ComLikeHandler = comLikeH.NewComLikeHandler(comLikeService)

	ComPictureHandler = comPictureH.NewComPictureHandler(comPictureService)

	ComReactionHandler = comReactionH.NewComReactionHandler(comReactionService)

	FeedbackHandler = feedbackH.NewFeedbackHandler(feedbackService)

}
