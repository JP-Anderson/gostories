package socketio

import (
	"fmt"
	"strconv"

	"gostories/engine/parser"
	"gostories/socket"
	"gostories/utils/strings"
)

// SocketInputOutputHandler manages the games user Input and Output through the gorilla
// websocket interface.
type SocketInputOutputHandler struct{}

// NewSocketInputOutputHandler creates a simple Input Output Handler for playing the game via console.
func NewSocketInputOutputHandler() *SocketInputOutputHandler {
	return &SocketInputOutputHandler{}
}

// NewLine takes a string and prints it to the console.
func (s *SocketInputOutputHandler) NewLine(output string) error {
	socket.Write(output)
	return nil
}

// NewLinef takes a format string and a series of values to interpolate in the format string.
func (s *SocketInputOutputHandler) NewLinef(output string, args ...interface{}) error {
	return s.NewLine(fmt.Sprintf(output, args...))
}

// ReadInt tries to parse console input as an int. It returns the int or errors.
func (s *SocketInputOutputHandler) ReadInt() (i int, e error) {
	input := strings.Trim(s.readString())
	return strconv.Atoi(input)
}

// ReadIntInRange tries to parse console input as an int in an inclusive range. It will continuously prompt the
// user until a valid integer within the desired range is provided.
func (s *SocketInputOutputHandler) ReadIntInRange(lowest, highest int) (i int) {
	valid := false
	for !valid {
		input, err := s.ReadInt()
		if err != nil {
			s.NewLine("Please enter an int")
			continue
		}
		if input < lowest || input > highest {
			s.NewLinef("Please enter an int in range %v <= x <= %v", lowest, highest)
			continue
		}
		valid = true
		i = input
	}
	return i
}

// SimpleParse passes input from a web socket to the command parser.
func (s *SocketInputOutputHandler) SimpleParse() (parser.Action, []string) {
	return parser.SimpleParse(s.readString)
}

func (s *SocketInputOutputHandler) readString() string {
	return socket.Read()
}
