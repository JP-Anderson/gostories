package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/things/area"
	tutils "gostories/utils/testing"
)

func TestAddExit(t *testing.T) {
	input := "add-exit(north,cat_room)"
	testArea := &area.Area{
		Look:  "Some room",
		Exits: make(map[area.Direction]area.Exit),
	}
	testGameState := tutils.TestState()
	testGameState.CurrentArea = testArea

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
