package logic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/things/area"
	tutils "gostories/utils/testing"
)

func TestMultiTriggerString(t *testing.T) {
	input := "add-exit(north,cat_room);add-exit(south,kitchen)"

	testGameState := tutils.TestState()
	testArea := &area.Area{
		Look: "An area",
		Exits: make(map[area.Direction]area.Exit),
	}
	testGameState.CurrentArea = testArea

	err := EvaluateTrigger(testGameState, input)
	assert.NoError(t, err)
	exitNorth := testGameState.CurrentArea.Exits["north"]
	exitSouth := testGameState.CurrentArea.Exits["south"]

	assert.Equal(t, exitNorth.From, testArea)
	assert.Equal(t, exitSouth.From, testArea)
	assert.True(t, strings.Contains(exitNorth.To.Look, "a small room"))
	assert.True(t, strings.Contains(exitSouth.To.Look, "kitchen"))
}

func TestAddExit(t *testing.T) {
	input := "add-exit(north,cat_room)"
	testArea := &area.Area{
		Look:  "Some room",
		Exits: make(map[area.Direction]area.Exit),
	}
	testGameState := tutils.TestState()
	testGameState.CurrentArea = testArea

	err := EvaluateTrigger(testGameState, input)
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

func TestChangeLookText(t *testing.T) {
	testGameState := tutils.TestState()
	testGameState.CurrentArea.LookTexts = []string{
		"one",
		"two",
	}
	assert.NotEqual(t, "two", testGameState.CurrentArea.LookText())	
	err := EvaluateTrigger(testGameState, "change-look-text(1)")
	assert.NoError(t, err)
	assert.Equal(t, "two", testGameState.CurrentArea.LookText())
}