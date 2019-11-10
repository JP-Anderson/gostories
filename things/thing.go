package things

import (
	"fmt"
	"strings"
)

// A Thing is anything in the game that can be interacted with in some way. All Things have
// a Name (identifying string), LookText (text which is output when the player executes the
// Look action on the Thing, a Visible attribute (is the Thing currently revealed to the
// player, and a set of Triggers. Triggers have effect on the game state, based on
// interactions the player makes with the object.
type Thing struct {
	Name     string
	Names []string
	LookText string
	Visible  bool
	// Triggers maps verb command strings to trigger functions. See trigger.go for trigger functions.
	Triggers map[string]Trigger
}

// A Trigger contains an Action string corresponding to the name of the action to trigger,
// e.g. "take-item", and a Target string corresponding to the target of the action, if any, e.g. "sword".
type Trigger struct {
	Target string
	Action string
}

// String returns the string representation of a trigger in the format TRIGGER(TARGET).
func (t *Trigger) String() string {
	return fmt.Sprintf("%s(%s)", t.Action, t.Target)
}

// MatchesName returns true if the provided string matches any value in []Names.
func (t *Thing) MatchesName(inputName string) bool {
	for _, name := range t.Names {
		if strings.ToLower(inputName) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// Show sets the Visible attribute of the Thing to true, revealing it to the player.
func (t *Thing) Show() {
	t.Visible = true
}

// Hide sets the Visible attribute of the Thing to false, hiding it from the player.
func (t *Thing) Hide() {
	t.Visible = false
}
