package users

import (
	"errors"
	"strings"

	"github.com/Zenk41/sipencari-rest-api/constant"
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/users"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/middlewares"
	"github.com/Zenk41/sipencari-rest-api/models"
	repository "github.com/Zenk41/sipencari-rest-api/repositories/users"
	"gorm.io/gorm"
)

type UserService interface {
	Create(payload payload.RegisterPayload, Role string) (response.User, error)
	Login(payload payload.LoginPayload) (response.Login, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ, Role string) (*gorm.DB, []response.User, error)
	GetByID(userID string) (response.User, error)
	GetByEmail(userEmail string) (bool, error)
	UpdateUser(payload payload.AccountPayload, userID string) (response.User, error)
	Delete(userID string) (bool, error)
	UpdatePicture(payload payload.ChangePicture, userID string) (response.User, error)
	UpdateAddress(payload payload.ChangeAddress, userID string) (response.User, error)
	ChangePassword(payload payload.ChangePasswordPayload, userID string) (response.User, error)
	CheckDuplicate(payload payload.CheckEmail) (bool, error)
}

type userService struct {
	repository repository.UserRepository
	jwtAuth    *middlewares.ConfigJWT
}

func NewUserService(repository repository.UserRepository, jwtAuth *middlewares.ConfigJWT) UserService {
	return &userService{repository: repository, jwtAuth: jwtAuth}
}

func (us *userService) Create(payload payload.RegisterPayload, Role string) (response.User, error) {

	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     constant.RoleUser,
	}

	if err := user.EncryptPassword(user.Password); err != nil {
		return response.User{}, err
	}

	user, err := us.repository.Create(user)
	if err != nil {
		return response.User{}, err
	}

	return *response.UserResponse(user), nil
}

func (us *userService) Login(payload payload.LoginPayload) (response.Login, error) {
	user, err := us.repository.GetByEmail(payload.Email)
	if err != nil {
		return response.Login{}, err
	}
	if err := user.CheckPassword(payload.Password); err != nil {
		return response.Login{}, err
	}

	token, err := us.jwtAuth.GenerateToken(user.UserID, user.Role.String())
	if err != nil {
		return response.Login{}, err
	}

	return response.Login{
		Token: token,
	}, nil
}

func (us *userService) GetAll(Page int, Size int, SortBy, Search, SearchQ, Role string) (*gorm.DB, []response.User, error) {
	var sort string
	var role string
	var roleQ string
	var search string
	var searchQ string
	if Role == "" {
		roleQ = ""
	} else if Role == constant.RoleAdmin.String() {
		roleQ = "role = ?"
		role = constant.RoleAdmin.String()
	} else if Role == constant.RoleUser.String() {
		roleQ = "role = ?"
		role = constant.RoleUser.String()
	} else if Role == constant.RoleSuperadmin.String() {
		roleQ = "role = ?"
		role = constant.RoleSuperadmin.String()
	}

	if SearchQ == "" {
		searchQ = ""
	} else {
		searchQ = SearchQ + " Like ? "
		search = "%" + Search + "%"
	}

	if SortBy != "" {
		if strings.HasPrefix(SortBy, "-") {
			sort = SortBy[1:] + " DESC"
		} else {
			sort = SortBy[0:] + " ASC"
		}
	} else {
		sort = ""
	}
	model, users, err := us.repository.GetAll(Page, Size, sort, search, searchQ, role, roleQ)
	if err != nil {
		return model, []response.User{}, err
	}

	return model, *response.UsersResponse(users), nil
}

func (us *userService) GetByID(userID string) (response.User, error) {

	user, err := us.repository.GetByID(userID)
	if err != nil {
		return response.User{}, err
	}
	return *response.UserResponse(user), nil
}

func (us *userService) GetByEmail(userEmail string) (bool, error) {
	user, err := us.repository.GetByEmail(userEmail)
	if err != nil || user.UserID == "" {
		return false, err
	}
	return true, nil
}

func (us *userService) UpdateUser(payload payload.AccountPayload, userID string) (response.User, error) {
	user, err := us.repository.GetByID(userID)
	if err != nil {
		return response.User{}, err
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}
	if payload.Email != "" {
		user.Email = payload.Email
	}
	if payload.Address != "" {
		user.Address = payload.Address
	}

	updatedUser, err := us.repository.Update(user)
	if err != nil {
		return response.User{}, err
	}
	return *response.UserResponse(updatedUser), nil
}

func (us *userService) Delete(userID string) (bool, error) {
	isDeleted, err := us.repository.Delete(userID)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}

func (us *userService) UpdatePicture(payload payload.ChangePicture, userID string) (response.User, error) {
	user, err := us.repository.GetByID(userID)
	if err != nil {
		return response.User{}, err
	}
	user.Picture = payload.Picture

	updatedUser, err := us.repository.Update(user)
	if err != nil {
		return response.User{}, err
	}
	return *response.UserResponse(updatedUser), nil
}

func (us *userService) UpdateAddress(payload payload.ChangeAddress, userID string) (response.User, error) {
	user, err := us.repository.GetByID(userID)
	if err != nil {
		return response.User{}, err
	}
	user.Address = payload.Address

	updatedUser, err := us.repository.Update(user)
	if err != nil {
		return response.User{}, err
	}
	return *response.UserResponse(updatedUser), nil
}

func (us *userService) ChangePassword(payload payload.ChangePasswordPayload, userID string) (response.User, error) {
	user, err := us.repository.GetByID(userID)
	if err != nil {
		return response.User{}, err
	}



	if payload.NewPassword == payload.OldPassword {
		return response.User{}, errors.New("Cant Use the same password")
	}

	if err := user.CheckPassword(payload.OldPassword); err != nil {
		return response.User{}, errors.New("Different old password")
	}

	if err := user.EncryptPassword(user.Password); err != nil {
		return response.User{}, err
	}

	updatedUser, err := us.repository.Update(user)

	if err != nil {
		return response.User{}, err
	}

	return *response.UserResponse(updatedUser), nil

}

func (us *userService) CheckDuplicate(payload payload.CheckEmail) (bool, error) {
	isDuplicate, err := us.repository.CheckDuplicate(payload.Email)

	if err != nil {
		return isDuplicate, err
	}
	return isDuplicate, nil
}
