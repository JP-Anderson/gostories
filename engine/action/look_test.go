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

func TestLookCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testGameState := tutils.TestState()
	testArea := testGameState.CurrentArea
	t.Run("valid item target", func(t *testing.T) {
		testItem := items.Get("sardines")
		testArea.Items = store.NewItemStore()
		testArea.Items.StoreItem(testItem)
		result := ExecuteLookCommand("sardines", testGameState)
		assert.Equal(t, testItem.GetThing(), result)
		mockedIOHandler.ExpectedStringEqualsNthOutputString(
			t,
			"A tin of tasty sardines preserved in olive oil.",
			1,
		)
	})
}

func TestLookCommandWithInvalidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testGameState := tutils.TestState()
	testArea := testGameState.CurrentArea
	testArea.Items = store.NewItemStore()

	result := ExecuteLookCommand("sardines", testGameState)
	assert.Nil(t, result)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Couldn't find a sardines to look at!",
		1,
	)
}
