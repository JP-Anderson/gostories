package things

// A Thing is anything in the game that can be interacted with in some way. All Things have a concept of visibility,
// one Thing might be hidden until one or more pre-requisites have been met (e.g. player has an item, or a skill)
type Thing struct {
	Name     string
	LookText string
	Visible  bool
	// Triggers maps verb command strings to trigger functions. See trigger.go for trigger functions.
	Triggers map[string]string
}

func (t *Thing) Show() {
	t.Visible = true
}

func (t *Thing) Hide() {
	t.Visible = false
}
