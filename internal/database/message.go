package database

import (
	"log"

	"github.com/valeelim/mahchat/pkg/dao"
)

func (c *Conn) CreateGroupMessage(msg *dao.GroupMessage) error {
	_, err := c.db.Exec(`
			INSERT INTO group_messages
				(channel_id, message_id, user_id, content)
			VALUES
				($1, $2, $3, $4)
		`, msg.ChannelID, msg.MessageID, msg.UserID, msg.Content)
	if err != nil {
		log.Println("create group message error", err)
		return err
	}
	return nil
}

func (c *Conn) GetGroupMessageByChannelID(channelID string) ([]dao.GroupMessage, error) {
	rows, err := c.db.Query(`
			SELECT
				(channel_id, message_id, user_id, content, created_at)
			FROM
				group_message
			WHERE
				channel_id = $1
		`, channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []dao.GroupMessage
	for rows.Next() {
		var msg dao.GroupMessage
		err := rows.Scan(&msg.ChannelID, &msg.MessageID, &msg.UserID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, msg)
	}
	return result, nil
}
