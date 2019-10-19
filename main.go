package main

import (
	"gostories/engine"
	"gostories/engine/store"
	"gostories/gen/features"
	"gostories/gen/items"
	"gostories/things"
	"gostories/things/area"
)

func main() {
	catRoom := catRoom()
	storeRoom := storeRoom()
	kitchen := kitchenRoom()

	// Add Exits
	catRoomToStockRoomExit := area.Exit{
		To:   &storeRoom,
		From: &catRoom,
	}
	storeRoomToCatRoomExit := area.Exit{
		To:   &catRoom,
		From: &storeRoom,
	}
	storeRoomToKitchenExit := area.Exit{
		To:   &kitchen,
		From: &storeRoom,
	}
	kitchenToStoreRoomExit := area.Exit{
		To:   &storeRoom,
		From: &kitchen,
	}
	catRoom.Exits[area.West] = catRoomToStockRoomExit
	storeRoom.Exits[area.East] = storeRoomToCatRoomExit
	storeRoom.Exits[area.North] = storeRoomToKitchenExit
	kitchen.Exits[area.South] = kitchenToStoreRoomExit

	// TODO: Perhaps also autogenerate the code for adding Items to Areas (currently only the reveal
	// trigger on a Feature is autogenerated.)
	// Add items
	storeRoom.Items.StoreItem(items.Item("collar"))
	kitchen.Items.StoreItem(items.Item("sardines"))

	stage := engine.Stage{}
	stage.Start(catRoom)
}

// TODO move area construction out of main
// TODO NewArea builder which creates empty stores etc
func catRoom() area.Area {
	return area.Area{
		Look:   "You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		Beings: []things.Being{things.NewBubbles()},
		Exits:  make(map[area.Direction]area.Exit),
		Items:  store.NewItemStore(),
	}
}

func storeRoom() area.Area {
	return area.Area{
		Look:   "You are in some kind of stockroom. There is one shelf stacked high against one wall, across from the entrance.",
		Beings: []things.Being{},
		Exits:  make(map[area.Direction]area.Exit),
		Items:  store.NewItemStore(),
		Features: []things.Feature{
			features.FeatureShelf,
		},
	}
}

func kitchenRoom() area.Area {
	return area.Area{
		Look:   "You are in a cramped kitchen, there is a fridge on the far side of the wall and one cupboard.",
		Beings: []things.Being{},
		Exits:  make(map[area.Direction]area.Exit),
		Items:  store.NewItemStore(),
		Features: []things.Feature{
			features.FeatureFridge,
		},
	}
}
