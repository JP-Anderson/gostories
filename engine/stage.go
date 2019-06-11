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
	for {
		NewLine(s.context.CurrentArea.Look)
		action, noun := SimpleParse()
		if action.Name == "travel" {
			trimmed := trim(strings.ToLower(noun))
			exit, exists := s.context.CurrentArea.Exits[things.Direction(trimmed)]
			if exists {
				NewLinef("Found exit")
				s.context.CurrentArea = exit.To
			} else {
				NewLinef("Could not find an exit to the %v", noun)
			}
		} else if action.Name == "talk" {
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
		} else if action.Name == "exit" {
			break
		}  else {
			NewLine("Unknown action")
		}
	}
}
