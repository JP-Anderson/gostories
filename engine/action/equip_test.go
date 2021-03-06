package action

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mockio "gostories/engine/io/mock"
	"gostories/gen/items"
	tutils "gostories/utils/testing"
)

func TestEquipCommandWithValidTarget(t *testing.T) {
	mockedHandler := mockio.NewMockHandler()
	testGameState := tutils.TestState()

	t.Run("valid item target", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.Get("collar"))
		assert.Equal(t, 1, testGameState.Inventory.Size())
		ExecuteEquipCommand("collar", testGameState)
		mockedHandler.ExpectedStringEqualsNthOutputString(
			t,
			"You equipped the collar.",
			1,
		)
		assert.Equal(t, 0, testGameState.Inventory.Size())
		assert.Equal(t, 1, testGameState.EquippedItems.Size())
	})
}

func TestEquipCommandWithInvalidTarget(t *testing.T) {
	mockedHandler := mockio.NewMockHandler()
	testGameState := tutils.TestState()

	t.Run("missing item", func(t *testing.T) {
		ExecuteEquipCommand("collar", testGameState)
		mockedHandler.ExpectedStringEqualsNthOutputString(
			t,
			"Do not have a collar to equip.",
			1,
		)
		assert.Equal(t, 0, testGameState.Inventory.Size())
		assert.Equal(t, 0, testGameState.EquippedItems.Size())
	})

	t.Run("non-equippable item", func(t *testing.T) {
		testGameState.Inventory.StoreItem(items.Get("sardines"))
		ExecuteEquipCommand("sardines", testGameState)
		mockedHandler.ExpectedStringEqualsNthOutputString(
			t,
			"How do you expect to equip the sardines?",
			2,
		)
		assert.Equal(t, 1, testGameState.Inventory.Size())
		assert.Equal(t, 0, testGameState.EquippedItems.Size())
	})
}
