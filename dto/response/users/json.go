package users

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/models"
)

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Picture   string    `json:"picture"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Token string `json:"token"`
}

func UserResponse(user models.User) *User {
	return &User{
		UserID:    user.UserID,
		Name:      user.Name,
		Role:      string(user.Role),
		Email:     user.Email,
		Picture:   user.Picture,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UsersResponse(users []models.User) *[]User {
	var usersResponse []User
	for _, user := range users {
		response := User{
			UserID:    user.UserID,
			Name:      user.Name,
			Role:      string(user.Role),
			Email:     user.Email,
			Picture:   user.Picture,
			Address:   user.Address,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersResponse = append(usersResponse, response)
	}
	return &usersResponse
}
