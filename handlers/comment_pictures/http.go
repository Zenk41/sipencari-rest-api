package comment_pictures

import (
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comment_pictures"
	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/services/comment_pictures"
	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type ComPictureHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Delete(c echo.Context) error
}

type comPictureHandler struct {
	validate *validator.Validate
	service  comment_pictures.ComPictureService
}

func NewComPictureHandler(service comment_pictures.ComPictureService) ComPictureHandler {
	validate := validator.New()
	return &comPictureHandler{validate: validate, service: service}
}

func (cph *comPictureHandler) Create(c echo.Context) error {
	var urlPictures []string
	commentID := c.Param("comment_id")
	form, err := c.MultipartForm()
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

	res, err := cph.service.Create(urlPictures, commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture created ", res)
}
func (cph *comPictureHandler) Update(c echo.Context) error {
	comPictureID := c.Param("picture_id")
	var result string
	var input payload.CommentPicture
	picture, _ := c.FormFile("url")

	if picture != nil {
		isValid, message := util.IsFileValid(picture)

		if !isValid {
			return response.NewResponseFailed(c, http.StatusBadRequest, "failed", message, nil, message)
		}
		picture.Filename = time.Now().String() + ".png"
		src, _ := picture.Open()
		defer src.Close()
		result, _ = helper.UploadToS3(c, "comment/", picture.Filename, src)
		input.URL = result
	}
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := cph.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	res, err := cph.service.UpdateComPicture(input, comPictureID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture updated ", res)
}
func (cph *comPictureHandler) GetAll(c echo.Context) error {
	commentID := c.Param("comment_id")
	res, err := cph.service.GetAll(commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "All Picture in comment retreived", res)
}
func (cph *comPictureHandler) GetByID(c echo.Context) error {
	comPictureID := c.Param("picture_id")

	res, err := cph.service.GetByID(comPictureID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture retreived by user", res)
}
func (cph *comPictureHandler) Delete(c echo.Context) error {
	comPictureID := c.Param("picture_id")

	result, err := cph.service.Delete(comPictureID)
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture deleted", nil)
}
