package main

import (
	"gostories/engine"
	"gostories/generator"
	"gostories/things"
)

func main() {
	catRoom := catRoom()
	storeRoom := storeRoom()

	// Add Exits
	catRoomToStockRoomExit := things.Exit{
		To:   &storeRoom,
		From: &catRoom,
	}
	storeRoomToCatRoomExit := things.Exit{
		To:   &catRoom,
		From: &storeRoom,
	}
	catRoom.Exits[things.West] = catRoomToStockRoomExit
	storeRoom.Exits[things.East] = storeRoomToCatRoomExit

	// Add item
	collar := things.CatCollarItem{}
	collar.Hide()
	storeRoom.Items = append(storeRoom.Items, collar)

	// TODO add some sort of feature which needs to be looked at to reveal item

	stage := engine.Stage{}
	stage.Start(catRoom)
}

// TODO move area construction out of main
func catRoom() things.Area {
	return things.Area{
		Look:   "You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		Beings: []things.Being{*cat()},
		Exits:  make(map[things.Direction]things.Exit),
		Items:  []things.Item{},
	}
}

func storeRoom() things.Area {
	return things.Area{
		Look:   "You are in some kind of stockroom. There is one shelf stacked high against one wall, across from the entrance.",
		Beings: []things.Being{},
		Exits:  make(map[things.Direction]things.Exit),
		Items:  []things.Item{},
	}
}

// Simple examples constructed from code for now
// Will either construct these from XML in future or find some way to autogenerate
func cat() *things.Being {
	translatedSpeech := generator.SpeechFromXMLFile("./generator/speech_data/bubbles_human.xml")
	catSpeech := generator.SpeechFromXMLFile("./generator/speech_data/bubbles.xml")
	return &things.Being{
		Name:      "Bubbles",
		Species:   "Cat",
		Speech:    translatedSpeech,
		AltSpeech: &catSpeech,
	}
}
