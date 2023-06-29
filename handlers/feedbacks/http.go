package feedbacks

import (
	"net/http"
	"strconv"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/feedbacks"
	"github.com/morkid/paginate"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/feedbacks"
	response "github.com/Zenk41/sipencari-rest-api/dto/response"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler interface {
	CreateFeedback(c echo.Context) error
	DeleteFeedback(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
}

type feedbackHandler struct {
	validate *validator.Validate
	service  feedbacks.FeedbackService
}

func NewFeedbackHandler(service feedbacks.FeedbackService) FeedbackHandler {
	validate := validator.New()
	return &feedbackHandler{validate: validate, service: service}
}

func (fdh *feedbackHandler) GetAll(c echo.Context) error {
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	sortBy := c.QueryParam("sort")
	search := c.QueryParam("search")
	searchQ := c.QueryParam("search-q")

	model, res, err := fdh.service.GetAll(page, size, sortBy, search, searchQ)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "success get all feedback", pg.Response(model, c.Request(), &res))
}

func (fdh *feedbackHandler) GetByID(c echo.Context) error {
	feedbackID := c.Param("feedback_id")
	res, err := fdh.service.GetByID(feedbackID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "get feedback", res)
}

func (fdh *feedbackHandler) Update(c echo.Context) error {
	feedbackID := c.Param("feedback_id")
	var input payload.Feedback
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := fdh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	res, err := fdh.service.UpdateFeedback(input, feedbackID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "feedback updated", res)
}

func (fdh *feedbackHandler) CreateFeedback(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	var input payload.Feedback
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := fdh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	res, err := fdh.service.Create(input, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "feedback created", res)
}

func (fdh *feedbackHandler) UpdateFeedback(c echo.Context) error {
	feedbackID := c.Param("feedback_id")
	var input payload.Feedback
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := fdh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	res, err := fdh.service.UpdateFeedback(input, feedbackID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "feedback updated", res)
}

func (fdh *feedbackHandler) DeleteFeedback(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	feedbackID := c.Param("feedback_id")
	var result bool
	discussion, err := fdh.service.GetByID(feedbackID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	if discussion.UserID != claims.ID {
		if claims.Role == constant.RoleAdmin.String() || claims.Role == constant.RoleSuperadmin.String() {
			result, err = fdh.service.Delete(feedbackID)
			
		} else {
			return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
		}
		// return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
	} else {
		result, err = fdh.service.Delete(feedbackID)
	}
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "feedback deleted", nil)
}
