package chat

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/valeelim/mahchat/pkg/dto"
	"github.com/valeelim/mahchat/pkg/socket"
)

type Chat struct {
	remote *socket.RedisBroker[dto.Message]
	local  *socket.IO[dto.Message]
	done   chan struct{}
}

// initializes a broker for a channel
func New(channel string, client *redis.Client) (*Chat, func()) {
	io := socket.NewIO[dto.Message]()
	broker, closeBroker := socket.NewRedisBroker[dto.Message](channel, client)

	log.Println("new broker for a channel")

	chat := &Chat{
		local:  io,
		remote: broker,
		done:   make(chan struct{}),
	}

	go chat.startChannel()

	return chat, closeBroker
}

// corresponds to a user connected to the remote broker
func (c *Chat) ServeWS(ctx *gin.Context) {
	// returns a new websocket credentials for that user
	log.Println("hahaha")
	socket, flush, err := c.local.ServeWS(ctx)
	if err != nil {
		log.Println("serve local error", err)
		return
	}
	defer flush()
	defer close(c.done)
	defer c.removeSession(c.remote.Channel, socket.ID)

	c.addSession(c.remote.Channel, socket.ID)


	go func() {
		ch, close := c.remote.Subscribe()
		defer close()

		for {
			select {
			case <-c.done:
				return
			case msg := <-ch:
				socket.Emit(msg)
			}
		}
	}()

	for msg := range socket.Listen() {
		log.Println("publishing message", msg)
		c.remote.Publish(context.Background(), msg)
	}

}

func (c *Chat) addSession(channel string, socketID int64) {
	c.remote.Client.SAdd(context.Background(), channel, []int64{socketID})
}

func (c *Chat) getSessions() ([]string, error) {
	return c.remote.Client.SMembers(context.Background(), c.remote.Channel).Result()
}

func (c *Chat) removeSession(channel string, socketID int64) {
	c.remote.Client.SRem(context.Background(), channel, []int64{socketID})
}

func (c *Chat) startChannel() {
	ch, close := c.remote.Subscribe()
	defer close()

	for {
		select {
		case <-c.done:
			return
		case msg := <-ch: // if there's a published message, push to local socket
			c.emitLocal(msg)
		}
	}
}

// push messages to everyone
func (c *Chat) emitLocal(msg dto.Message) {
	subs, err := c.getSessions()
	if err != nil {
		log.Println("Emit local error", err)
		return
	}
	for _, sub := range subs {
		socketID, _ := strconv.ParseInt(sub, 10, 64)	
		c.local.EmitMsg(socketID, msg)
	}
}
