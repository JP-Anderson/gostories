package runner

import (
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/engine/speech"
	"gostories/engine/state"
)

// RunWithAlt takes two speech.Trees, if the primary speech passes a specific condition in the root Tree node, it
// will begin parsing the Tree with Run. Otherwise it will parse the alternate Tree. This can be used to design
// alernate major speech paths based on conditions in the game (e.g. player is wearing/holding an item).
func RunWithAlt(speech *speech.Tree, alt *speech.Tree, gameState state.State) {
	ran := Run(speech, gameState)
	if ran {
		return
	}
	if alt != nil {
		Run(alt, gameState)
	}
}

// Run takes a speech.Tree and a game State, starting from the first node of the SpeechTree, it will parse the speech
// tree based on input/output to/from the player.
func Run(tree *speech.Tree, gameState state.State) bool {
	p := tree.Start()
	curr := &p
	onFirstRun := true
	for {
		if curr == nil {
			return false
		}
		if curr.Condition != "" {
			if !logic.EvaluateCondition(gameState, curr.Condition) {
				if onFirstRun {
					return false
				}
				return true
			}
		}
		onFirstRun = false
		io.Handler.NewLine(curr.Speech)
		if curr.Trigger != "" {
			io.Handler.NewLine(curr.Trigger)
			err := logic.EvaluateTrigger(&gameState, curr.Trigger)
			if err != nil {
				io.Handler.NewLinef("%v", err)
			}
		}
		if curr.Checkpoint != nil {
			io.Handler.NewLinef("New checkpoint!")
			tree.SetStart(curr.Checkpoint)
		}
		if curr.Responses != nil && len(curr.Responses) > 0 {
			choice := printResponsesAndGetChoice(curr, gameState)
			response := curr.Responses[choice]
			io.Handler.NewLine(response.ResponseStr)
			if response.Trigger != "" {
				io.Handler.NewLine(response.Trigger)
				err := logic.EvaluateTrigger(&gameState, response.Trigger)
				if err != nil {
					io.Handler.NewLinef("%v", err)
				}
			}
			curr = &curr.Responses[choice].Next
		} else if curr.Next != nil {
			curr = curr.Next
		} else {
			return true
		}
	}
}

func printResponsesAndGetChoice(speechEvent *speech.Event, gameState state.State) int {
	// Remove responses we don't pass the conditions for first.
	availableResponses := []speech.Response{}
	for _, x := range speechEvent.Responses {
		if x.Condition == "" || logic.EvaluateCondition(gameState, x.Condition) {
			availableResponses = append(availableResponses, x)
		}
	}

	speechEvent.Responses = availableResponses
	for i, option := range speechEvent.Responses {
		io.Handler.NewLinef("%v - \"%v\"", i, option.ResponseStr)
	}

	last := len(speechEvent.Responses) - 1
	for {
		selection, err := io.Handler.ReadInt()
		if err != nil {
			io.Handler.NewLinef("%v", err)
		}
		if selection < 0 || selection > last {
			io.Handler.NewLinef("Enter option number from %v to %v", 0, last)
		}
		return selection
	}
}
