package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/things"
)

func ExecuteLookCommand(lookTarget string, gameState state.State) (target *things.Thing) {
	defer func() {
		if target != nil {
			io.NewLine(target.LookText)
		}
	}()

	if lookTarget == "" {
		io.NewLine(gameState.CurrentArea.Look)
	}

	target = gameState.CurrentArea.CheckAreaItemsForThing(lookTarget)
	if target != nil {
		return
	}

	target = gameState.CurrentArea.CheckAreaFeaturesForThing(lookTarget)
	if target != nil {
		return
	}

	target = gameState.CurrentArea.CheckAreaBeingsForThing(lookTarget)
	if target != nil {
		return
	}

	io.NewLinef("Couldn't find a %v to look at!", lookTarget)
	return
}

