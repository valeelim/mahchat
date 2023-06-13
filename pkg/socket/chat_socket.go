package socket

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/valeelim/mahchat/pkg/dto"
)

// represents a group websocket connections

const (
	readBufferSize  = 1024 * 1024 * 1024
	writeBufferSize = 1024 * 1024 * 1024
)

type IO[T any] struct {
	websocket.Upgrader
	mu      sync.RWMutex
	sockets map[int64]*Socket[T]
}

func NewIO[T any]() *IO[T] {
	return &IO[T]{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sockets: make(map[int64]*Socket[T]),
	}
}

// initiates websocket connection
func (io *IO[T]) ServeWS(c *gin.Context) (*Socket[T], func(), error) {
	log.Println("sus")
	conn, err := io.Upgrade(c.Writer, c.Request, nil)
	log.Println("very sus")
	if err != nil {
		log.Fatal("serve ws fail", err)
		return nil, nil, errors.New("Failed to upgrade websocket")
	}
	socket := NewSocket[T](conn)

	io.registerSocket(socket)
	return socket, func() {
		io.deregisterSocket(socket)
	}, nil
}

func (io *IO[T]) EmitMsg(socketID int64, msg dto.Message) {
	io.mu.Lock()
	socket, ok := io.sockets[socketID]
	io.mu.Unlock()
	if !ok {
		log.Println("socket id does not exist")
		return
	}
	socket.Emit(msg)
}

func (io *IO[T]) registerSocket(socket *Socket[T]) {
	io.mu.Lock()
	io.sockets[socket.ID] = socket
	io.mu.Unlock()
}

func (io *IO[T]) deregisterSocket(socket *Socket[T]) {
	io.mu.Lock()
	if socket, ok := io.sockets[socket.ID]; ok {
		log.Println("A client disconnects")
		socket.Close()
		delete(io.sockets, socket.ID)
	}
	io.mu.Unlock()
}
