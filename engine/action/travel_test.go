package action

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/things/area"
	tutils "gostories/utils/testing"
)

func TestTravelCommandWithValidTarget(t *testing.T) {
	testArea := &area.Area{}
	testArea2 := &area.Area{}

	area1To2Exit := area.Exit{
		From: testArea,
		To:   testArea2,
	}

	testArea.Exits = map[area.Direction]area.Exit{
		area.North: area1To2Exit,
	}

	testGameState := tutils.TestState()
	testGameState.CurrentArea = testArea

	t.Run("valid exit to north", func(t *testing.T) {
		result := ExecuteTravelCommand("north", testGameState)
		assert.True(t, result)
	})
}

func TestTravelCommandWithInvalidTarget(t *testing.T) {
	testGameState := tutils.TestState()
	testArea := testGameState.CurrentArea
	testArea.Exits = map[area.Direction]area.Exit{}

	t.Run("invalid exit to north", func(t *testing.T) {
		result := ExecuteTravelCommand("north", testGameState)
		assert.False(t, result)
	})
}
