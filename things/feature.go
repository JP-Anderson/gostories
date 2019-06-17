package things

type Feature interface {

	// Get the name of the Feature
	GetName() string

	// Description given when Looking at the Feature
	GetLookText() string
}


func NewShelfFeature() ShelfFeature {
	return ShelfFeature{
		Thing{Triggers: map[string]Trigger{
			"" : RevealItemTrigger{
				CatCollarItem{}.Thing,
			},
		}},
	}
}

type ShelfFeature struct {
	Thing
}

func (c ShelfFeature) GetName() string { return "shelf" }

func (c ShelfFeature) GetLookText() string {
	return "The shelf seems to contain a few old magazines and a cat collar"
}

