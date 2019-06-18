package things

type Item interface {

	// Get the Name of the item
	GetName() string

	// Description given when Looking at the item
	GetLookText() string

	// Take the item into the inventory
	Take()

	// Get the Thing
	GetThing() Thing
}

type Equippable interface {

	// Toggle equips an Equippable if it is unequipped, and unequips it if it is equipped.
	Toggle()
}

// TODO create items dir. Eventually create from xml/json
type CatCollarItem struct {
	Thing
}

func NewCatCollarItem() CatCollarItem {
	c := CatCollarItem{}
	c.Name = "collar"
	c.LookText = "A small red cat collar with a bell."
	return c
}

func (c CatCollarItem) GetName() string { return c.Name }

func (c CatCollarItem) GetLookText() string { return c.GetLookText() }

func (c CatCollarItem) Take() {}

func (c CatCollarItem) GetThing() Thing { return c.Thing }

func (c CatCollarItem) Toggle() {}

func (c CatCollarItem) Show() {
	c.Visible = true
}

func (c CatCollarItem) Hide() {
	c.Visible = false
}

type ShrubberyItem struct {
	Thing
}

func NewShrubberyItem() ShrubberyItem {
	s := ShrubberyItem{}
	s.Name = "shrubbery"
	s.LookText = "A small but rather well cared for shrubbery."
	return s
}

func (s ShrubberyItem) GetName() string { return s.Name }

func (s ShrubberyItem) GetLookText() string { return s.LookText }

func (s ShrubberyItem) Take() {}

func (s ShrubberyItem) GetThing() Thing { return s.Thing }

func (s ShrubberyItem) Show() {
	s.Visible = true
}

func (s ShrubberyItem) Hide() {
	s.Visible = false
}
