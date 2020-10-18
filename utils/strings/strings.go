package strings

import (
	"strings"
)

const (
	space         = " "
	emptyString   = ""
	linuxCutset   = "\n"
	windowsCutset = "\r" + linuxCutset
)

// ToIDString converts a string to a string which can be used for ID comparisons. All IDs are treated as lower case, and
// spaces are ignored, which will simplify storing, retrieving, and comparing IDs across the system.
func ToIDString(input string) string {
	return strings.ToLower(strings.Replace(input, space, emptyString, -1))
}

// Trim returns a string with spaces to the right trimmed, and a linux line ending.
func Trim(input string) string {
	return strings.TrimRight(input, linuxCutset)
}
