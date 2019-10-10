package state

import (
	"gostories/engine/store"
	"gostories/things/area"
)

// State contains the current context of the game state, including the current area the player is in,
// and the players carried items.
type State struct {
	CurrentArea   *area.Area
	Inventory     *store.Inventory
	EquippedItems *store.EquippedItems
}
