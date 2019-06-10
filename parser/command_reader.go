package parser

func ParseInput(in string) (Action, Noun)  {
	return actionFromString(in), nil
}

type Action struct { Name string }

type Noun interface {

}

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
