package items

import "gostories/things"

var Item_Collar = NewCollarItem()

type CollarItem struct {
	things.Thing
}

func NewCollarItem() CollarItem {
	c := CollarItem{}
	c.Name = "collar"
	c.LookText = "A small red cat collar with a bell."
	return c
}

func (c CollarItem) GetName() string { return c.Name }

func (c CollarItem) GetLookText() string { return c.LookText }

func (c CollarItem) Take() {}

func (c CollarItem) GetThing() things.Thing { return c.Thing }

func (c CollarItem) Toggle() {}

func (c CollarItem) Show() {
	c.Visible = true
}

func (c CollarItem) Hide() {
	c.Visible = false
}

var Item_Shrubbery = NewShrubberyItem()

type ShrubberyItem struct {
	things.Thing
}

func NewShrubberyItem() ShrubberyItem {
	c := ShrubberyItem{}
	c.Name = "shrubbery"
	c.LookText = "A small but rather well cared for shrubbery."
	return c
}

func (c ShrubberyItem) GetName() string { return c.Name }

func (c ShrubberyItem) GetLookText() string { return c.LookText }

func (c ShrubberyItem) Take() {}

func (c ShrubberyItem) GetThing() things.Thing { return c.Thing }

func (c ShrubberyItem) Toggle() {}

func (c ShrubberyItem) Show() {
	c.Visible = true
}

func (c ShrubberyItem) Hide() {
	c.Visible = false
}
