package users

import (
	"errors"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

type UserRepository interface {
	Create(User models.User) (models.User, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ, Role, RoleQ string) (*gorm.DB, []models.User, error)
	GetByID(UserID string) (models.User, error)
	GetByEmail(UserEmail string) (models.User, error)
	Update(User models.User, userID string) (models.User, error)
	Delete(UserID string) (bool, error)
	CheckDuplicate(Email string) (bool, error)
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Create(User models.User) (models.User, error) {
	User.SetId(ur.conn)
	err := ur.conn.Create(&User).Error
	return User, err
}

func (ur *userRepository) GetAll(Page int, Size int, SortBy, Search, SearchQ, Role, RoleQ string) (*gorm.DB, []models.User, error) {
	var rec []models.User
	var model *gorm.DB
	if SortBy == "" && SearchQ == "" && RoleQ == "" {
		model = ur.conn.Model(&rec)
	} else if SortBy != "" && SearchQ == "" && RoleQ == "" {
		model = ur.conn.Model(&rec).Order(SortBy)
	} else if SortBy == "" && SearchQ != "" && RoleQ == "" {
		model = ur.conn.Model(&rec).Where(SearchQ, Search)
	} else if SortBy == "" && SearchQ == "" && RoleQ != "" {
		model = ur.conn.Model(&rec).Where(RoleQ, Role)
	} else if SortBy != "" && SearchQ != "" && RoleQ == "" {
		model = ur.conn.Model(&rec).Order(SortBy).Where(SearchQ, Search)
	} else {
		model = ur.conn.Model(&rec).Order(SortBy).Where(RoleQ, Role).Where(SearchQ, Search).Where(RoleQ, Role)
	}

	if err := model.Find(&rec); err != nil {
		return model, []models.User{}, err.Error
	}

	return model, rec, nil
}

func (ur *userRepository) GetByID(UserID string) (models.User, error) {
	var rec models.User
	error := ur.conn.Where("users.user_id = ?", UserID).First(&rec).Error
	return rec, error
}

func (ur *userRepository) GetByEmail(UserEmail string) (models.User, error) {
	var rec models.User
	error := ur.conn.Where("users.email = ?", UserEmail).First(&rec).Error
	return rec, error
}

func (ur *userRepository) Update(User models.User, userID string) (models.User, error) {
	err := ur.conn.Updates(User).Error
	return User, err
}

func (ur *userRepository) Delete(UserID string) (bool, error) {
	rec, err := ur.GetByID(UserID)
	if err != nil {
		return false, err
	}
	if rec.Role == constant.RoleSuperadmin {
		return false, errors.New("Forbidden")
	}
	if result := ur.conn.Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}

func (ur *userRepository) CheckDuplicate(Email string) (bool, error) {
	var rec []models.User

	err := ur.conn.Where("users.email = ?", Email).First(&rec).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, err
		}
		return false, err
	}
	return true, nil
}
