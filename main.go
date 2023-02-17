package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gostories/engine"
	"gostories/engine/io"
	"gostories/gen/areas"
	"gostories/socket"
)

func main() {
	mode := getRunMode()
	io.SetIOMode(mode)
	if mode == io.SocketIOMode {
		initServer()
	}
	startRoom := areas.Get("cat_room")
	stage := engine.Stage{}
	stage.Start(*startRoom)
}

func getRunMode() int {
	if len(os.Args) > 1 {
		if os.Args[1] == "cmd" {
			fmt.Println("Running CONSOLE IO Mode")
			return io.ConsoleIOMode
		}
	}
	fmt.Println("Running SOCKET IO Mode")
	return io.SocketIOMode
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
