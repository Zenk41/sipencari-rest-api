package discussions

import (
	"net/http"
	"strconv"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/discussions"
	"github.com/Zenk41/sipencari-rest-api/util"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussions"
	response "github.com/Zenk41/sipencari-rest-api/dto/response"
	resDiscussion "github.com/Zenk41/sipencari-rest-api/dto/response/discussions"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
)

type DiscussionHandler interface {
	CreateDiscussion(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	GetByUserID(e echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetMine(c echo.Context) error
}

type discussionHandler struct {
	validate *validator.Validate
	service  discussions.DiscussionService
}

func NewDiscussionHandler(service discussions.DiscussionService) DiscussionHandler {
	validate := validator.New()
	return &discussionHandler{validate: validate, service: service}
}

func (dh *discussionHandler) CreateDiscussion(c echo.Context) error {
	user := middlewares.DecodeTokenClaims(c)
	var urlPictures []string
	var input payload.CreateDiscussion
	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}
	if err := dh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}
	form, err := c.MultipartForm()
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
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

	locationName, err := helper.GetAdressFromLatLng(c, input.Lat, input.Lng)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	dis, err := dh.service.Create(input, user.ID, urlPictures, locationName.FormatedAddress, user.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture has changed", dis)
}

func (dh *discussionHandler) GetMine(c echo.Context) error {
	claim := middlewares.DecodeTokenClaims(c)

	res, err := dh.service.GetMyDiscussion(claim.ID, claim.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "discussion retreived by user", &res)
}

func (dh *discussionHandler) GetAll(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	sortBy := c.QueryParam("sort")
	search := c.QueryParam("search")
	searchQ := c.QueryParam("search-q")
	status := c.QueryParam("status")
	privacy := c.QueryParam("privacy")

	model, res, err := dh.service.GetAll(page, size, sortBy, status, privacy, search, searchQ, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "get users failed", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "All discussion retieved", pg.Response(model, c.Request(), &res))
}

func (dh *discussionHandler) GetByID(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	discussionID := c.Param("discussion_id")
	res, err := dh.service.GetByID(discussionID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "discussion retieved", res)
}

func (dh *discussionHandler) GetByUserID(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	userID := c.Param("user_id")
	privacy := c.QueryParam("privacy")

	res, err := dh.service.GetByUserID(userID, privacy, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "discussion retreived by user", res)
}

func (dh *discussionHandler) Update(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	var input payload.UpdateDiscussion
	var res resDiscussion.Discussion
	discussionID := c.Param("discussion_id")

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := dh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	discussion, err := dh.service.GetByID(discussionID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	locationName, err := helper.GetAdressFromLatLng(c, input.Lat, input.Lng)

	if discussion.UserID != claims.ID {
		if claims.Role == constant.RoleAdmin.String() || claims.Role == constant.RoleSuperadmin.String() {
			res, err = dh.service.Update(input, discussionID, locationName.FormatedAddress, claims.ID)

		} else {
			print("role"+claims.Role);
			print("role constant"+constant.RoleAdmin.String())
			return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
		}
		// return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
	} else {
		res, err = dh.service.Update(input, discussionID, locationName.FormatedAddress, claims.ID)
	}
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}


	return response.NewResponseSuccess(c, http.StatusOK, "success", "data updated", &res)
}

func (dh *discussionHandler) Delete(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	discussionID := c.Param("discussion_id")
	var result bool

	discussion, err := dh.service.GetByID(discussionID, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	if discussion.UserID != claims.ID {
		if claims.Role == constant.RoleAdmin.String() || claims.Role == constant.RoleSuperadmin.String() {
			result, err = dh.service.Delete(discussionID)
		} else {
			return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
		}
		// return response.NewResponseFailed(c, http.StatusForbidden, "failed", "user doesnt have access", nil, "")
	} else {
		result, err = dh.service.Delete(discussionID)
	}
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "discussion deleted", nil)
}
