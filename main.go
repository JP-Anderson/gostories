package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"gostories/engine"
	"gostories/gen/areas"
	"gostories/server"
)

func main() {
	initServer()
	startRoom := areas.Get("cat_room")
	stage := engine.Stage{}
	stage.Start(*startRoom)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func initServer() {
	webSocketReadyChan := make(chan bool, 1)
	go func() {
		http.HandleFunc("/sock", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				panic(fmt.Sprintf("failed call to web socket Upgrade: %v", err))
			}
			server.Conn = conn
			server.Start(webSocketReadyChan)
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "websockets.html")
		})
		http.ListenAndServe(":8080", nil)
	}()
	<-webSocketReadyChan
}
