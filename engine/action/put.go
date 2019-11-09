package action

import (
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/engine/state"
	"gostories/things"
	"gostories/things/area"
)

// ExecutePlaceCommand can take one or two string targets. The first target must always match to an item that can
// be placed in the players inventory. If no second target is provided, the plan is to allow the player to drop items
// into the area with this command (TODO).
// If the second target is provided, certain interactions could occur of this, depending on how the place command will
// affect different targets (and also different items) (TODO)
func ExecutePlaceCommand(placeTarget string, placeSecondTarget *string, gameState *state.State) *things.Thing {
	notAnItem := gameState.CurrentArea.CheckAreaForThing(placeTarget, area.CheckBeings, area.CheckFeatures)
	if notAnItem != nil {
		io.ActiveInputOutputHandler.NewLine("How do you expect to place the " + notAnItem.Name + "!?")
		return nil
	}

	item, err := gameState.Inventory.GetItemWithName(placeTarget)
	if err != nil {
		item, err = gameState.EquippedItems.GetItemWithName(placeTarget)
		if err != nil {
			io.ActiveInputOutputHandler.NewLinef("Do not have a %s to put anywhere.", placeTarget)
			return nil
		}
	}

	actualItem := *item

	if placeSecondTarget == nil {
		// TODO: does this need to be handled differently for triggers?
		io.ActiveInputOutputHandler.NewLinef("Are you sure you want to drop the %s?", placeTarget)
		return nil
	}

	secondTargetString := *placeSecondTarget
	secondTarget := gameState.CurrentArea.CheckAreaForThing(secondTargetString, area.CheckBeings, area.CheckFeatures)
	// TODO run triggers off certain action interactions
	if secondTarget != nil {
		io.ActiveInputOutputHandler.NewLinef("placed %s on %s", actualItem.GetName(), secondTarget.Name)
		trigger, ok := secondTarget.Triggers["put"]
		if ok {
			err = logic.EvaluateTrigger(gameState, trigger.Action)
			if err != nil {
				io.ActiveInputOutputHandler.NewLinef("trigger error %v", err)
			}
		}
		return secondTarget
	}

	io.ActiveInputOutputHandler.NewLinef("Not sure how to place the %s on the %s!", actualItem.GetName(), secondTargetString)
	return nil
}
