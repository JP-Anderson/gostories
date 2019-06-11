package parser

func ParseInput(arg1, arg2 string) (Action, string)  {
	return actionFromString(arg1), arg2
}

type Action struct { Name string }
func actionFromString(in string) Action {
	action, found := actions[in]; if found {
		return action
	}
	return UnknownAction
}

var actions = map[string]Action{
	"speak" : TalkAction,
	"talk" : TalkAction,
	"chat" : TalkAction,
}

var UnknownAction = Action{ "unknown" }
var TalkAction = Action{ "talk" }
