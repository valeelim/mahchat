package repository

import "github.com/valeelim/mahchat/pkg/dao"

type Channel interface {
	CreateChannel(channel *dao.Channel) error
	GetChannelByID(id int64) (dao.Channel, error)
}
