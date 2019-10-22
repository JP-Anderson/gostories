package action

import (
	"testing"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/state"
	"gostories/engine/store"
	"gostories/gen/items"
	"gostories/things/area"
)

func TestPlaceItemNotInInventory(t * testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}

	str := ""
	ExecutePlaceCommand("collar", &str, testGameState)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Do not have a collar to put anywhere.",
		1,
	)
}

func TestPlaceItemInvalidTarget(t * testing.T) {
        mockedIOHandler := mockio.NewMockInputOutputHandler()
        io.ActiveInputOutputHandler = mockedIOHandler

        testArea := &area.Area{}
        testGameState := &state.State{
                CurrentArea:   testArea,
                Inventory:     store.NewInventory(),
                EquippedItems: store.NewEquippedItems(),
        }

	str := "lion"
	testGameState.Inventory.StoreItem(items.Item("collar"))
        ExecutePlaceCommand("collar", &str, testGameState)
        mockedIOHandler.ExpectedStringEqualsNthOutputString(
                t,
                "Not sure how to place the collar on the lion!",
                1,
        )
}


// TODO: test place being and feature. Test place with valid targets

