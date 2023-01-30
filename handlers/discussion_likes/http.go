package discussion_likes

import (
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/discussion_likes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DisLikeHandler interface {
	GetAllLike(c echo.Context) error
	GetByID(c echo.Context) error
	Like(c echo.Context) error
}

type disLikeHandler struct {
	validate *validator.Validate
	service  discussion_likes.DisLikeService
}

func NewDisLikeHandler(service discussion_likes.DisLikeService) DisLikeHandler {
	validate := validator.New()
	return &disLikeHandler{validate: validate, service: service}
}

func (dlh disLikeHandler) GetAllLike(c echo.Context) error {
	discussionID := c.Param("discussion_id")

	res, err := dlh.service.GetAll(discussionID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "All Like in discussion retreived", res)
}

func (dlh disLikeHandler) GetByID(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	discussionID := c.Param("discussion_id")

	res, err := dlh.service.GetByID(discussionID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "Like in discussion retreived", res)

}
func (dlh disLikeHandler) Like(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	discussionID := c.Param("discussion_id")

	res, msg, err := dlh.service.Like(discussionID, claims.ID)
	if res == true {
		return response.NewResponseSuccess(c, http.StatusOK, "success", msg, nil)
	}
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", msg, nil)
}
