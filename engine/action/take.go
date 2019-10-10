package action

import (
	"strings"

	"gostories/engine/io"
	"gostories/engine/state"
)

// ExecuteTakeCommand given a target string will search the Item objects in the current game Area
// of the State. If any Item names match the target, they will be added to the player inventory.
func ExecuteTakeCommand(takeTarget string, state *state.State) {
	item := state.CurrentArea.FindItemByName(takeTarget)
	if item != nil && item.GetThing().Visible {
		io.ActiveInputOutputHandler.NewLinef("You take the %v.", item.GetName())
		state.Inventory.StoreItem(item)
		state.CurrentArea.Items.RemoveItem(item)
		return
	}

	feature := state.CurrentArea.CheckAreaFeaturesForThing(takeTarget)
	if feature != nil && strings.ToLower(feature.Name) == strings.ToLower(takeTarget) {
		io.ActiveInputOutputHandler.NewLinef("You can't really take the %v...", feature.Name)
		return
	}
	io.ActiveInputOutputHandler.NewLinef("Couldn't find a %v to pick up.", takeTarget)
}
