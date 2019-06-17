package things

type Item interface {

	// Get the name of the item
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

func (c CatCollarItem) GetName() string { return "collar" }

func (c CatCollarItem) GetLookText() string {
	return "A small red cat collar with a bell."
}

func (c CatCollarItem) Take() {}

func (c CatCollarItem) GetThing() Thing { return c.Thing }

func (c CatCollarItem) Toggle() {}

func (c *CatCollarItem) Show() {
	c.visible = true
}

func (c *CatCollarItem) Hide() {
	c.visible = false
}

func (c *CatCollarItem) Visible() bool {
	return c.visible
}


type ShrubberyItem struct {
	Thing
}

func (s ShrubberyItem) GetName() string { return "shrubbery" }

func (s ShrubberyItem) GetLookText() string {
	return "A small but rather well cared for shrubbery."
}

func (s ShrubberyItem) Take() {}

func (s ShrubberyItem) GetThing() Thing { return s.Thing }

func (s ShrubberyItem) Show() {
	s.visible = true
}

func (s ShrubberyItem) Hide() {
	s.visible = false
}

func (s ShrubberyItem) Visible() bool {
	return s.visible
}
