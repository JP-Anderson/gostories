package engine

import (
	"gostories/engine/context"
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/speech"
)

func RunWithAlt(speech speech.Tree, alt *speech.Tree, gameContext context.Context) {
	ran := Run(speech, gameContext)
	if ran {
		return
	}
	if alt != nil {
		Run(*alt, gameContext)
	}
}

func Run(speech speech.Tree, gameContext context.Context) bool {
	curr := &speech.Event
	onFirstRun := true
	for {
		if curr == nil {
			return false
		}
		if curr.Condition != "" {
			if !logic.EvaluateCondition(gameContext, curr.Condition) {
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
			err := logic.EvaluateTrigger(gameContext, curr.Trigger); if err != nil {
				io.NewLinef("%v", err)
			}
		}
		if curr.Responses != nil && len(curr.Responses) > 0 {
			choice := printResponsesAndGetChoice(curr, gameContext)
			response := curr.Responses[choice]
			io.NewLine(response.ResponseStr)
			if response.Trigger != "" {
				io.NewLine(response.Trigger)
				err := logic.EvaluateTrigger(gameContext, response.Trigger); if err != nil {
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

func printResponsesAndGetChoice(speechEvent *speech.Event, gameContext context.Context) int {
	// Remove responses we don't pass the conditions for first.
	availableResponses := []speech.Response{}
	for _, x := range speechEvent.Responses {
		if x.Condition == "" || logic.EvaluateCondition(gameContext, x.Condition) {
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
