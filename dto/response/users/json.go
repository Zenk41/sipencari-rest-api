package users

import (
	"github.com/Zenk41/sipencari-rest-api/models"

)

type User struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type Login struct {
	Token string `json:"token"`
}

func UserResponse(user models.User) *User {
	return &User{
		UserID:  user.UserID,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
	}
}

func UsersResponse(users []models.User) *[]User {
	var usersResponse []User
	for _, user := range users {
		response := User{
			UserID:  user.UserID,
			Name:    user.Name,
			Email:   user.Email,
			Picture: user.Picture,
		}
		usersResponse = append(usersResponse, response)
	}
	return &usersResponse
}
