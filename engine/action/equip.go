package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/things"
)

// ExecuteEquipCommand takes an equip target string and a game state. If there is an equippable item in
// the Inventory attached to the game state, with a name matching the target,  it will be moved to the
// equipped Items.
func ExecuteEquipCommand(equipTarget string, state *state.State) {
	item, err := state.Inventory.GetItemWithName(equipTarget)
	if err == nil {
		itemInterface := *item
		_, ok := itemInterface.(things.Equippable)
		if ok {
			item, err := state.Inventory.RemoveItemWithName(equipTarget)
			if item != nil && err == nil {
				state.EquippedItems.StoreItem(*item)
				io.ActiveInputOutputHandler.NewLinef("You equipped the %v.", equipTarget)
			} else {
				io.ActiveInputOutputHandler.NewLinef("Failed to equip item... %v", err)
			}
		} else {
			io.ActiveInputOutputHandler.NewLinef("How do you expect to equip the %v?", equipTarget)
		}
	} else {
		io.ActiveInputOutputHandler.NewLinef("Do not have a %v to equip.", equipTarget)
	}
}
