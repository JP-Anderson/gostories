package action

import (
	"fmt"
	"strings"

	"gostories/engine/io"
	"gostories/engine/parser"
	"gostories/engine/state"
)

// ExecuteHelpCommand prints the available actions and their trigger commands to the player.
func ExecuteHelpCommand(gameState *state.State) {
	io.Handler.NewLine("Available actions:")
	for action, commands := range parser.Actions() {
		commandStr := ""
		for _, command := range commands {
			commandStr = commandStr + " " + command
		}
		io.Handler.NewLinef(
			" %s %s",
			padToLength(action, 12),
			commandStr,
		)
	}
}

func padToLength(s string, length int) string {
	padLength := length - len(s)
	return fmt.Sprintf("%s %s", s, strings.Repeat("-", padLength))
}
