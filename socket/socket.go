package socket

import (
	"github.com/PangQiMing/GoGoChat/entity"
	"log"
	"net/http"
)

func ServeWs1(hub *Hub, goGoID string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan entity.Message), goGoID: goGoID}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
