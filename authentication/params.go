package authentication

import "encoding/json"

type LoginParameter struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterParameter struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgetPasswordParameter struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordParameter struct {
	Password string `json:"password" binding:"required"`
}

type UpdateUserParamater struct {
	ID       json.Number `json:"id" binding:"required"`
	Fullname string      `json:"fullname" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required"`
}

type DeleteUserParameter struct {
	ID json.Number `json:"id" binding:"required"`
}
