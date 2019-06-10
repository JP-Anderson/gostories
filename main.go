package main

import (
	"gostories/engine"
	"gostories/things"
)

func main() {
	context := engine.Context{
		CurrentArea: catRoom(),
	}
	engine.NewLine(context.CurrentArea.Look)
	in, err := engine.Reader.ReadString('\n')
	if err != nil {

	}
	engine.SimpleParse(in)
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