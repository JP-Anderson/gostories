package action

import (
	"gostories/engine/inventory"
	"testing"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/state"
	"gostories/gen/items"
	"gostories/things/area"
)

func TestEquipCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     inventory.NewInventory(),
		EquippedItems: inventory.NewEquippedItems(),
	}

	t.Run("valid item target", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.ItemCollar)
		ExecuteEquipCommand("collar", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"You equipped the collar.",
			1,
		)
	})
}

func TestEquipCommandWithInvalidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     inventory.NewInventory(),
		EquippedItems: inventory.NewEquippedItems(),
	}

	t.Run("missing item", func(t *testing.T) {
		ExecuteEquipCommand("collar", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"Do not have a collar to equip.",
			1,
		)
	})

	t.Run("non-equippable item", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.ItemSardines)
		ExecuteEquipCommand("sardines", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"How do you expect to equip the sardines?",
			2,
		)
	})
}
