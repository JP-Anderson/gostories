package action

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/state"
	"gostories/engine/store"
	"gostories/gen/items"
	"gostories/things/area"
)

func TestEquipCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}

	t.Run("valid item target", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.Items["collar"])
		assert.Equal(t, 1, testGameState.Inventory.Size())
		ExecuteEquipCommand("collar", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"You equipped the collar.",
			1,
		)
		assert.Equal(t, 0, testGameState.Inventory.Size())
		assert.Equal(t, 1, testGameState.EquippedItems.Size())
	})
}

func TestEquipCommandWithInvalidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea:   testArea,
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}

	t.Run("missing item", func(t *testing.T) {
		ExecuteEquipCommand("collar", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"Do not have a collar to equip.",
			1,
		)
		assert.Equal(t, 0, testGameState.Inventory.Size())
		assert.Equal(t, 0, testGameState.EquippedItems.Size())
	})

	t.Run("non-equippable item", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.Items["sardines"])
		ExecuteEquipCommand("sardines", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"How do you expect to equip the sardines?",
			2,
		)
		assert.Equal(t, 1, testGameState.Inventory.Size())
		assert.Equal(t, 0, testGameState.EquippedItems.Size())
	})
}
