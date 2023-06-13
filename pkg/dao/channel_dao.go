package dao

import (
	"time"

	"github.com/valeelim/mahchat/pkg/utils"
)

type Channel struct {
	ID			int64		`json:"id,omitempty"`
	Name		string		`json:"name,omitempty"`
	CreatedAt 	time.Time	`json:"created_at,omitempty"`
}

func NewChannel(name string) (*Channel, error) {
	id, err := utils.GenerateSnowflake()
	if err != nil {
		return nil, err
	}
	return &Channel{
		ID: id,
		Name: name,
	}, nil
}
