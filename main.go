package main

import (
	"log"
	"net/http"
	"os"

	"gostories/engine"
	"gostories/gen/areas"
	"gostories/socket"
)

func main() {
	initServer()
	startRoom := areas.Get("cat_room")
	stage := engine.Stage{}
	stage.Start(*startRoom)
}

func initServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Deploying to %s", port)
	}
	webSocketReadyChan := make(chan bool, 1)
	go func() {
		http.HandleFunc("/sock", func(w http.ResponseWriter, r *http.Request) {
			socket.Start(webSocketReadyChan, w, r)
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "websockets.html")
		})
		http.ListenAndServe(":"+port, nil)
	}()
	<-webSocketReadyChan
}
