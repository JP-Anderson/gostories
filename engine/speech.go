package engine

import (
	"gostories/engine/io"
	"gostories/speech"
)

func RunWithAlt(speech speech.Tree, alt *speech.Tree, gameContext Context) {
	ran := Run(speech, gameContext)
	if ran {
		return
	}
	if alt != nil {
		Run(*alt, gameContext)
	}
}

func Run(speech speech.Tree, gameContext Context) bool {
	curr := &speech.Event
	onFirstRun := true
	for {
		if curr == nil {
			return false
		}
		if curr.Condition != "" {
			if !EvaluateCondition(gameContext, curr.Condition) {
				if onFirstRun {
					return false
				}
				return true
			}
		}
		onFirstRun = false
		io.NewLine(curr.Speech)
		if curr.Responses != nil && len(curr.Responses) > 0 {
			choice := printResponsesAndGetChoice(curr.Responses)
			io.NewLine(curr.Responses[choice].ResponseStr)
			curr = &curr.Responses[choice].Next
		} else if curr.Next != nil {
			curr = curr.Next
		}
		return true
	}
}

// TODO: Add Conditions to Responses
func printResponsesAndGetChoice(responseOptions []speech.Response) int {
	for i, option := range responseOptions {
		io.NewLinef("%v - \"%v\"", i, option.ResponseStr)
	}
	last := len(responseOptions)-1
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
