package module_socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

//
//func WebSocketStart() {
//	hub := newHub()
//	go hub.run()
//	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
//		serveWs(hub, w, r)
//	})
//}

var hub = newHub()

func RegisterWebSocket(c *gin.Context) {
	go hub.run()
	//go broadcast()
	var w = c.Writer
	var r = c.Request
	serveWs(hub, w, r)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//
	//})
}

func Broadcast(message string) {
	m := []byte(message)
	hub.broadcast <- m
}
