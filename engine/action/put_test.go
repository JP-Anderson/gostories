package action

import (
	"testing"

	mockio "gostories/engine/io/mock"
	"gostories/gen/features"
	"gostories/gen/items"
	tutils "gostories/utils/testing"
)

func TestPlaceItemNotInInventory(t *testing.T) {
	mockedHandler := mockio.NewMockHandler()

	testGameState := tutils.TestState()

	str := ""
	ExecutePlaceCommand("collar", &str, testGameState)
	mockedHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Do not have a collar to put anywhere.",
		1,
	)
}

func TestPlaceItemInvalidTarget(t *testing.T) {
	mockedHandler := mockio.NewMockHandler()

	testGameState := tutils.TestState()

	str := "lion"
	testGameState.Inventory.StoreItem(items.Get("collar"))
	ExecutePlaceCommand("collar", &str, testGameState)
	mockedHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Not sure how to place the collar on the lion!",
		1,
	)
}

func TestPlaceShrubberyOnStand(t *testing.T) {
	mockedHandler := mockio.NewMockHandler()

	testGameState := tutils.TestState()
	shrubbery := items.Get("shrubbery")
	testGameState.Inventory.StoreItem(shrubbery)
	stand := features.Get("stand")
	standName := stand.GetName()
	testGameState.CurrentArea.AddFeature(stand)

	ExecutePlaceCommand("shrubbery", &standName, testGameState)
	
	mockedHandler.ExpectedStringEqualsNthOutputString(
		t,
		"placed shrubbery on stand",
		1,
	)
}

// TODO: test place being and feature.
