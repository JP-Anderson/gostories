package engine

import (
	"gostories/engine/state"
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/speech"
)

// RunWithAlt takes two speech.Trees, if the primary speech passes a specific condition in the root Tree node, it
// will begin parsing the Tree with Run. Otherwise it will parse the alternate Tree. This can be used to design
// alernate major speech paths based on conditions in the game (e.g. player is wearing/holding an item).
func RunWithAlt(speech speech.Tree, alt *speech.Tree, gameState state.State) {
	ran := Run(speech, gameState)
	if ran {
		return
	}
	if alt != nil {
		Run(*alt, gameState)
	}
}

// Run takes a speech.Tree and a game State, starting from the first node of the SpeechTree, it will parse the speech
// tree based on input/output to/from the player.
func Run(speech speech.Tree, gameState state.State) bool {
	curr := &speech.Event
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
		io.NewLine(curr.Speech)
		if curr.Trigger != "" {
			io.NewLine(curr.Trigger)
			err := logic.EvaluateTrigger(gameState, curr.Trigger); if err != nil {
				io.NewLinef("%v", err)
			}
		}
		if curr.Responses != nil && len(curr.Responses) > 0 {
			choice := printResponsesAndGetChoice(curr, gameState)
			response := curr.Responses[choice]
			io.NewLine(response.ResponseStr)
			if response.Trigger != "" {
				io.NewLine(response.Trigger)
				err := logic.EvaluateTrigger(gameState, response.Trigger); if err != nil {
					io.NewLinef("%v", err)
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
		io.NewLinef("%v - \"%v\"", i, option.ResponseStr)
	}

	last := len(speechEvent.Responses) - 1
	for {
		selection, err := io.ReadInt()
		if err != nil {
			io.NewLinef("%v", err)
		}
		if selection < 0 || selection > last {
			io.NewLinef("Enter option number from %v to %v", 0, last)
		}
		return selection
	}
}
