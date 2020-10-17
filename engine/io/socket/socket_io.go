package socketio

import (
	"fmt"
	"strconv"
	"strings"

	"gostories/engine/parser"
	"gostories/server"
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
	server.Write(output)
	return nil
}

// NewLinef takes a format string and a series of values to interpolate in the format string.
func (s *SocketInputOutputHandler) NewLinef(output string, args ...interface{}) error {
	return s.NewLine(fmt.Sprintf(output, args...))
}

// ReadInt tries to parse console input as an int. It returns the int or errors.
func (s *SocketInputOutputHandler) ReadInt() (i int, e error) {
	input := Trim(s.readString())
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

// SimpleParse parses input from the user. Currently only one or two (space-separated) strings can
// be parsed. SimpleParse returns the first string as an action (if recognised), and the second
// string (the target verb) as is.
func (s *SocketInputOutputHandler) SimpleParse() (parser.Action, []string) {
	input := s.readString()
	split := strings.Split(input, " ")
	len := len(split)
	if len > 2 {
		return parser.ParseInput(split...)
	} else if len == 2 {
		return parser.ParseInput(Trim(split[0]), Trim(split[1]))
	} else if len == 1 {
		return parser.ParseInput(Trim(split[0]), "")
	}
	return parser.Unknown(), []string{""}
}

func (s *SocketInputOutputHandler) readString() string {
	return server.Read()
}

const linuxCutset = "\n"
const windowsCutset = "\r" + linuxCutset

// Trim returns a string with spaces to the right trimmed, and a line ending.
func Trim(input string) string {
	return strings.TrimRight(input, linuxCutset)
}
