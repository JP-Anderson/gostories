package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewLine(output string) {
	fmt.Println(output)
	fmt.Print()
}

var Reader = bufio.NewReader(os.Stdin)

func SimpleParse(input string) {
	split := strings.Split(input, " ")
	for _, s := range split {
	    NewLine(s)
	}
}