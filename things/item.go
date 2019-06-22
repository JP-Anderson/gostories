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
