package things

// An Item is something the player can pick up and place in their store.
type Item interface {

	// Get the Name of the item.
	GetName() string

	// Description given when Looking at the item.
	GetLookText() string

	// Get the Thing returns the base Thing struct attached to the Item
	GetThing() *Thing
}

// An Equippable is a sub-class of Item which can be equipped on the player for
// additional benefits.
type Equippable interface {

	// Toggle equips an Equippable if it is unequipped, and unequips it if it is equipped.
	Toggle()
}
