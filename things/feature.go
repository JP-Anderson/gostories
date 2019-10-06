package things

type Feature interface {

	// Get the Name of the Feature
	GetName() string

	// Description given when Looking at the Feature
	GetLookText() string

	// Get the Thing
	GetThing() *Thing
}

