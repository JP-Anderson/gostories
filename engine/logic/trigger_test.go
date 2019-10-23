package logic

import (
	"gostories/engine/state"
	"gostories/engine/store"
	"gostories/things/area"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddExit(t *testing.T) {
	input := "add-exit(north,cat_room)"
	testArea := &area.Area{
		Look:  "Some room",
		Exits: make(map[area.Direction]area.Exit),
	}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}

	err := triggerAddExit(testGameState, input)
	assert.NoError(t, err)

	area1 := testGameState.CurrentArea

	exitNorth := area1.Exits["north"]
	assert.NotNil(t, exitNorth)

	addedRoom := exitNorth.To
	assert.Equal(
		t,
		"You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		addedRoom.Look,
	)

	reverseExit := addedRoom.Exits["south"]
	assert.NotNil(t, reverseExit)
	assert.Equal(t, area1, reverseExit.To)
}

// func triggerAddExit(gameState state.State, input string) error {
// 	stringSlice := strings.Split(input,",")
// 	a := areas.Get(stringSlice[1])
// 	if a != nil {
// 		dir := area.StringToDirection[stringSlice[0]]
// 		exit := area.Exit{
// 			To: a,
// 			From: gameState.CurrentArea,
// 		}
// 		gameState.CurrentArea.Exits[dir] = exit
// 	}
// 	return nil
// }
