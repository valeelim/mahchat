package repository

import (
	"github.com/valeelim/mahchat/pkg/dao"
)

type User interface {
	CreateUser(user *dao.User) error
	GetUserById(id string) (dao.User, error)
	GetUserByEmail(email string) (dao.User, error)
	GetAllUsers() ([]dao.User, error)
}


