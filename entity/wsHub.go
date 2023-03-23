package entity

import (
	"fmt"
	"sync"
)

type Hub struct {
	Clients          map[*Client]bool //注册客户端，用户在线设置为True
	Broadcast        chan *Message    //广播消息，把消息发送给另一个客户端或者多个客户端
	RegisterClient   chan *Client     //注册客户端
	UnRegisterClient chan *Client     //注销客户端
	Lock             sync.RWMutex     //使用读写互斥锁，防止并发时出现数据不同步
}

func NewHub() *Hub {
	return &Hub{
		Clients:          make(map[*Client]bool),
		Broadcast:        make(chan *Message),
		RegisterClient:   make(chan *Client),
		UnRegisterClient: make(chan *Client),
	}
}

// Run Hub的事件处理,注册客户端，注销客户端，广播消息
func (h *Hub) Run() {
	for {
		select {
		//注册客户端
		case client := <-h.RegisterClient:
			h.Lock.Lock()
			h.Clients[client] = true
			h.Lock.Unlock()
		//注销客户端
		case client := <-h.UnRegisterClient:
			if _, ok := h.Clients[client]; ok {
				h.Lock.Lock()
				delete(h.Clients, client)
				h.Lock.Unlock()
				close(client.SendMsg)
			}
		//广播消息
		case message := <-h.Broadcast:
			h.Lock.Lock()
			fmt.Println("Broadcast:", message)
			//遍历Clients注册表里的用户
			for client := range h.Clients {
				select {
				//client里所有用户接收来自Hub的广播消息
				case client.SendMsg <- message:
				default:
					close(client.SendMsg)
					delete(h.Clients, client)
				}
			}
			h.Lock.Unlock()
		}
	}
}
