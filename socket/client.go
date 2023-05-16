package socket

import (
	"fmt"
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

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
	send   chan entity.Message
}

func (c *Client) readPump() {
	for {
		var message entity.Message
		err := c.conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			log.Println("11" + err.Error())
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, _ := <-c.send:
			err := c.conn.WriteJSON(&message)
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
