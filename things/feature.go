package things

// A Feature is an item or part of an Area which is not an Item (i.e. the player
// cannot pick it up). However it is not a Being, meaning the player cannot interact
// conversationally with it. Features can be interacted with to reveal Items, Features,
// Beings, etc.
type Feature interface {

	// Get the Name of the Feature
	GetName() string

	// Description given when Looking at the Feature
	GetLookText() string

	// Get the Thing
	GetThing() *Thing
}

