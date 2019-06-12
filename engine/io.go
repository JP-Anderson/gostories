package engine

import (
	"bufio"
	"fmt"
	"os"
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

func SimpleParse() (parser.Action, string) {
	input, err := reader.ReadString('\n')
	if err != nil {
		// todo
	}
	split := strings.Split(input, " ")
	if len(split) >= 2 {
		return parser.ParseInput(split[0], trim(split[1]))
	} else if len(split) == 1 {
		return parser.ParseInput(trim(split[0]), "")
	}
	return parser.Unknown(), ""
}

const linuxCutset = "\n"
const windowsCutset = "\r" + linuxCutset

func trim(input string) string {
	return strings.TrimRight(input, windowsCutset)
}
