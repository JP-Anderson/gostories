package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/things"
)

// ExecuteLookCommand is given the string name of a target, and a game State. It searches through the
// Items, Features, and Beings in the current area of the game State. If any "Thing" has a name matching
// the target, the Look text for the Thing is sent through the output handler, and the Thing is returned
// to run any optional post Action triggers attached to the Thing.
func ExecuteLookCommand(lookTarget string, gameState *state.State) (target *things.Thing) {
	defer func() {
		if target != nil {
			io.ActiveInputOutputHandler.NewLine(target.LookText)
		}
	}()

	if lookTarget == "" {
		io.ActiveInputOutputHandler.NewLine(gameState.CurrentArea.Look)
		return
	}

	item := gameState.CurrentArea.FindItemByName(lookTarget)
	if item != nil {
		return item.GetThing()
	}

	target = gameState.CurrentArea.CheckAreaFeaturesForThing(lookTarget)
	if target != nil {
		return
	}

	target = gameState.CurrentArea.CheckAreaBeingsForThing(lookTarget)
	if target != nil {
		return
	}

	io.ActiveInputOutputHandler.NewLinef("Couldn't find a %v to look at!", lookTarget)
	return
}
