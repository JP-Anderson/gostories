package main

import (
	"net/http"

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

func initServer() {
	webSocketReadyChan := make(chan bool, 1)
	go func() {
		http.HandleFunc("/sock", func(w http.ResponseWriter, r *http.Request) {
			server.Start(webSocketReadyChan, w, r)
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "websockets.html")
		})
		http.ListenAndServe(":8080", nil)
	}()
	<-webSocketReadyChan
}
