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
	Discussion := v1.Group("/discussions", middleware.JWTWithConfig(initialize.AuthHandler))
	Discussion.POST("", initialize.DisHandler.CreateDiscussion)
	Discussion.GET("", initialize.DisHandler.GetAll)
	Discussion.GET("/:discussion_id", initialize.DisHandler.GetByID)
	Discussion.PUT("/:discussion_id", initialize.DisHandler.Update)
	Discussion.DELETE("/:discussion_id", initialize.DisHandler.Delete)
	User.GET("/:user_id/discussions", initialize.DisHandler.GetByUserID)

	// Discussion Like
	DiscussionLike := Discussion.Group("/:discussion_id/likes")
	DiscussionLike.GET("", initialize.DisLikeHandler.GetAllLike)
	DiscussionLike.GET("/user", initialize.DisLikeHandler.GetByID)
	DiscussionLike.POST("", initialize.DisLikeHandler.Like)

	// Discussion picture

	DiscussionPicture := Discussion.Group("/:discussion_id/pictures")
	DiscussionPicture.GET("", initialize.DisPictureHandler.GetAll)
	DiscussionPicture.POST("", initialize.DisPictureHandler.Create)
	Picture := v1.Group("/discussions/pictures/:picture_id", middleware.JWTWithConfig(initialize.AuthHandler))
	Picture.GET("", initialize.DisPictureHandler.GetByID)
	Picture.DELETE("", initialize.DisPictureHandler.Delete)
	Picture.PUT("", initialize.DisPictureHandler.Update)

	// Discussion Location
	DiscussionLocation := Discussion.Group("/:discussion_id/location")
	DiscussionLocation.GET("", initialize.DisLocationHandler.GetByDiscussionID)
	DiscussionLocation.PUT("", initialize.DisLocationHandler.UpdateByDiscussionID)
	DisLocation := v1.Group("/locations", middleware.JWTWithConfig(initialize.AuthHandler))
	DisLocation.GET("", initialize.DisLocationHandler.GetAll)
	DisLocation.GET("/:location_id", initialize.DisLocationHandler.GetByID)
	DisLocation.PUT("/:location_id", initialize.DisLocationHandler.Update)

	// Comment
	DiscussionComment := Discussion.Group("/:discussion_id/comments")
	DiscussionComment.GET("", initialize.CommentHandler.GetAll)
	DiscussionComment.POST("", initialize.CommentHandler.Create)
	Comment := v1.Group("/comments/:comment_id", middleware.JWTWithConfig(initialize.AuthHandler))
	Comment.GET("", initialize.CommentHandler.GetByID)
	Comment.PUT("", initialize.CommentHandler.Update)
	Comment.DELETE("", initialize.CommentHandler.Delete)

	// Comment Like
	CommentLike := Comment.Group("/likes")
	CommentLike.GET("", initialize.ComLikeHandler.GetAllLike)
	CommentLike.GET("/:user_id", initialize.ComLikeHandler.GetByID)
	CommentLike.POST("", initialize.ComLikeHandler.Like)

	// Comment Picture
	CommentPicture := Comment.Group("/pictures")
	CommentPicture.GET("", initialize.ComPictureHandler.GetAll)
	CommentPicture.POST("", initialize.ComPictureHandler.Create)
	CPicture := v1.Group("/comments/pictures/:picture_id")
	CPicture.DELETE("", initialize.ComPictureHandler.Delete)
	CPicture.GET("", initialize.ComPictureHandler.GetByID)
	CPicture.PUT("", initialize.ComPictureHandler.Update)

	// Comment Reaction
	CommentReaction := Comment.Group("/reactions")
	CommentReaction.GET("", initialize.ComReactionHandler.GetAll)
	CommentReaction.POST("", initialize.ComReactionHandler.React)
	CommentReaction.GET("/:user_id", initialize.ComReactionHandler.GetByID)

	// Feedback
	Feedback := v1.Group("/feedbacks", middleware.JWTWithConfig(initialize.AuthHandler))
	Feedback.POST("", initialize.FeedbackHandler.CreateFeedback)
	Feedback.GET("", initialize.FeedbackHandler.GetAll)
	Feedback.PUT("/:feedback_id", initialize.FeedbackHandler.Update)
	Feedback.GET("/:feedback_id", initialize.FeedbackHandler.GetByID)
	Feedback.DELETE("/:feedback_id", initialize.FeedbackHandler.DeleteFeedback)

	// Superadmin & Admin
	admin := v1.Group("/admin", middleware.JWTWithConfig(initialize.AuthHandler), middlewares.AuthorizedUserAs(constant.RoleAdmin.String(), constant.RoleSuperadmin.String()))
	// User
	adminUser := admin.Group("/users")
	adminUser.GET("", initialize.UserHandler.GetAll)
	adminUser.PUT("/:user_id", initialize.UserHandler.Update)
	adminUser.DELETE(":user_id", initialize.UserHandler.DeleteByAdmin)
	adminUser.GET(":user_id/discussions", initialize.DisHandler.GetByUserID)

	adminDiscussion := admin.Group("/discussions")
	adminDiscussion.GET("", initialize.DisHandler.GetAll)
	adminDiscussion.PUT("/:discussion_id", initialize.DisHandler.Update)
	adminDiscussion.GET("/:discussion_id", initialize.DisHandler.GetByID)
	adminDiscussion.DELETE("/:discussion_id", initialize.DisHandler.Delete)

}
