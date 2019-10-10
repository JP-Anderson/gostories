package consoleio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gostories/parser"
)

// ConsoleInputOutputHandler manages the games user Input and Output through some "standard" Go
// packages such as "os" and "bufio".
type ConsoleInputOutputHandler struct {
	reader *bufio.Reader
}

// NewConsoleInputOutputHandler creates a simple Input Output Handler for playing the game via console.
func NewConsoleInputOutputHandler() *ConsoleInputOutputHandler {
	return &ConsoleInputOutputHandler{
		reader: bufio.NewReader(os.Stdin),
	}
}

// NewLine takes a string and prints it to the console.
func (c *ConsoleInputOutputHandler) NewLine(output string) error {
	fmt.Println(output)
	return nil
}

// NewLinef takes a format string and a series of values to interpolate in the format string.
func (c *ConsoleInputOutputHandler) NewLinef(output string, args ...interface{}) error {
	c.NewLine(fmt.Sprintf(output, args...))
	return nil
}

// ReadInt tries to parse console input as an int. It returns the int or errors.
func (c *ConsoleInputOutputHandler) ReadInt() (i int, e error) {
	input := Trim(c.readString())
	return strconv.Atoi(input)
}

// SimpleParse parses input from the user. Currently only one or two (space-separated) strings can
// be parsed. SimpleParse returns the first string as an action (if recognised), and the second
// string (the target verb) as is.
func (c *ConsoleInputOutputHandler) SimpleParse() (parser.Action, string) {
	input := c.readString()
	split := strings.Split(input, " ")
	if len(split) >= 2 {
		return parser.ParseInput(split[0], Trim(split[1]))
	} else if len(split) == 1 {
		return parser.ParseInput(Trim(split[0]), "")
	}
	return parser.Unknown(), ""
}

func (c *ConsoleInputOutputHandler) readString() string {
	input, err := c.reader.ReadString('\n')
	if err != nil {
		c.NewLinef("ReadString error: %v", err)
	}
	return input
}

const linuxCutset = "\n"
const windowsCutset = "\r" + linuxCutset

// Trim returns a string with spaces to the right trimmed, and a line ending.
func Trim(input string) string {
	return strings.TrimRight(input, windowsCutset)
}
