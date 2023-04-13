package service

import (
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWS WS服务
func ServeWS(goGoID uint64, hub *entity.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	client := &entity.Client{
		GoGoID:  goGoID,
		Conn:    conn,
		SendMsg: make(chan *entity.Message),
		Hun:     hub,
	}

	msg := &entity.Message{
		Avatar:   "00",
		Nickname: "jack",
		Content:  "hello",
	}
	client.SendMsg <- msg

	hub.RegisterClient <- client
	go client.Read()
	go client.Write()
}
