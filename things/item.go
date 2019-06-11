package things

type Item interface {

	// Get the name of the item
	GetName() string

	// Description given when Looking at the item
	GetLookText() string

}

// TODO create items dir. Eventually create from xml/json
type CatCollarItem struct {
	Thing
}

func (c CatCollarItem) GetName() string { return "cat_collar" }

func  (c CatCollarItem) GetLookText() string {
	return "A small red cat collar with a bell."
}