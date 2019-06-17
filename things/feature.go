package things

type Feature interface {

	// Get the name of the Feature
	GetName() string

	// Description given when Looking at the Feature
	GetLookText() string

	// Get the Thing
	GetThing() Thing
}


func NewShelfFeature() ShelfFeature {
	return ShelfFeature{
		Thing{Triggers: map[string]Trigger{
			"look" : RevealItemTrigger{
				CatCollarItem{}.Thing,
			},
		}},
	}
}

type ShelfFeature struct {
	Thing
}

func (s ShelfFeature) GetName() string { return "shelf" }

func (s ShelfFeature) GetLookText() string {
	return "The shelf seems to contain a few old magazines and a cat collar"
}

func (s ShelfFeature) GetThing() Thing { return s.Thing }
