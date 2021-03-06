package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleParse(t *testing.T) {

	t.Run("parse verb action and noun", func(t *testing.T) {
		action, strings := SimpleParse(func() string { return "talk person" })
		require.Equal(t, talkAction, action)
		require.Equal(t, []string{"person"}, strings)
	})

	t.Run("parse single verb action", func(t *testing.T) {
		action, strings := SimpleParse(func() string { return "quit" })
		require.Equal(t, quitAction, action)
		require.Equal(t, []string{""}, strings)
	})

	t.Run("parse verb action and nouns", func(t *testing.T) {
		action, strings := SimpleParse(func() string { return "put coffee on table" })
		require.Equal(t, placeAction, action)
		require.Equal(t, []string{"coffee", "table"}, strings)
	})
}

func TestParseTwoTokenInputWithValidActions(t *testing.T) {
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
			action, _ := parseTwoTokenInput(testCase.action, testCase.target)
			assertions.Equal(testCase.expectedAction, action)
		})
	}
}

func TestParseTwoTokenInputWithInvalidActionsGivesUnknownAction(t *testing.T) {
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
			action, _ := parseTwoTokenInput(testCase.action, testCase.target)
			assertions.Equal(unknownAction, action)
		})
	}
}

func TestParseInputWithMultiStrings(t *testing.T) {
	validInputs := []struct {
		testName        string
		strings         []string
		expectedAction  Action
		expectedTargets []string
	}{
		// Place/put actions
		{"", []string{"put", "coffee", "on", "table"}, placeAction, []string{"coffee", "table"}},
		{"", []string{"put", "the", "coffee", "on", "to", "the", "table"}, placeAction, []string{"coffee", "table"}},
		{"", []string{"place", "coffee", "on", "table"}, placeAction, []string{"coffee", "table"}},
		{"", []string{"p", "coffee", "on", "table"}, placeAction, []string{"coffee", "table"}},
		{"", []string{"put", "coffee", "on", "table"}, placeAction, []string{"coffee", "table"}},
		// Open/unlock actions
		{"", []string{"unlock", "door", "with", "key"}, unlockAction, []string{"door", "key"}},
		{"", []string{"unlock", "the", "door", "with", "key"}, unlockAction, []string{"door", "key"}},
		{"", []string{"open", "gate", "with", "amulet"}, unlockAction, []string{"gate", "amulet"}},
		// TODO: Ignore adjectives?
		{"", []string{"open", "the", "large", "gate", "with", "amulet"}, unlockAction, []string{"large", "gate", "amulet"}},
	}
	for _, testCase := range validInputs {
		t.Run(testCase.testName, func(t *testing.T) {
			assertions := require.New(t)
			action, targets := parseMultiTokenInput(testCase.strings...)
			assertions.Equal(testCase.expectedAction, action)
			assertions.Equal(testCase.expectedTargets, targets)
		})
	}
}
