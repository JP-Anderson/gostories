package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gostories/parser"
)

// NewLine takes a string and prints it to the console.
func NewLine(output string) {
	fmt.Println(output)
}

// NewLinef takes a format string and a series of values to interpolate in the format string.
func NewLinef(output string, args ...interface{}) {
	NewLine(fmt.Sprintf(output, args...))
}

var reader = bufio.NewReader(os.Stdin)

// ReadInt tries to parse console input as an int. It returns the int or errors. 
func ReadInt() (i int, e error) {
	input := Trim(readString())
	return strconv.Atoi(input)
}

func readString() string {
	input, err := reader.ReadString('\n')
	if err != nil {
		NewLinef("ReadString error: %v", err)
	}
	return input
}

// SimpleParse parses input from the user. Currently only one or two (space-separated) strings can
// be parsed. SimpleParse returns the first string as an action (if recognised), and the second
// string (the target verb) as is.
func SimpleParse() (parser.Action, string) {
	input := readString()
	split := strings.Split(input, " ")
	if len(split) >= 2 {
		return parser.ParseInput(split[0], Trim(split[1]))
	} else if len(split) == 1 {
		return parser.ParseInput(Trim(split[0]), "")
	}
	return parser.Unknown(), ""
}

const linuxCutset = "\n"
const windowsCutset = "\r" + linuxCutset

// Trim returns a string with spaces to the right trimmed, and a line ending.
func Trim(input string) string {
	return strings.TrimRight(input, windowsCutset)
}
