package action

import (
	"strings"

	"gostories/engine/io"
	"gostories/engine/io/console"
	"gostories/engine/state"
	"gostories/things/area"
)

func ExecuteTravelCommand(travelTarget string, state *state.State) bool {
	trimmed := consoleio.Trim(strings.ToLower(travelTarget))
	exit, exists := state.CurrentArea.Exits[area.Direction(trimmed)]
	if exists {
		state.CurrentArea = exit.To
		return true
	}
	io.ActiveInputOutputHandler.NewLinef("Could not find an exit to the %v", trimmed)
	return false
}

