package main

import (
	"strings"

	"gostories/engine"
	"gostories/things"
)

func main() {
	area := catRoom()
	engine.NewLine(area.Beings[0].Name)
	context := engine.Context{
		CurrentArea: area,
	}
	engine.NewLine(context.CurrentArea.Look)

	// TODO: move this to engine
	for {
		action, noun := engine.SimpleParse()
		if action.Name == "exit" {
			break
		}
		if action.Name == "talk" {
			found := false
			for _, being := range context.CurrentArea.Beings {
				if strings.ToLower(being.Name) == strings.ToLower(noun) {
					found = true
					engine.NewLine(being.Speech[0])
				}
			}
			if !found {
				engine.NewLinef("Could not find a %v to talk to!", noun)
			}
		} else {
			engine.NewLinef("Unknown action %v", action.Name)
		}
	}
}

func catRoom() things.Area {
	return things.Area{
		Look: "You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		Beings: []things.Being{ *cat() },
		Exits: []things.Exit{},
	}
}

func storeRoom() things.Area {
	return things.Area{
		Look: "You are in some kind of stockroom. There is one shelf stacked high against one wall, across from the entrance.",
		Beings: []things.Being{},
		Exits: []things.Exit{ },
	}
}

// Simple examples constructed from code for now
// Will either construct these from XML in future or find some way to autogenerate
func cat() *things.Being {
	return &things.Being{
		Name: "Bubbles",
		Species: "Cat",
		Speech: []string{
			"Meeeeeeeeow!",
			"...mew!",
			"Rrrrrrewwwwwowow!",
		},
	}
}