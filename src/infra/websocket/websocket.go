package ws

import (
	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

const (
	TextMessageType = iota
	BinaryMessageType
	CloseMessageType
)

// Upgrader Upgrade http into websocket.
//
// ```
// NewWebsocketUpgrader()
// ```
type Upgrader struct {
	Ref websocket.Upgrader
}

type WebsocketUpgraderOptions = func(*Upgrader)

func NewWebsocketUpgrader(options ...WebsocketUpgraderOptions) *Upgrader {
	var upgrader = &Upgrader{
		Ref: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	for _, opt := range options {
		if opt != nil {
			opt(upgrader)
		}
	}
	return upgrader
}

func (u Upgrader) Upgrade(writer http.ResponseWriter, request *http.Request, header http.Header) (*Connection, error) {
	conn, err := u.Ref.Upgrade(writer, request, header)
	if err != nil {
		return nil, err
	}
	wsConn := &Connection{
		mutex:     sync.Mutex{},
		Session:   uuid.NewString(),
		WebSocket: conn,
		InChan:    make(chan *Message, 256),
		OutChan:   make(chan *Message, 256),
		CloseChan: make(chan byte),
		Closed:    false,
	}

	go wsConn.startReadLoop()
	go wsConn.startWriteLoop()

	return wsConn, nil
}

type Message struct {
	Type int
	Data []byte
}

type Connection struct {
	mutex     sync.Mutex
	Session   Session
	WebSocket *websocket.Conn
	InChan    chan *Message
	OutChan   chan *Message
	CloseChan chan byte
	Closed    bool
}

func (c *Connection) WriteMessage(message Message) error {
	select {
	case c.OutChan <- &message:
	case <-c.CloseChan:
	}
	return nil
}

func (c *Connection) ReadMessage() (*Message, error) {
	select {
	case message := <-c.InChan:
		return message, nil
	case <-c.CloseChan:
		return nil, errors.Errorf("websocket closed with session %s", c.Session)
	}
}

func (c *Connection) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.Closed {
		c.Closed = true
		c.CloseChan <- 0
		c.WebSocket.Close()
	}
}

func (c *Connection) startReadLoop() {
	var (
		message     *Message
		messageType int
		data        []byte
		err         error
	)

	for {
		if messageType, data, err = c.WebSocket.ReadMessage(); err != nil {
			c.Close()
			break
		}
		message = &Message{
			Type: messageType,
			Data: data,
		}

		select {
		case c.InChan <- message:
		case <-c.CloseChan:
			break
		}
	}
}

func (c *Connection) startWriteLoop() {
	var (
		message *Message
		err     error
	)

	for {
		select {
		case message = <-c.OutChan:
			if err = c.WebSocket.WriteMessage(message.Type, message.Data); err != nil {
				break
			}
		case <-c.CloseChan:
			break
		}
	}
}
