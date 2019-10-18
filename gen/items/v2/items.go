package v2

import (
	"gostories/things"
)

// Items contains all the items a player can pick up. It is currently indexed by the item name, however ideally
// it should be indexed by an enum/custom int and this should be the only place to access items.
var Items = buildItemsMap()

type itemToAdd struct {
	Name         string
	LookText     string
	IsVisible    bool
	IsEquippable bool
}

type item struct {
	t *things.Thing
	// Name         string
	// LookText     string
	// IsVisible    bool
	// IsEquippable bool
}

func (i *item) GetName() string {
	return i.t.Name
}

func (i *item) GetLookText() string {
	return i.t.LookText
}

func (i *item) GetThing() *things.Thing {
	return i.t
}

type equippable struct {
	t *things.Thing
	// Name         string
	// LookText     string
	// IsVisible    bool
	// IsEquippable bool
}

func (i *equippable) GetName() string {
	return i.t.Name
}

func (i *equippable) GetLookText() string {
	return i.t.LookText
}

func (i *equippable) GetThing() *things.Thing {
	return i.t
}

func (i *equippable) Toggle() {}

// WriteItemsFile outputs a Go source file containing the in-game Items to generate
func buildItemsMap() (items map[string]things.Item) {
	itemsToAdd := []itemToAdd{
		itemToAdd{"collar", "A small red cat collar with a bell.", false, true},
		itemToAdd{"shrubbery", "A small but rather well cared for shrubbery.", true, false},
		itemToAdd{"sardines", "A tin of tasty sardines preserved in olive oil.", true, false},
	}
	items = make(map[string]things.Item, len(itemsToAdd))
	for _, itemToAdd := range itemsToAdd {
		if itemToAdd.IsEquippable {
			items[itemToAdd.Name] = &equippable{
				&things.Thing{
					Name:     itemToAdd.Name,
					LookText: itemToAdd.LookText,
					Visible:  itemToAdd.IsVisible,
				},
			}
		} else {
			items[itemToAdd.Name] = &item{
				&things.Thing{
					Name:     itemToAdd.Name,
					LookText: itemToAdd.LookText,
					Visible:  itemToAdd.IsVisible,
				},
			}
		}

	}
	return
}
