package testing

import (
	"gostories/engine/state"
	"gostories/engine/store"
	"gostories/things/area"
)

// TestState returns a game State for tests, which comes with an empty area, an (empty) inventory, and an
// (empty) store of equipped items.
func TestState() *state.State {
	return &state.State{
		CurrentArea:   &area.Area{},
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}
}
