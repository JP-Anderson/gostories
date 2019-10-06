package state

import (
	"gostories/engine/inventory"
	"gostories/things"
)

// State contains the current context of the game state, including the current area the player is in,
// and the players carried items.
type State struct {
	CurrentArea   things.Area
	Inventory     *inventory.Inventory
	EquippedItems *inventory.EquippedItems
}

