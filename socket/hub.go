package socket

import (
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/PangQiMing/GoGoChat/service"
	"log"
	"strconv"
)

type Hub struct {
	// Registered onlineClients.
	onlineClients map[string]*Client

	// Inbound messages from the onlineClients.
	broadcast chan entity.Message

	// Register requests from the onlineClients.
	register chan *Client

	// Unregister requests from onlineClients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:     make(chan entity.Message),
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
			if message.MessageType == "private" {
				err := service.SaveMessage(message)
				if err != nil {
					continue
				}
				if localClient, ok := h.onlineClients[message.From]; ok {
					localClient.send <- message
					if client, ok := h.onlineClients[message.To]; ok {
						client.send <- message
					}
				} else {
					log.Println("该用户不在线")
				}
			} else if message.MessageType == "public" {
				groups := service.GetGroupMemberByToId(message.To)
				err := service.SaveMessage(message)
				if err != nil {
					continue
				}
				for _, member := range groups {
					if memberClient, ok := h.onlineClients[strconv.FormatUint(member.MemberID, 10)]; ok {
						memberClient.send <- message
					}
				}
				//log.Println(groups)
			}
		}
	}
}
