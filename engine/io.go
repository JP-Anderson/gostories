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
	// trim right for Windows. "\n" for Linux
	return parser.ParseInput(split[0], strings.TrimRight(split[1], "\r\n"))
}
