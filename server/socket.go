package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

var MsgType int

func Write(msg string) {
	fmt.Printf("message type is %d\n", MsgType)
	err := Conn.WriteMessage(MsgType, []byte(msg))
	if err != nil {
		// TODO handle errors or return
		panic(err)
	}
}
