package entity

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	GoGoID   uint64          //客户端ID
	Conn     *websocket.Conn //webSocket连接接口
	SendMsg  chan *Message   //需要发送的消息
	Hun      *Hub            //连接hub的接口，使用hub可以注册，注销客户端，广播消息
	Avatar   string          //客户端头像
	Nickname string          //客户端呢称
}

// 客户端从WebSocket连接读取数据
func (c *Client) Read() {
	//向Hub发送注销当前客户端操作，并关闭WebSocket连接
	defer func() {
		c.Hun.UnRegisterClient <- c
		c.Conn.Close()
	}()

	for {
		message := &Message{}
		//从WebSocket里读取Message数据
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Println("error:", err)
			break
		}
		//设置客户端的头像、呢称
		c.Avatar = message.Avatar
		c.Nickname = message.Nickname
		//把Message发送给Hub.Broadcast channel
		c.Hun.Broadcast <- message
		fmt.Printf("read: %v", message)
	}
}

// 客户端从WebSocket连接写入数据
func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.SendMsg:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			fmt.Println("write: ", message)
			err := c.Conn.WriteJSON(message)
			if err != nil {
				log.Println("error", err)
				break
			}
		}
	}
}
