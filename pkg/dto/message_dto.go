package dto

import "github.com/valeelim/mahchat/pkg/dao"

type Message struct {
	Sender    int64  `json:"sender,omitempty"`
	Recipient int64  `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type CreateGroupMessageRequest struct {
	ChannelID int64  `json:"channel_id,omitempty"`
	UserID    int64  `json:"user_id,omitempty"`
	Content   string `json:"content,omitempty"`
}

type GetGroupMessages struct {
	Data []dao.GroupMessage `json:"data,omitempty"`
}
