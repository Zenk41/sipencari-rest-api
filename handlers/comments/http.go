package comments

import (
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/comments"
	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comments"
	response "github.com/Zenk41/sipencari-rest-api/dto/response"
	// resComment "github.com/Zenk41/sipencari-rest-api/dto/response/comments"
)

type CommentHandler interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type commentHandler struct {
	validate *validator.Validate
	service  comments.CommentService
}

func NewCommentHandler(service comments.CommentService) CommentHandler {
	validate := validator.New()
	return &commentHandler{validate: validate, service: service}
}

func (ch *commentHandler) Create(c echo.Context) error {
	claim := middlewares.DecodeTokenClaims(c)
	discussionID := c.Param("discussion_id")
	var urlPictures []string
	var input payload.Comment
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := ch.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	form, err := c.MultipartForm()
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	files := form.File["comment_pictures"]
	if files != nil {
		isValid, Msg := util.IsFilesValid(files)
		if !isValid {
			return response.NewResponseFailed(c, http.StatusBadRequest, "failed", Msg, nil, "")
		}
		urlPictures, err = helper.MultipleUploadS3(c, files, "comment/")
		if err != nil {
			return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
		}
	}
	res, err := ch.service.Create(claim.ID, discussionID, urlPictures, input)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "comment created", res)

}

func (ch *commentHandler) GetAll(c echo.Context) error {
	discussionID := c.Param("discussion_id")
	res, err := ch.service.GetAll(discussionID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "all comment in discussion", res)
}

func (ch *commentHandler) GetByID(c echo.Context) error {
	commentID := c.Param("comment_id")

	res, err := ch.service.GetByID(commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "get comment", res)
}

func (ch *commentHandler) Update(c echo.Context) error {
	commentID := c.Param("comment_id")
	var input payload.UpdateComment
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := ch.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	res, err := ch.service.Update(commentID, input)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "comment updated", res)
}

func (ch *commentHandler) Delete(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	commentID := c.Param("comment_id")
	var result bool

	comment, err := ch.service.GetByID(commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	if comment.UserID != claims.ID {
		if claims.Role != constant.RoleAdmin.String() || claims.Role != constant.RoleSuperadmin.String() {
			return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
		} else {
			result, err = ch.service.Delete(commentID)
			if err != nil {
				return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
			}
		}
		return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
	} else {
		result, err = ch.service.Delete(commentID)
	}
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "comment deleted", nil)
}
