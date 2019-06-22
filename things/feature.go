package things

type Feature interface {

	// Get the Name of the Feature
	GetName() string

	// Description given when Looking at the Feature
	GetLookText() string

	// Get the Thing
	GetThing() Thing
}

type ShelfFeature struct {
	Thing
}

func (s ShelfFeature) GetName() string { return s.Name }

func (s ShelfFeature) GetLookText() string { return s.LookText }

func (s ShelfFeature) GetThing() Thing { return s.Thing }
