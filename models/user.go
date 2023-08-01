package models

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID    string            `json:"user_id" gorm:"size:255;primaryKey"`
	Name      string            `json:"name" gorm:"size:100;"`
	Email     string            `json:"email" gorm:"size:255;unique"`
	Password  string            `json:"password" gorm:"size:16;"`
	Picture   string            `json:"picture" gorm:"size:255;"`
	Role      constant.RoleEnum `json:"role" gorm:"size:10;"`
	Address   string            `json:"address" gorm:"size:255;"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

func (u *User) SetId(db *gorm.DB) {
	u.UserID = uuid.New().String()
}

func (u *User) EncryptPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pass)
	return nil
}

func (u *User) CheckPassword(encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(encryptedPassword))
	if err != nil {
		return err
	}
	return nil
}
