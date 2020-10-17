package socket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// conn provides access to the web socket.
var conn *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Write writes the provided string to the web socket.
func Write(msg string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		// TODO handle errors or return
		panic(err)
	}
}

// Read will block execution until a message (from the web socket) is delivered by channel.
func Read() string {
	msg := <-out
	cmsg := msg.(chanMessage)
	return cmsg.Message
}

// The in channel is continuously populated by the web socket connection ReadMessage() call.
var in chan<- interface{}

// The out channel will output messages taken from the in channel above. It can be accessed in a blocking
// manner through the Read() method.
var out <-chan interface{}

// chanMessage is used internally to satisfy the interface{} method required for a buffered channel.
type chanMessage struct {
	Message string
}

// Start performs the required work to intitialise the buffered channel used for reading messages from the
// web socket from the game engine, in a blocking manner. It takes bool channel it uses to indicate when
// it is ready for the game engine to start (once the channels are set up)
func Start(done chan bool, w http.ResponseWriter, r *http.Request) {
	initWebSocket(w, r)
	initChans()
	done <- true
	readForever()
}

func initWebSocket(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(fmt.Sprintf("failed call to web socket Upgrade: %v", err))
	}
}

// initChans sets up the buffered channel used to recieve messages from the web socket and block on reads.
func initChans() {
	in, out = setUpBufferedChannel()
}

func setUpBufferedChannel() (chan<- interface{}, <-chan interface{}) {
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

func readForever() {
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("failed to Read Message: %v", err)
			return
		}
		cmsg := chanMessage{
			Message: string(msg),
		}
		in <- cmsg
	}
}
