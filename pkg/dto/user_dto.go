package dto

import "github.com/valeelim/mahchat/pkg/dao"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

type RegisterResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	UserId      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type GetUsersResponse struct {
	Data []dao.User `json:"data,omitempty"`
}

func NewRegisterRequest(email, password, name string) *RegisterRequest {
	return &RegisterRequest{
		Email:    email,
		Password: password,
		Name:     name,
	}
}

func NewLoginRequest(email, password string) *LoginRequest {
	return &LoginRequest{
		Email:    email,
		Password: password,
	}
}
