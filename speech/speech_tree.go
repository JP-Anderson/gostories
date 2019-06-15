package speech

import (
	"gostories/engine/io"
)

type Tree struct {
	Event Event `xml:"Event"`
}

// An Event should have either a single next speech Event in `Next`. Or
// it should have 2 or more Response structs in `Responses` with user
// selectable options leading to speech Events.
type Event struct {
	// optional
	Next      *Event `xml:"Event"`
	// optional
	Responses []Response `xml:"Responses>Response"`
	Speech    string `xml:"Speech"`
}

type Response struct {
	ResponseStr string `xml:"ResponseStr"`
	Next        Event `xml:"Event"`
}

func Run(speech Tree) {
	curr := &speech.Event
	for {
		if curr == nil {
			io.NewLine("immediately breaking")
			break
		}
		io.NewLine(curr.Speech)
		if curr.Responses != nil && len(curr.Responses) > 0 {
			choice := printResponsesAndGetChoice(curr.Responses)
			curr = &curr.Responses[choice].Next
		} else if curr.Next != nil {
			curr = curr.Next
		} else {
			break
		}
	}
}

func printResponsesAndGetChoice(responseOptions []Response) int {
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
