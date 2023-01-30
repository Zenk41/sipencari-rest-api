package comment_reactions

import (
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/comment_reactions"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comment_reactions"
	response "github.com/Zenk41/sipencari-rest-api/dto/response"
)

type ComReactionHandler interface {
	React(c echo.Context) error
	GetByID(c echo.Context) error
	GetAll(c echo.Context) error
}

type comReactionHandler struct {
	validate *validator.Validate
	service  comment_reactions.ComReactionService
}

func NewComReactionHandler(service comment_reactions.ComReactionService) ComReactionHandler {
	validate := validator.New()
	return &comReactionHandler{validate: validate, service: service}
}

func (crh *comReactionHandler) React(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	var input payload.CommentReaction
	commentID := c.Param("comment_id")
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := crh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	_, res, err := crh.service.React(input, commentID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "reaction success", res)
}

func (crh *comReactionHandler) GetByID(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	commentID := c.Param("comment_id")
	res, err := crh.service.GetByID(commentID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "reaction success", res)
}
func (crh *comReactionHandler) GetAll(c echo.Context) error {
	commentID := c.Param("comment_id")
	res, err := crh.service.GetAll(commentID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "reaction success", res)
}
