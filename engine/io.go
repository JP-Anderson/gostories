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

var Reader = bufio.NewReader(os.Stdin)

func SimpleParse(input string) (parser.Action, string) {
	split := strings.Split(input, " ")
	//debug
	NewLine("0 -> " + split[0])
	NewLine("1 -> " + split[1])
	return parser.ParseInput(split[0], strings.TrimRight(split[1], "\r\n"))
}
