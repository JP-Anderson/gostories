
// Do not edit! Autogenerated file //

package items
// items package

import "gostories/things"

var Item_Collar = NewCollarItem()
var collar_Item *CollarItem

type CollarItem struct {
	things.Thing
}

func (c CollarItem) GetName() string { return c.Name }

func (c CollarItem) GetLookText() string { return c.LookText }

func (c CollarItem) Toggle() {}

func (c *CollarItem) Show() { c.Thing.Visible = true }

func (c *CollarItem) Hide() { c.Thing.Visible = false }

func (c CollarItem) GetThing() things.Thing { return c.Thing }


func (c CollarItem) Take() {}

func NewCollarItem() *CollarItem {
	if collar_Item == nil {
		collar_Item = &CollarItem{}
		collar_Item.Name = "collar"
		collar_Item.LookText = "A small red cat collar with a bell."
	}

	return collar_Item
}

var Item_Shrubbery = NewShrubberyItem()
var shrubbery_Item *ShrubberyItem

type ShrubberyItem struct {
	things.Thing
}

func (c ShrubberyItem) GetName() string { return c.Name }

func (c ShrubberyItem) GetLookText() string { return c.LookText }

func (c ShrubberyItem) Toggle() {}

func (c *ShrubberyItem) Show() { c.Thing.Visible = true }

func (c *ShrubberyItem) Hide() { c.Thing.Visible = false }

func (c ShrubberyItem) GetThing() things.Thing { return c.Thing }


func (c ShrubberyItem) Take() {}

func NewShrubberyItem() *ShrubberyItem {
	if shrubbery_Item == nil {
		shrubbery_Item = &ShrubberyItem{}
		shrubbery_Item.Name = "shrubbery"
		shrubbery_Item.LookText = "A small but rather well cared for shrubbery."
	}

		shrubbery_Item.Show()

	return shrubbery_Item
}

var Item_Sardines = NewSardinesItem()
var sardines_Item *SardinesItem

type SardinesItem struct {
	things.Thing
}

func (c SardinesItem) GetName() string { return c.Name }

func (c SardinesItem) GetLookText() string { return c.LookText }

func (c SardinesItem) Toggle() {}

func (c *SardinesItem) Show() { c.Thing.Visible = true }

func (c *SardinesItem) Hide() { c.Thing.Visible = false }

func (c SardinesItem) GetThing() things.Thing { return c.Thing }


func (c SardinesItem) Take() {}

func NewSardinesItem() *SardinesItem {
	if sardines_Item == nil {
		sardines_Item = &SardinesItem{}
		sardines_Item.Name = "sardines"
		sardines_Item.LookText = "A tin of tasty sardines preserved in olive oil."
	}

		sardines_Item.Show()

	return sardines_Item
}
