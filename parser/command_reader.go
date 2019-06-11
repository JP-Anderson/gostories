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
	"speak" : talkAction,
	"talk" :  talkAction,
	"chat" :  talkAction,
	"quit" :  exitAction,
	"exit" :  exitAction,
}

var unknownAction = Action{"unknown" }
var talkAction = Action{"talk" }
var exitAction = Action {"exit" }

func Unknown() Action {
	return unknownAction
}