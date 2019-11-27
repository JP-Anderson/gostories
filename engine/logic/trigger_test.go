package logic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	mockio "gostories/engine/io/mock"
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

func TestChangeLookTextParameters(t *testing.T) {
	io := mockio.NewMockHandler()
	expectedOutputText0 := "Print this"
	expectedOutputText1 := "Some other text"
	inputs := []string{
		fmt.Sprintf("0, %s", expectedOutputText0),
		fmt.Sprintf("1, %s", expectedOutputText1),
	}
	testGameState := tutils.TestState()

	for _, input := range inputs {
		err := EvaluateTrigger(testGameState, fmt.Sprintf("change-look-text(%s)", input))
		assert.NoError(t, err)
	}
	io.ExpectedStringEqualsNthOutputString(
		t,
		expectedOutputText0,
		1,
	)
	io.ExpectedStringEqualsNthOutputString(
		t,
		expectedOutputText1,
		2,
	)
}

func TestChangeLookTextWithCommasInText(t *testing.T) {
	io := mockio.NewMockHandler()
	input := "1, hello, this is a test."
	testGameState := tutils.TestState()

	assert.NoError(t, EvaluateTrigger(testGameState, fmt.Sprintf("change-look-text(%s)", input)))
	io.ExpectedStringEqualsNthOutputString(t, "hello, this is a test.", 1)
}

