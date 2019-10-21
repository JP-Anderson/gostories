package parser

// ParseInput takes two strings from the user input (which will already have been split on spaces). It attempts to match
// the first arg to an Action, and will return any matching Action (or an UnknownAction if it cannot find a match).
// Currently the second argument is simply returned as is.
func ParseInput(arg1, arg2 string) (Action, string) {
	return actionFromString(arg1), arg2
}

func parseMultiTokenInput(ts ...string) (action Action, targets []string) {
	targets = []string{}
	action = unknownAction
	for _, token := range ts {
		println(token)
		_, isPreposition := prepositions[token]
		newAction := actionFromString(token)

		if !isPreposition {
			println("not preposition")
			if newAction == unknownAction {
				println("not an action either")
				targets = append(targets, token)
			} else {
				println("is an action")
				action = newAction
			}
		}
	}
	return action, targets
}

// Action is a type representing an action the player can execute in the game. Currently it just wraps a string
// which matches the verb the player types to carry out the action.
type Action struct{ Name string }

func actionFromString(in string) Action {
	action, found := actions[in]
	if found {
		return action
	}
	return unknownAction
}

var actions = map[string]Action{
	"speak": talkAction,
	"talk":  talkAction,
	"chat":  talkAction,
	"t":     talkAction,

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

	"quit": quitAction,
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
var quitAction = Action{"quit"}

// Unknown returns an unknownAction, which is used when user input cannot be parsed.
func Unknown() Action {
	return unknownAction
}
