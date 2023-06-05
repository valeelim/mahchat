package dao

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string    `json:"id,omitempty"`
	Email          string    `json:"email,omitempty"`
	Name           string    `json:"name,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func NewUser(email, name string) *User {
	return &User{
		Email: email,
		Name:  name,
	}
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
