package logic

import "strings"

func parseFuncParam(input string) string {
	return input[strings.Index(input, "(")+1 : strings.Index(input, ")")]
}

func parseFuncName(input string) string {
	return input[:strings.Index(input, "(")]
}
