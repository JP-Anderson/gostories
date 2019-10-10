package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
)

// ExecuteEquipCommand takes an equip target string and a game state. If there is an equippable item in
// the Inventory attached to the game state, with a name matching the target,  it will be moved to the
// equipped Items.
func ExecuteEquipCommand(equipTarget string, state *state.State) {
	item, err := state.Inventory.GetItemWithName(equipTarget)
	if err != nil {
		io.ActiveInputOutputHandler.NewLinef("Do not have a %v to equip.", equipTarget)
		return
	}
	equipped := state.EquippedItems.StoreItem(*item)
	if equipped {
		state.Inventory.RemoveItem(*item)
	}
}
