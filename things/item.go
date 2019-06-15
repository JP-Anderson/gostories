package things

type Item interface {

	// Get the name of the item
	GetName() string

	// Description given when Looking at the item
	GetLookText() string
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

func (c CatCollarItem) Toggle() {}

type ShrubberyItem struct {
	Thing
}

func (c ShrubberyItem) GetName() string { return "shrubbery" }

func (c ShrubberyItem) GetLookText() string {
	return "A small but rather well cared for shrubbery."
}

func (c ShrubberyItem) Toggle() {}
