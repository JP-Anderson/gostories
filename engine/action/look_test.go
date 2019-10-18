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

func TestLookCommandWithValidTarget(t *testing.T) {
	mockedIOHandler := mockio.NewMockInputOutputHandler()
	io.ActiveInputOutputHandler = mockedIOHandler

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea: testArea,
	}

	t.Run("valid item target", func(t *testing.T) {
		testItem := items.Items["sardines"]
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

	testArea := &area.Area{}
	testGameState := &state.State{
		CurrentArea: testArea,
	}
	testArea.Items = store.NewItemStore()

	result := ExecuteLookCommand("sardines", testGameState)
	assert.Nil(t, result)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Couldn't find a sardines to look at!",
		1,
	)
}
