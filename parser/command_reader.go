package parser

func ParseInput(arg1, arg2 string) (Action, string)  {
	return actionFromString(arg1), arg2
}

type Action struct { Name string }
func actionFromString(in string) Action {
	action, found := actions[in]; if found {
		return action
	}
	return unknownAction
}

var actions = map[string]Action{
	"speak" :  talkAction,
	"talk" :   talkAction,
	"chat" :   talkAction,
	"quit" :   quitAction,
	"exit" :   travelAction,
	"travel" : travelAction,
}

var unknownAction = Action{"unknown" }
var talkAction = Action{"talk" }
var quitAction = Action {"quit" }
var travelAction = Action { "travel" }

func Unknown() Action {
	return unknownAction
}