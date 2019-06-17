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

func NewCatCollarItem() CatCollarItem {
	c := CatCollarItem{}
	c.name = "collar"
	c.lookText = "A small red cat collar with a bell."
	return c
}

func (c CatCollarItem) GetName() string { return c.Name() }

func (c CatCollarItem) GetLookText() string { return c.GetLookText() }

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

func NewShrubberyItem() ShrubberyItem {
	s := ShrubberyItem{}
	s.name = "shrubbery"
	s.lookText = "A small but rather well cared for shrubbery."
	return s
}

func (s ShrubberyItem) GetName() string { return s.Name() }

func (s ShrubberyItem) GetLookText() string { return s.LookText() }

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
