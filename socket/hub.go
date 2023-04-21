package socket

import (
	"log"
)

type Message struct {
	Avatar       string `json:"avatar"`
	FromUsername string `json:"from_username"`
	From         string `json:"from"`
	To           string `json:"to"`
	Content      string `json:"content"`
	MessageType  string `json:"message_type"`
}

type Hub struct {
	// Registered onlineClients.
	onlineClients map[string]*Client

	// Inbound messages from the onlineClients.
	broadcast chan Message

	// Register requests from the onlineClients.
	register chan *Client

	// Unregister requests from onlineClients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:     make(chan Message),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		onlineClients: make(map[string]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.onlineClients[client.goGoID] = client
		case client := <-h.unregister:
			if _, ok := h.onlineClients[client.goGoID]; ok {
				delete(h.onlineClients, client.goGoID)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Println(message)
			//var msg Message
			//err := json.Unmarshal(message, &msg)
			//if err != nil {
			//	log.Println(err.Error())
			//	return
			//}

			if message.MessageType == "private" {
				if localClient, ok := h.onlineClients[message.From]; ok {
					localClient.send <- message
					if client, ok := h.onlineClients[message.To]; ok {
						client.send <- message
					}
				} else {
					log.Println("该用户不在线")
				}
			}
		}
	}
}
