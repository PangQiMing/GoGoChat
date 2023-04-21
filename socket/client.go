package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//var (
//	newline = []byte{'\n'}
//	space   = []byte{' '}
//)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	goGoID string
	hub    *Hub
	conn   *websocket.Conn
	send   chan Message
}

func (c *Client) readPump() {
	for {
		var message Message
		err := c.conn.ReadJSON(&message)
		//_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			log.Println("11" + err.Error())
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, _ := <-c.send:
			err := c.conn.WriteJSON(&message)
			//err := c.conn.WriteMessage(1, message)
			if err != nil {
				log.Println("22" + err.Error())
				return
			}

			_, err = c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("error 11")
				return
			}
		}
	}
}
