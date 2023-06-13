package repository

import (
	"context"
)

type Cache interface {
	GetAccessToken(ctx context.Context, key string) (map[string]string, error)
	SetAccessToken(ctx context.Context, key string, values ...interface{}) error
}
