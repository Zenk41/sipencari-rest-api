package comment_likes

import (
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/comment_likes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ComLikeHandler interface {
	GetAllLike(c echo.Context) error
	GetByID(c echo.Context) error
	Like(c echo.Context) error
}

type comLikeHandler struct {
	validate *validator.Validate
	service  comment_likes.ComLikeService
}

func NewComLikeHandler(service comment_likes.ComLikeService) ComLikeHandler {
	validate := validator.New()
	return &comLikeHandler{validate: validate, service: service}
}

func (clh *comLikeHandler) GetAllLike(c echo.Context) error {
	commentID := c.Param("comment_id")

	res, err := clh.service.GetAll(commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "All Like in comment retreived", res)
}

func (clh *comLikeHandler) GetByID(c echo.Context) error {
	userID := c.Param("user_id")
	commentID := c.Param("comment_id")

	res, err := clh.service.GetByID(commentID, userID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "Like in comment retreived", res)

}

func (clh *comLikeHandler) Like(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	commentID := c.Param("comment_id")

	res, msg, err := clh.service.Like(commentID, claims.ID)
	if res == true {
		return response.NewResponseSuccess(c, http.StatusOK, "success", msg, nil)
	}
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", msg, nil)
}
