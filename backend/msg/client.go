package msg

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type ClientMsg struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *ClientMsg) Read() {
	defer func() {
		c.Pool.Unregister <- c
		_ = c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{
			Type: messageType,
			Body: string(p),
		}
		c.Pool.BroadCast <- message
		fmt.Printf("Message Receiver: %+v\n", message)
	}
}
