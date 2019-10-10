package action

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/engine/state"
	"gostories/gen/items"
	"gostories/things"
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
		testItem := items.ItemSardines
		testArea.Items = []things.Item{
			testItem,
		}
		result := ExecuteLookCommand("sardines", *testGameState)
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

	result := ExecuteLookCommand("sardines", *testGameState)
	assert.Nil(t, result)
	mockedIOHandler.ExpectedStringEqualsNthOutputString(
		t,
		"Couldn't find a sardines to look at!",
		1,
	)
}
