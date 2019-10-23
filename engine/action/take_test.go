package action

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/store"
	"gostories/gen/items"
	tutils "gostories/utils/testing"
)

func TestTakeCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testGameState := tutils.TestState()
	testArea := testGameState.CurrentArea

	t.Run("item added to inventory and removed from area", func(t *testing.T) {
		testItem := items.Get("sardines")
		testItem.GetThing().Visible = true
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
