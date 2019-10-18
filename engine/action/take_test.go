package action

import (
	"gostories/engine/store"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/state"
	"gostories/gen/items"
	"gostories/things/area"
)

func TestTakeCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea: testArea,
		Inventory:   store.NewInventory(),
	}

	t.Run("item added to inventory and removed from area", func(t *testing.T) {
		testItem := items.Items["sardines"]
		testArea.Items = store.NewItemStore()
		testArea.Items.StoreItem(testItem)
		ExecuteTakeCommand("sardines", testGameState)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"You take the sardines.",
			1,
		)
		assert.Equal(t, 1, testGameState.Inventory.Size())
		assert.Equal(t, 0, testArea.Items.Size())
	})

}
