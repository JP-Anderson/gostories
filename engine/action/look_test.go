package action

import (
	"testing"

        "github.com/stretchr/testify/assert"

	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/things/area"
	"gostories/things"
	"gostories/gen/items"
)

func TestLookCommandWithValidTarget(t *testing.T) {
        testArea := &area.Area{}
        testGameState := &state.State {
                CurrentArea: testArea,
        }

	t.Run("valid item target", func(t *testing.T) {
		testItem := items.ItemSardines
		testArea.Items = []things.Item{
			testItem,
		}
		io.NewLinef("%#v", testGameState)
		io.NewLinef("%#v", testArea)
		io.NewLinef("%#v", testArea.Items)
		result := ExecuteLookCommand("sardines", *testGameState)
		assert.Equal(t, testItem.GetThing(), result)
		// TODO: create io mock/monkey patch for asserting text output of the look command.
		//  or find better way to test this.
	})

}

