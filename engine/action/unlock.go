package action

import (
	"strings"

	"gostories/engine/logic"
	"gostories/engine/state"
	"gostories/things/area"
)

// ExecuteUnlockCommand given a target string will trigger an unlock, open, or access command
// on some in game Object, if it exists.
func ExecuteUnlockCommand(unlockTarget string, state *state.State) {
     f := state.CurrentArea.CheckAreaForThing(unlockTarget, area.CheckFeatures)
     if f != nil && strings.ToLower(f.Name) == strings.ToLower(unlockTarget) {
     	// do something...
	trigger, ok := f.Triggers["unlock"]
	if !ok {
	   return
	}
	err := logic.EvaluateTrigger(state, trigger.Action)
	if err != nil {
	return
	}  
     }
}