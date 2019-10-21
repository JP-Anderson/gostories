package io

import (
	console "gostories/engine/io/console"
	"gostories/engine/parser"
)

// ActiveInputOutputHandler bridges the game engine with whatever Input/Output handler is in use. For unit
// testing, this variable is monkey patched. In the future, will review if this should be injected rather
// than monkey patched.
var ActiveInputOutputHandler InputOutputHandler = console.NewConsoleInputOutputHandler()

// InputOutputHandler specifies the behaviour a component must implement in order to pass input into
// the game engine, and receive output.
type InputOutputHandler interface {

	// NewLine outputs the provided string to the player. Or returns an error.
	NewLine(output string) error

	// NewLinef outputs the provided formatted string/args to the player. Or returns an error.
	NewLinef(output string, args ...interface{}) error

	// ReadInt returns an integer as input from the player, or an error.
	ReadInt() (int, error)

	// SimpleParse parses textual input from the player and returns an Action and a target
	// for the action in-game. It cannot error, but can return an UnknownAction.
	SimpleParse() (parser.Action, string)
}
