package socketio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gostories/engine/parser"
	"gostories/server"
)

// SocketInputOutputHandler manages the games user Input and Output through the gorilla
// websocket interface.
type SocketInputOutputHandler struct {
	reader *bufio.Reader
}

// NewSocketInputOutputHandler creates a simple Input Output Handler for playing the game via console.
func NewSocketInputOutputHandler() *SocketInputOutputHandler {
	return &SocketInputOutputHandler{
		reader: bufio.NewReader(os.Stdin),
	}
}

// NewLine takes a string and prints it to the console.
func (c *SocketInputOutputHandler) NewLine(output string) error {
	fmt.Println(output)
	server.Write(output)
	return nil
}

// NewLinef takes a format string and a series of values to interpolate in the format string.
func (c *SocketInputOutputHandler) NewLinef(output string, args ...interface{}) error {
	c.NewLine(fmt.Sprintf(output, args...))
	return nil
}

// ReadInt tries to parse console input as an int. It returns the int or errors.
func (c *SocketInputOutputHandler) ReadInt() (i int, e error) {
	input := Trim(c.readString())
	return strconv.Atoi(input)
}

// ReadIntInRange tries to parse console input as an int in an inclusive range. It will continuously prompt the
// user until a valid integer within the desired range is provided.
func (c *SocketInputOutputHandler) ReadIntInRange(lowest, highest int) (i int) {
	valid := false
	for !valid {
		input, err := c.ReadInt()
		if err != nil {
			c.NewLine("Please enter an int")
			continue
		}
		if input < lowest || input > highest {
			c.NewLinef("Please enter an int in range %v <= x <= %v", lowest, highest)
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
func (c *SocketInputOutputHandler) SimpleParse() (parser.Action, []string) {
	input := c.readString()
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

func (c *SocketInputOutputHandler) readString() string {
	print(">> ")
	s := <-server.Out
	cmsg := s.(server.ChanMessage)
	return cmsg.Message
}

const linuxCutset = "\n"
const windowsCutset = "\r" + linuxCutset

// Trim returns a string with spaces to the right trimmed, and a line ending.
func Trim(input string) string {
	return strings.TrimRight(input, linuxCutset)
}