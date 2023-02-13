package users

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,min=4,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AccountPayload struct {
	Name    string `json:"name" validate:"min=4,max=100"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address" validate:"min=10,max=50"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type ResetPasswordPayload struct {
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type ChangePicture struct {
	Picture string `json:"picture" form:"picture" validate:"required"`
}

type ChangeAddress struct {
	Address string `json:"address" validate:"required,min=10,max=50"`
}

type CheckEmail struct {
	Email string `json:"email" validate:"required,email"`
}
