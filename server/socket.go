package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Conn provides access to the web socket.
var Conn *websocket.Conn

// Write writes the provided string to the web socket.
func Write(msg string) {
	err := Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		// TODO handle errors or return
		panic(err)
	}
}

// Channels used to block the game engine to wait for user input from a web socket.
var In chan<- interface{}
var Out <-chan interface{}

type ChanMessage struct {
	Message string
}

func InitChans() {
	In, Out = initChans()
}

func initChans() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})
	go func() {
		var q []interface{}
		outChan := func() chan interface{} {
			if len(q) == 0 {
				return nil
			}
			return out
		}
		curVal := func() interface{} {
			if len(q) == 0 {
				return nil
			}
			return q[0]
		}
	loop:
		for {
			select {
			case v, ok := <-in:
				if !ok {
					break loop
				} else {
					q = append(q, v)
				}
			case outChan() <- curVal():
				q = q[1:]
			}
		}
		close(out)
	}()
	return in, out
}

func ReadForever() {
	InitChans()
	for {
		// Read message from browser
		_, msg, err := Conn.ReadMessage()
		if err != nil {
			fmt.Printf("failed to Read Message: %v", err)
			return
		}
		cmsg := ChanMessage{
			Message: string(msg),
		}
		In <- cmsg
	}
}
