package repository

import "github.com/valeelim/mahchat/pkg/dao"

type GroupMessage interface {
	CreateGroupMessage(msg *dao.GroupMessage) error
	GetGroupMessageByChannelID(channelID string) ([]dao.GroupMessage, error)
}