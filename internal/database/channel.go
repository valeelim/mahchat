package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/valeelim/mahchat/pkg/dao"
)

func (c *Conn) CreateChannel(channel *dao.Channel) error {
	_, err := c.db.Exec("INSERT INTO channels (id, name) VALUES ($1, $2)",
		channel.ID, channel.Name)
	if err != nil {
		log.Println("create channel error")
		return err
	}
	return nil
}

func (c *Conn) GetChannelByID(id int64) (channel dao.Channel, err error) {
	err = c.db.QueryRow(`
			SELECT 
				id,
				name,
				created_at
			FROM
				channels
			WHERE
				id = $1
		`, id).Scan(&channel.ID, &channel.Name, &channel.CreatedAt)
	if err == sql.ErrNoRows {
		err = errors.New("channel not found")
	}
	return
}
