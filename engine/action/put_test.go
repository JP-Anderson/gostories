package action

import (
	"testing"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/gen/items"
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

// TODO: test place being and feature. Test place with valid targets
