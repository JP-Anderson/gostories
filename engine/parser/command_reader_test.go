package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInputWithValidActions(t *testing.T) {

	validActionStrings := []struct {
		action         string
		target         string
		expectedAction Action
	}{
		{"speak", "person", talkAction},
		{"talk", "person", talkAction},
		{"chat", "person", talkAction},
		{"t", "person", talkAction},
		{"look", "person", lookAction},
		{"l", "person", lookAction},
		{"exit", "person", travelAction},
		{"w", "person", travelAction},
		{"go", "person", travelAction},
		{"travel", "person", travelAction},
		{"walk", "person", travelAction},
		{"take", "person", takeAction},
		{"grab", "person", takeAction},

		{"equip", "person", equipAction},
		{"wear", "person", equipAction},
		{"hold", "person", equipAction},
		{"e", "person", equipAction},

		{"inventory", "person", inventoryAction},
		{"bag", "person", inventoryAction},
		{"pack", "person", inventoryAction},
		{"i", "person", inventoryAction},
	}

	for _, testCase := range validActionStrings {
		t.Run(testCase.action, func(t *testing.T) {
			assertions := require.New(t)
			action, _ := ParseInput(testCase.action, testCase.target)
			assertions.Equal(testCase.expectedAction, action)
		})
	}
}

func TestParseInputWithInvalidActionsGivesUnknownAction(t *testing.T) {

	unknownActionInputs := []struct {
		action string
		target string
	}{
		{"exfoliate", "person"},
		{"tweet", "person"},
		{"rub", "person"},
		{"embezzle", "person"},
		{"snipe", "person"},
	}

	for _, testCase := range unknownActionInputs {
		t.Run(testCase.action, func(t *testing.T) {
			assertions := require.New(t)
			action, _ := ParseInput(testCase.action, testCase.target)
			assertions.Equal(unknownAction, action)
		})
	}
}
