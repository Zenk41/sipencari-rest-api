package discussion_locations

import (
	"net/http"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussion_locations"
	"github.com/Zenk41/sipencari-rest-api/dto/response"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/services/discussion_locations"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DisLocationHandler interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	GetByDiscussionID(c echo.Context) error
	Update(c echo.Context) error
	UpdateByDiscussionID(c echo.Context) error
}

type disLocationHandler struct {
	validate *validator.Validate
	service  discussion_locations.DisLocationService
}

func NewDisLocationHandler(service discussion_locations.DisLocationService) DisLocationHandler {
	validate := validator.New()
	return &disLocationHandler{validate: validate, service: service}
}

func (dlh *disLocationHandler) GetAll(c echo.Context) error {
	res, err := dlh.service.GetAll()
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "success get all locations", res)
}
func (dlh *disLocationHandler) GetByID(c echo.Context) error {
	locationID := c.Param("location_id")
	res, err := dlh.service.GetByID(locationID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "success get location", res)
}
func (dlh *disLocationHandler) GetByDiscussionID(c echo.Context) error {
	discussionID := c.Param("discussion_id")
	res, err := dlh.service.GetByDiscussionID(discussionID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "success get location", res)
}

func (dlh *disLocationHandler) Update(c echo.Context) error {
	locationID := c.Param("location_id")
	var input payload.DiscussionLocation

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := dlh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	locationName, err := helper.GetAdressFromLatLng(c, input.Lat, input.Lng)

	res, err := dlh.service.Update(input, locationID, locationName.FormatedAddress)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "lcoation updated", res)
}

func (dlh *disLocationHandler) UpdateByDiscussionID(c echo.Context) error {
	discussionID := c.Param("discussion_id")
	var input payload.DiscussionLocation

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := dlh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	locationName, err := helper.GetAdressFromLatLng(c, input.Lat, input.Lng)

	res, err := dlh.service.UpdateByDiscussionID(input, discussionID, locationName.FormatedAddress)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "lcoation updated", res)
}
