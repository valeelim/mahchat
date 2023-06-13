package dto

import "time"

type CreateChannelRequest struct {
	Name string `json:"name,omitempty"`
}

type CreateChannelResponse struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetChannelResponse struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
