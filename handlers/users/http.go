package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/helper"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/services/users"
	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/users"
	response "github.com/Zenk41/sipencari-rest-api/dto/response"
)

type UserHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	ChangePictureByUser(c echo.Context) error
	ChangePassword(c echo.Context) error
	ChangeAddress(c echo.Context) error
	MyProfile(c echo.Context) error
	UserProfile(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	UpdateByAdmin(c echo.Context) error
	DeleteByAdmin(c echo.Context) error
}

type userHandler struct {
	validate *validator.Validate
	service  users.UserService
}

func NewUserHandler(service users.UserService) UserHandler {
	validate := validator.New()
	return &userHandler{validate: validate, service: service}
}

func (uh *userHandler) Register(c echo.Context) error {
	var input payload.RegisterPayload

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	emailExist, _ := uh.service.GetByEmail(input.Email)
	if emailExist {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "Email Already Exist", nil, "Email Already Exist")
	}

	userResponse, err := uh.service.Create(input, constant.RoleUser.String())
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "user registered", userResponse)
}

func (uh *userHandler) Login(c echo.Context) error {
	var input payload.LoginPayload

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	loginResponse, err := uh.service.Login(input)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "login success", loginResponse)
}

func (uh *userHandler) ChangePictureByUser(c echo.Context) error {
	var result string
	var input payload.ChangePicture
	picture, _ := c.FormFile("picture")
	user := middlewares.DecodeTokenClaims(c)
	if user.ID == "" {
		return response.NewResponseFailed(c, http.StatusUnauthorized, "failed", "unauthorized", nil, "")
	}

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if picture != nil {
		isValid, message := util.IsFileValid(picture)

		if !isValid {
			return response.NewResponseFailed(c, http.StatusBadRequest, "failed", message, nil, message)
		}
		picture.Filename = time.Now().String() + ".png"
		src, _ := picture.Open()
		defer src.Close()
		result, _ = helper.UploadToS3(c, "profile/", picture.Filename, src)
		input.Picture = result
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	userResponse, err := uh.service.UpdatePicture(input, user.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}

	return response.NewResponseSuccess(c, http.StatusOK, "success", "picture has changed", userResponse)
}

func (uh *userHandler) ChangePassword(c echo.Context) error {
	var input payload.ChangePasswordPayload

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	claims := middlewares.DecodeTokenClaims(c)

	res, err := uh.service.ChangePassword(input, claims.ID)

	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "password changed", res)
}

func (uh *userHandler) ChangeAddress(c echo.Context) error {
	var input payload.ChangeAddress

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	claims := middlewares.DecodeTokenClaims(c)

	res, err := uh.service.UpdateAddress(input, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "address changed", res)
}

func (uh *userHandler) MyProfile(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)

	res, err := uh.service.GetByID(claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "my profile", res)
}

func (uh *userHandler) UserProfile(c echo.Context) error {
	idUser := c.Param("user_id")
	res, err := uh.service.GetByID(idUser)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "my profile", res)
}

func (uh *userHandler) Update(c echo.Context) error {
	claims := middlewares.DecodeTokenClaims(c)
	var input payload.AccountPayload
	emailExist := false

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	if input.Email != "" {
		emailExist, _ = uh.service.GetByEmail(input.Email)
	}

	if emailExist {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "Email Already Exist", nil, "Email Already Exist")
	}

	res, err := uh.service.UpdateUser(input, claims.ID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "data updated", &res)
}

func (uh *userHandler) GetAll(c echo.Context) error {
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	sortBy := c.QueryParam("sort")
	search := c.QueryParam("search")
	searchQ := c.QueryParam("search-q")
	role := c.QueryParam("role")

	model, res, err := uh.service.GetAll(page, size, sortBy, search, searchQ, role)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "get users failed", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "success get users", pg.Response(model, c.Request(), &res))
}

func (uh *userHandler) UpdateByAdmin(c echo.Context) error {
	userID := c.Param("user_id")
	var input payload.AccountPayload
	emailExist := false

	if err := c.Bind(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "invalid request", nil, err.Error())
	}

	if err := uh.validate.Struct(&input); err != nil {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "validation failed", nil, err.Error())
	}

	if input.Email != "" {
		emailExist, _ = uh.service.GetByEmail(input.Email)
	}
	if emailExist {
		return response.NewResponseFailed(c, http.StatusBadRequest, "failed", "Email Already Exist", nil, "Email Already Exist")
	}

	res, err := uh.service.UpdateUser(input, userID)
	if err != nil {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "data updated", &res)
}
func (uh *userHandler) DeleteByAdmin(c echo.Context) error {
	userID := c.Param("user_id")
	result, err := uh.service.Delete(userID)
	if err != nil || !result {
		return response.NewResponseFailed(c, http.StatusInternalServerError, "failed", "internal server error", nil, err.Error())
	}
	return response.NewResponseSuccess(c, http.StatusOK, "success", "user deleted", nil)

}
