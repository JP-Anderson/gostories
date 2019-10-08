package things

// A Thing is anything in the game that can be interacted with in some way. All Things have
// a Name (identifying string), LookText (text which is output when the player executes the
// Look action on the Thing, a Visible attribute (is the Thing currently revealed to the 
// player, and a set of Triggers. Triggers have effect on the game state, based on
// interactions the player makes with the object.
type Thing struct {
	Name     string
	LookText string
	Visible  bool
	// Triggers maps verb command strings to trigger functions. See trigger.go for trigger functions.
	Triggers map[string]string
}

// Show sets the Visible attribute of the Thing to true, revealing it to the player.
func (t *Thing) Show() {
	t.Visible = true
}

// Hide sets the Visible attribute of the Thing to false, hiding it from the player.
func (t *Thing) Hide() {
	t.Visible = false
}
