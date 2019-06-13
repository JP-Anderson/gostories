package parser

func ParseInput(arg1, arg2 string) (Action, string) {
	return actionFromString(arg1), arg2
}

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

	"look":    lookAction,
	"examine": lookAction,
	"search":  lookAction,
	"scan":    lookAction,

	"exit":   travelAction,
	"walk":   travelAction,
	"travel": travelAction,

	"take": takeAction,
	"grab": takeAction,

	"inventory": inventoryAction,
	"bag":       inventoryAction,
	"pack":      inventoryAction,
	"i":         inventoryAction,

	"quit": quitAction,
}

var unknownAction = Action{"unknown"}
var talkAction = Action{"talk"}
var lookAction = Action{"look"}
var travelAction = Action{"travel"}
var takeAction = Action{"take"}
var inventoryAction = Action{"inventory"}
var quitAction = Action{"quit"}

func Unknown() Action {
	return unknownAction
}
