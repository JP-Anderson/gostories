package engine

import (
	"strings"

	"gostories/things"
)

type Stage struct {
	context Context
}

func (s Stage) Start(area things.Area) {
	s.context = Context{CurrentArea: area}
	newArea := true
	for {
		if newArea {
			NewLine(s.context.CurrentArea.Look)
			newArea = false
		}
		action, noun := SimpleParse()
		if action.Name == "look" {
			// TODO also Look at NPCs
			if noun == "" {
				NewLine(s.context.CurrentArea.Look)
			} else {
				found := false
				for _, item := range s.context.CurrentArea.Items {
					if strings.ToLower(item.GetName()) == strings.ToLower(noun) {
						found = true
						NewLine(item.GetLookText())
					}
				}
				if !found {
					NewLinef("Couldn't find a %v to look at!", noun)
				}
			}
		} else if action.Name == "travel" {
			trimmed := trim(strings.ToLower(noun))
			exit, exists := s.context.CurrentArea.Exits[things.Direction(trimmed)]
			if exists {
				s.context.CurrentArea = *exit.To
				newArea = true
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
		} else {
			NewLine("Unknown action")
		}
	}
}
