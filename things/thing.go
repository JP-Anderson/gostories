package things

// A Thing is anything in the game that can be interacted with in some way. All Things have a concept of visibility,
// one Thing might be hidden until one or more pre-requisites have been met (e.g. player has an item, or a skill)
type Thing struct {
	visible bool
	Triggers Triggers
}

type Perceivable interface {
	Show()
	Hide()
	Visible() bool
}

func (t *Thing) Show() {
	t.visible = true
}

func (t *Thing) Hide() {
	t.visible = false
}

func (t *Thing) Visible() bool {
	return t.visible
}
