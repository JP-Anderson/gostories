package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gostories/parser"
)

func NewLine(output string) {
	fmt.Println(output)
}

func NewLinef(output string, args ...interface{}) {
	NewLine(fmt.Sprintf(output, args...))
}

var reader = bufio.NewReader(os.Stdin)

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

func Trim(input string) string {
	return strings.TrimRight(input, windowsCutset)
}
