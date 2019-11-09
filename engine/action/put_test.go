package action

import (
	"fmt"
	"testing"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/logic"
	"gostories/gen/features"
	"gostories/gen/items"
	"gostories/things/area"
	tutils "gostories/utils/testing"
)

func TestPlaceItemNotInInventory(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testGameState := tutils.TestState()

	str := ""
	ExecutePlaceCommand("collar", &str, testGameState)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Do not have a collar to put anywhere.",
		1,
	)
}

func TestPlaceItemInvalidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testGameState := tutils.TestState()

	str := "lion"
	testGameState.Inventory.StoreItem(items.Get("collar"))
	ExecutePlaceCommand("collar", &str, testGameState)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Not sure how to place the collar on the lion!",
		1,
	)
}

func TestPlaceShrubberyOnStand(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	// TODO: make a test area builder method
	testArea := &area.Area{
		Look: "some area",
		Exits: map[area.Direction]area.Exit{},
	}

	testGameState := tutils.TestState()
	shrubbery := items.Get("shrubbery")
	testGameState.Inventory.StoreItem(shrubbery)
	testGameState.CurrentArea = testArea

	stand := features.Get("stand")
	standName := stand.GetName()
	testGameState.CurrentArea.AddFeature(stand)

	println(fmt.Sprintf("%#v", stand.GetThing().Triggers["put"]))

	ExecutePlaceCommand("shrubbery", &standName, testGameState)
	
	trigger := stand.GetThing().Triggers["put"]

	logic.EvaluateTrigger(testGameState, trigger.String())
	
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"placed shrubbery on stand",
		1,
	)
}

// TODO: test place being and feature.
