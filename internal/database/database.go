package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

type Conn struct {
	db *sql.DB
}

func New(user, port, password, dbname, host string) (*Conn, error) {
	connStr := fmt.Sprintf("user=%s port=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("DBUSER"), os.Getenv("DBPORT"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), os.Getenv("DBHOST"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}
	fmt.Println("Connected!")
	return &Conn{db: db}, nil
}

func (c *Conn) Close() error {
	return c.db.Close()	
}


