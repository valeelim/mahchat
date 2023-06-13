package database

import (
	"database/sql"
	"errors"

	"github.com/valeelim/mahchat/pkg/dao"
)

func (c *Conn) GetUserById(id string) (user dao.User, err error) {
	err = c.db.QueryRow(`
		SELECT 
			id,
			email,
			name,
			password,
			created_at
		FROM 
			users
		WHERE 
			id = $1`,
		id).Scan(&user.ID, &user.Email, &user.Name, &user.HashedPassword, &user.CreatedAt)
	if err != sql.ErrNoRows {
		err = errors.New("user not found")
	}
	return
}

func (c *Conn) CreateUser(user *dao.User) error {
	_, err := c.db.Exec(`INSERT INTO users (id, name, password, email) VALUES ($1, $2, $3, $4)`,
		user.ID, user.Name, user.HashedPassword, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) GetUserByEmail(email string) (user dao.User, err error) {
	err = c.db.QueryRow(`
		SELECT 
			id,
			email,
			name,
			password,
			created_at
		FROM 
			users
		WHERE 
			email = $1`,
		email).Scan(&user.ID, &user.Email, &user.Name, &user.HashedPassword, &user.CreatedAt)
	if err == sql.ErrNoRows {
		err = errors.New("user not found")
	}
	return
}

func (c *Conn) GetAllUsers() ([]dao.User, error) {
	rows, err := c.db.Query(`
		SELECT
			id,
			email,
			name,
			password,
			created_at
		FROM
			users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dao.User
	for rows.Next() {
		var user dao.User
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.HashedPassword, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}
