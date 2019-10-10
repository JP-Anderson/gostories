package action

import (
	"strings"

	"gostories/engine/io"
	consoleio "gostories/engine/io/console"
	"gostories/engine/state"
	"gostories/things/area"
)

// ExecuteTravelCommand takes a direction target string, and a game state struct. If the current
// area of the game State has an exit matching the provided direction, the player will travel to
// a new area via the exit. The function returns true if exit is valid and the player has travelled.
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
