package main

import (
	"gostories/engine"
	"gostories/gen/items"
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
	collar := items.Item_Collar
	collar.Hide()
	storeRoom.Items = append(storeRoom.Items, collar)

	stage := engine.Stage{}
	stage.Start(catRoom)
}

// TODO move area construction out of main
func catRoom() things.Area {
	return things.Area{
		Look:   "You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		Beings: []things.Being{things.NewBubbles()},
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
		Features: []things.Feature{
			items.Feature_Shelf,
		},
	}
}
