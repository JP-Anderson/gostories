package action

import (
	"gostories/engine/io"
	"gostories/engine/speech/runner"
	"gostories/engine/state"
	"strings"
)

// ExecuteTalkCommand takes a being target string, and a game state struct. If the current
// area of the game State has a Being with a name matching the provided target string, the player will
// initiate conversation with that Being, if possible.
func ExecuteTalkCommand(talkTarget string, state *state.State) {
	for _, being := range state.CurrentArea.Beings {
		io.Handler.NewLine(being.Name)
		if strings.ToLower(being.Name) == strings.ToLower(talkTarget) {
			io.Handler.NewLinef("You speak to %v.", being.Name)
			runner.RunWithAlt(&being.Speech, being.AltSpeech, *state)
			return
		}
	}
	io.Handler.NewLinef("Could not find a %v to talk to!", talkTarget)
}
