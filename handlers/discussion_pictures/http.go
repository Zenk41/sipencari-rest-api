package discussion_pictures

import (
	"net/http"
	"time"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussion_pictures"
	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/services/discussion_pictures"
	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DisPictureHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Delete(c echo.Context) error
}

type disPictureHandler struct {
	validate *validator.Validate
	service  discussion_pictures.DisPictureService
}

func NewDisPictureHandler(service discussion_pictures.DisPictureService) DisPictureHandler {
	validate := validator.New()
	return &disPictureHandler{validate: validate, service: service}
}

func (dph *disPictureHandler) Create(c echo.Context) error {
	var urlPictures []string
	discussionID := c.Param("discussion_id")
	form, err := c.MultipartForm()
	files := form.File["discussion_pictures"]
	if files != nil {
		isValid, Msg := util.IsFilesValid(files)
		if !isValid {
			return response.NewResponseFailed(c, http.StatusBadRequest, "failed", Msg, nil, "")
		}
		urlPictures, err = helper.MultipleUploadS3(c, files, "discussion/")
		if err != nil {
			return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
		}
	}

	res, err := dph.service.Create(urlPictures, discussionID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture created ", res)
}

func (dph *disPictureHandler) Update(c echo.Context) error {
	disPictureID := c.Param("picture_id")
	var result string
	var input payload.DiscussionPicture
	picture, _ := c.FormFile("url")

	if picture != nil {
		isValid, message := util.IsFileValid(picture)

		if !isValid {
			return response.NewResponseFailed(c, http.StatusBadRequest, "failed", message, nil, message)
		}
		picture.Filename = time.Now().String() + ".png"
		src, _ := picture.Open()
		defer src.Close()
		result, _ = helper.UploadToS3(c, "discussion/", picture.Filename, src)
		input.URL = result
	}
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := dph.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	res, err := dph.service.UpdateDisPicture(input, disPictureID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture updated ", res)
}

func (dph *disPictureHandler) GetAll(c echo.Context) error {
	discussionID := c.Param("discussion_id")
	res, err := dph.service.GetAll(discussionID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "All Picture in discussion retreived", res)
}
func (dph *disPictureHandler) GetByID(c echo.Context) error {
	disPictureID := c.Param("picture_id")

	res, err := dph.service.GetByID(disPictureID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture retreived by user", res)
}

func (dph *disPictureHandler) Delete(c echo.Context) error {
	disPictureID := c.Param("picture_id")

	result, err := dph.service.Delete(disPictureID)
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture deleted", nil)
}
