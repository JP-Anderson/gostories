package parser

import (
	"strings"

	gstrings "gostories/utils/strings"
)

// SimpleParse parses input from the user. Currently only one or two (space-separated) strings can
// be parsed. SimpleParse returns the first string as an action (if recognised), and the second
// string (the target verb) as is.
func SimpleParse(getStringFunc func() string) (Action, []string) {
	input := getStringFunc()
	split := strings.Split(input, " ")
	len := len(split)
	if len > 2 {
		return ParseInput(split...)
	} else if len == 2 {
		return ParseInput(gstrings.Trim(split[0]), gstrings.Trim(split[1]))
	} else if len == 1 {
		return ParseInput(gstrings.Trim(split[0]), "")
	}
	return Unknown(), []string{""}
}

// ParseInput takes a slice of strings from the user input (which will already have been split on spaces). It will pass
// into a token parsing function based on the number of tokens in the slice.
func ParseInput(tokens ...string) (Action, []string) {
	if len(tokens) == 1 {
		return parseTwoTokenInput(tokens[0], "")
	} else if len(tokens) == 2 {
		return parseTwoTokenInput(tokens[0], tokens[1])
	}
	return parseMultiTokenInput(tokens...)
}

func parseTwoTokenInput(t1, t2 string) (Action, []string) {
	return actionFromString(t1), []string{t2}
}

func parseMultiTokenInput(ts ...string) (action Action, targets []string) {
	targets = []string{}
	action = unknownAction
	for _, token := range ts {
		token := strings.TrimSuffix(strings.TrimSpace(token), "\n")
		_, isArticle := articles[token]
		if isArticle {
			continue
		}
		_, isConjunction := conjunctions[token]
		if isConjunction {
			continue
		}
		_, isPreposition := prepositions[token]
		if isPreposition {
			continue
		}
		newAction := actionFromString(token)
		if newAction == unknownAction {
			targets = append(targets, token)
		} else {
			action = newAction
		}
	}
	return action, targets
}

// Action is a type representing an action the player can execute in the game. Currently it just wraps a string
// which matches the verb the player types to carry out the action.
type Action struct {
	Name string
}

func actionFromString(in string) Action {
	action, found := actions[in]
	if found {
		return action
	}
	return unknownAction
}

// Actions returns the Action names mapped to a []string containing the commands to trigger the Action, primarily
// for printing the actions and commands in the Help command.
func Actions() map[string][]string {
	actionStrings := map[string][]string{}
	for command, a := range actions {
		name := a.Name
		_, ok := actionStrings[name]
		if !ok {
			actionStrings[name] = []string{command}
		} else {
			actionStrings[name] = append(actionStrings[name], command)
		}
	}
	return actionStrings
}

var actions = map[string]Action{
	"speak": talkAction,
	"talk":  talkAction,
	"chat":  talkAction,
	"t":     talkAction,

	// TODO: review splitting look and examine.
	"look":    lookAction,
	"examine": lookAction,
	"search":  lookAction,
	"scan":    lookAction,
	"l":       lookAction,

	"exit":   travelAction,
	"walk":   travelAction,
	"travel": travelAction,
	"go":     travelAction,
	"w":      travelAction,

	"take": takeAction,
	"grab": takeAction,

	"equip": equipAction,
	"wear":  equipAction,
	"hold":  equipAction,
	"e":     equipAction,

	"inventory": inventoryAction,
	"bag":       inventoryAction,
	"pack":      inventoryAction,
	"i":         inventoryAction,

	"put":   placeAction,
	"place": placeAction,
	"p":     placeAction,

	"unlock": unlockAction,
	"open":   unlockAction,
	"access": unlockAction,

	"help": helpAction,
	"h":    helpAction,

	"quit": quitAction,
	"q":    quitAction,
}

var unknownAction = Action{"unknown"}
var talkAction = Action{"talk"}
var lookAction = Action{"look"}
var travelAction = Action{"travel"}
var takeAction = Action{"take"}
var equipAction = Action{"equip"}
var inventoryAction = Action{"inventory"}
var placeAction = Action{"place"}
var unlockAction = Action{"unlock"}
var helpAction = Action{"help"}
var quitAction = Action{"quit"}

// Unknown returns an unknownAction, which is used when user input cannot be parsed.
func Unknown() Action {
	return unknownAction
}
