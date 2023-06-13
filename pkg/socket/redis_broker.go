package socket

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisBroker[T any] struct {
	Client  *redis.Client
	Channel string
	done    chan struct{}
}

func NewRedisBroker[T any](channel string, client *redis.Client) (*RedisBroker[T], func()) {	
	rdb := &RedisBroker[T]{
		Client:  client,
		Channel: channel,
		done: make(chan struct{}),
	}
	close := func() {
		log.Println("broker  closed")
		close(rdb.done)
	}
	return rdb, close
}

func (rdb *RedisBroker[T]) Publish(ctx context.Context, msg T) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := rdb.Client.Publish(ctx, rdb.Channel, string(json)).Err(); err != nil {
		return err
	}
	return nil
}

func (rdb *RedisBroker[T]) Subscribe() (<-chan T, func()) {
	ch := make(chan T)
	go func() {
		rdb.subscribe(ch)
	}()
	close := func() { 
		close(ch)
	}
	return ch, close
}

func (rdb *RedisBroker[T]) subscribe(ch chan<- T) {
	ctx := context.Background()

	pubsub := rdb.Client.Subscribe(ctx, rdb.Channel)
	defer pubsub.Close()

	if _, err := pubsub.Receive(ctx); err != nil {
		log.Println("Subscribe fails")
	}
	for {
		select {
		case <-rdb.done:
			return
		case data := <-pubsub.Channel():
			var msg T

			if err := json.Unmarshal([]byte(data.Payload), &msg); err != nil {
				log.Println("Error unmarshal json")
				continue
			}
			select {
			case <-rdb.done:
				return
			case ch <- msg:
			}
		}
	}
}
