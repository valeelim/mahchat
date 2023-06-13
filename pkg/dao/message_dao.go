package dao

import (
	"time"

	"github.com/valeelim/mahchat/pkg/utils"
)

type Message struct {
	ID        int64     `json:"id,omitempty"`
	Sender    int64     `json:"sender" binding:"required"`
	Recipient int64     `json:"recipient" binding:"required"`
	Content   string    `json:"content,omitempty" binding:"required,min=1"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type GroupMessage struct {
	ChannelID int64     `json:"channel_id,omitempty"`
	MessageID int64     `json:"message_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewGroupMessage(channelID, userID int64, content string) (*GroupMessage, error) {
	id, err := utils.GenerateSnowflake()
	if err != nil {
		return nil, err
	}
	return &GroupMessage{
		ChannelID: channelID,
		MessageID: id,
		UserID:    userID,
		Content:   content,
	}, nil
}
