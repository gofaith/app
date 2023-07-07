package app

import (
	"log"
	"net/http"

	"github.com/StevenZack/tools/strToolkit"
	"github.com/gorilla/websocket"
)

type bridgeServer struct {
	port     string
	upgrader websocket.Upgrader
}

func newBridgeServer() *bridgeServer {
	v := &bridgeServer{
		port: strToolkit.RandomPort(),
	}
	http.HandleFunc("/ws", v.ws)
	return v
}

func (b *bridgeServer) ws(w http.ResponseWriter, r *http.Request) {
	conn, e := b.upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Println(e)
		return
	}
	defer conn.Close()

	for {
		_, m, e := conn.ReadMessage()
		if e != nil {
			log.Println(e)
			return
		}
		
	}
}

func (b *bridgeServer) Run() error {
	e := http.ListenAndServe(":"+b.port, nil)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}
