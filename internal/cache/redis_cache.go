package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func New(addr, password string, db int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatal("redis not connected", err)
		return nil, err
	}
	return &RedisCache{Client: client}, nil
}

func (r *RedisCache) SetAccessToken(ctx context.Context, key string, values ...interface{}) error {
	pipe := r.Client.TxPipeline()

	pipe.HSet(ctx, key, values...)
	pipe.Expire(ctx, key, 1*time.Hour)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisCache) GetAccessToken(ctx context.Context, key string) (map[string]string, error) {
	result, err := r.Client.HGetAll(ctx, key).Result()
	if len(result) == 0 {
		return nil, redis.Nil
	}
	return result, err
}

func (r *RedisCache) GetClient() (*redis.Client, error) {
	return r.Client, nil
}
