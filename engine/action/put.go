package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
)

// ExecutePlaceCommand can take one or two string targets. The first target must always match to an item that can
// be placed in the players inventory. If no second target is provided, the plan is to allow the player to drop items
// into the area with this command (TODO).
// If the second target is provided, certain interactions could occur of this, depending on how the place command will
// affect different targets (and also different items) (TODO)
func ExecutePlaceCommand(placeTarget string, placeSecondTarget *string, gameState *state.State) {
	being := gameState.CurrentArea.CheckAreaBeingsForThing(placeTarget)
	if being != nil {
		io.ActiveInputOutputHandler.NewLine("How do you expect to place " + being.Name + "!?")
		return
	}

	ftr := gameState.CurrentArea.CheckAreaFeaturesForThing(placeTarget)
	if ftr != nil {
		io.ActiveInputOutputHandler.NewLine("How do you expect to place the " + ftr.Name + "!?")
		return
	}

	item, err := gameState.Inventory.GetItemWithName(placeTarget)
	if err != nil {
		item, err = gameState.EquippedItems.GetItemWithName(placeTarget)
		if err != nil {
			io.ActiveInputOutputHandler.NewLinef("Do not have a %s to put anywhere.", placeTarget)
			return
		}
	}

	actualItem := *item

	if placeSecondTarget == nil {
		io.ActiveInputOutputHandler.NewLinef("Are you sure you want to drop the %s?", placeTarget)
		return
	}

	secondTargetString := *placeSecondTarget

	// TODO run triggers off certain action interactions
	being2 := gameState.CurrentArea.CheckAreaBeingsForThing(secondTargetString)
	if being2 != nil {
		io.ActiveInputOutputHandler.NewLinef("placed %s on %s", actualItem.GetName(), being2.Name)
		return
	}

	ftr2 := gameState.CurrentArea.CheckAreaFeaturesForThing(*placeSecondTarget)
	if ftr2 != nil {
		io.ActiveInputOutputHandler.NewLinef("placed %s on %s", actualItem.GetName(), ftr2.Name)
		return
	}

	io.ActiveInputOutputHandler.NewLinef("Not sure how to place the %s on the %s!", actualItem.GetName(), secondTargetString)

}
