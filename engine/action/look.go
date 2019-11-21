package action

import (
	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/things"
	"gostories/things/area"
)

// ExecuteLookCommand is given the string name of a target, and a game State. It searches through the
// Items, Features, and Beings in the current area of the game State. If any "Thing" has a name matching
// the target, the Look text for the Thing is sent through the output handler, and the Thing is returned
// to run any optional post Action triggers attached to the Thing.
func ExecuteLookCommand(lookTarget string, gameState *state.State) (target *things.Thing) {
	if lookTarget == "" {
		io.Handler.NewLine(gameState.CurrentArea.LookText())
		return nil
	}
	target = gameState.CurrentArea.CheckAreaForThing(lookTarget, area.CheckBeings, area.CheckItems, area.CheckFeatures)
	if target != nil {
		io.Handler.NewLine(target.LookText)
		return target
	}
	io.Handler.NewLinef("Couldn't find a %v to look at!", lookTarget)
	return target
}
