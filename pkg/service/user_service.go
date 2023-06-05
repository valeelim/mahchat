package service

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valeelim/mahchat/pkg/dao"
	"github.com/valeelim/mahchat/pkg/dto"
	"github.com/valeelim/mahchat/pkg/repository"
	"github.com/valeelim/mahchat/pkg/utils"
)

type UserService struct {
}

func New() *UserService {
	return new(UserService)
}

type RegisterUser func(c *gin.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)

type LoginUser func(c *gin.Context, req dto.LoginRequest) (*dto.LoginResponse, error)

type GetUsers func(c *gin.Context) (*dto.GetUsersResponse, error)

func RegisterUserService(repo repository.User) RegisterUser {
	return func(c *gin.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
		if _, err := repo.GetUserByEmail(req.Email); err == nil {
			return nil, err
		}
		user := dao.NewUser(req.Email, req.Name)
		if err := user.SetPassword(req.Password); err != nil {
			return nil, err
		}
		if err := repo.CreateUser(user); err != nil {
			return nil, err
		}
		return &dto.RegisterResponse{
			Email: req.Email,
			Name: req.Name,
		}, nil
	}
}

func LoginUserService(repo repository.User, cache repository.Cache) LoginUser {
	return func(c *gin.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
		user, err := repo.GetUserByEmail(req.Email)
		if err != nil {
			return nil, err
		}
		if err := user.ComparePassword(req.Password); err != nil {
			return nil, err
		}

		tokenLength := 32
		token, err := utils.GenerateOpaqueToken(tokenLength)
		if err != nil {
			return nil, err
		}

		if err := cache.SetAccessToken(context.Background(), token, map[string]interface{}{
			"user_id": user.ID,
			"access_token": token,
			"expires_at": time.Now().Add(1 * time.Hour).Unix(),
		}); err != nil {
			return nil, err
		}
		return &dto.LoginResponse{
			UserId: user.ID,
			AccessToken: token,
		}, nil
	}
}

func GetAllUsersService(repo repository.User) GetUsers {
	return func(c *gin.Context) (*dto.GetUsersResponse, error) {
		result, err := repo.GetAllUsers()
		if err != nil {
			return nil, err
		}
		return &dto.GetUsersResponse{Data: result}, nil
	}
}
