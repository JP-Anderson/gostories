package engine

import (
	"strings"

	"gostories/things"
)

type Stage struct {
	context Context
}

func (s Stage) Start(area things.Area) {
	s.context = Context{ CurrentArea: area }
	NewLine(s.context.CurrentArea.Look)
	for {
		action, noun := SimpleParse()
		if action.Name == "exit" {
			break
		}
		if action.Name == "talk" {
			found := false
			for _, being := range s.context.CurrentArea.Beings {
				if strings.ToLower(being.Name) == strings.ToLower(noun) {
					found = true
					NewLine(being.Speech[0])
				}
			}
			if !found {
				NewLinef("Could not find a %v to talk to!", noun)
			}
		} else {
			NewLine("Unknown action")
		}
	}
}
