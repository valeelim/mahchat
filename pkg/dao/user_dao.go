package dao

import (
	"time"

	"github.com/valeelim/mahchat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64    `json:"id,omitempty"`
	Email          string    `json:"email,omitempty"`
	Name           string    `json:"name,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func NewUser(email, name string) (*User, error) {
	id, err := utils.GenerateSnowflake() 
	if err != nil {
		return nil, err
	}
	return &User{
		ID: id,
		Email: email,
		Name:  name,
	}, nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err
}
