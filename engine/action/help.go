package action

import (
	"gostories/engine/io"
	"gostories/engine/parser"
	"gostories/engine/state"
)

func ExecuteHelpCommand(gameState *state.State) {
	io.Handler.NewLine("Available actions:")
	for action, commands := range parser.Actions() {
		commandStr := ""
		for _, command := range commands {
			commandStr = commandStr + " " + command
		}
		io.Handler.NewLinef(" > %s -- %s", action, commandStr)
	}
}
