package socket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/valeelim/mahchat/pkg/utils"
)

const (
	writeTimeout = 30 * time.Second

	pongTimeout = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongTimeout * 9) / 10

	maxMessageSize = 512
)

type Socket[T any] struct {
	ID   int64
	conn *websocket.Conn

	quit    sync.Once
	done    chan struct{}
	readCh  chan T 
	writeCh chan any 
}

func NewSocket[T any](conn *websocket.Conn) *Socket[T] {
	id, _ := utils.GenerateSnowflake()
	socket := &Socket[T]{
		ID:   id,
		conn: conn,

		done:    make(chan struct{}),
		readCh:  make(chan T),
		writeCh: make(chan any),
	}

	go socket.writer()
	go socket.reader()
	return socket
}

func (s *Socket[T]) Close() {
	s.quit.Do(func() {
		log.Println("Socket connection closed")
		close(s.done)
		if err := s.conn.Close(); err != nil {
			log.Println("closing error", err)
		}
	})
}

// writes to websocket
func (s *Socket[T]) writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Println("writer stopped")
		ticker.Stop()
		close(s.writeCh)
	}()
	for {
		select {
		case <-s.done:
			s.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		case msg := <-s.writeCh:
			s.conn.SetWriteDeadline(time.Now().Add(writeTimeout))

			if err := s.conn.WriteJSON(msg); err != nil {
				log.Println("write json error", err)
				return
			}
		case <-ticker.C:
			s.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			// send heartbeat to websocket
			if err := s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// reads from websocket to read channel
func (s *Socket[T]) reader() {
	defer func() {
		log.Println("reader closed")
		close(s.readCh)
	}()
	for {
		s.conn.SetReadLimit(maxMessageSize)
		s.conn.SetReadDeadline(time.Now().Add(pongTimeout))
		s.conn.SetPongHandler(func(string) error {
			return s.conn.SetReadDeadline(time.Now().Add(pongTimeout))
		})

		var msg T 
		if err := s.conn.ReadJSON(&msg); err != nil {
			log.Println("Read json error", err, msg)
			return
		}

		select {
		case <-s.done:
			log.Println("reader done called")
			s.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		case s.readCh <- msg:
		}
	}
}

func (s *Socket[T]) Listen() <-chan T {
	return s.readCh
}

func (s *Socket[T]) Emit(msg any) {
	select {
	case <-s.done:
		return
	case s.writeCh <- msg:
	}	
}

