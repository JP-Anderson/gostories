package area

import (
	"strings"

	"gostories/engine/store"
	"gostories/things"
)

// Area represents an in-game location. It can be connected to other Areas by exits, and it contains
// game objects (Items, Beings, Features, etc.) which the player can interact with if they are within
// the Area.
type Area struct {
	Look     string
	Exits    map[Direction]Exit
	Items    *store.ItemStore
	Beings   []*things.Being
	Features []things.Feature
}

// Exit represents a path the player can navigate through to leave an Area. Each Exit has a pointer
// to the Area it leads to (To), and the Area it exits from (From). Currently Exits in an area can
// be mapped to Directions, but ideally in the future Exits will be assignable to arbitrary objects.
type Exit struct {
	To   *Area
	From *Area
}

// Direction is a custom string type, and is used to map known direction names to Exits in an Area.
type Direction string

// This const block contains the current supported Direction string names.
const (
	North = "north"
	East  = "east"
	South = "south"
	West  = "west"
)

// StringToDirection maps a string to the area.Direction value.
var StringToDirection = map[string]Direction{
	North: North,
	East:  East,
	South: South,
	West:  West,
}

// OppositeDirection maps a string to it's opposite area.Direction value.
var OppositeDirection = map[string]Direction{
	North: South,
	East:  West,
	South: North,
	West:  East,
}

// AddFeature adds a feature to the Area.
func (a *Area) AddFeature(feature things.Feature) {
	if a.Features == nil {
		a.Features = []things.Feature{}
	}
	a.Features = append(a.Features, feature)
}

// Checker is a func which given an *Area and string, returns any things.Thing that has a name matching the
// provided string.
type Checker func(*Area, string) *things.Thing

// CheckAreaForThing takes a thing name string, and a variable number of Checker functions. It will return a
// pointer to any Thing (matching the provided name) that has been yielded by the checkers.
func (a *Area) CheckAreaForThing(thingName string, checkers ...Checker) *things.Thing {
	for _, checker := range checkers {
		res := checker(a, thingName)
		if res != nil {
			return res
		}
	}
	return nil
}

// FindItemByName takes a target Item name, it iterates through the Item objects stored in the Area, and
// returns any Item with a matching name (not case-sensitive). Note, this method does not take into
// account if the Item is visible to the player.
func (a Area) FindItemByName(targetName string) things.Item {
	item, err := a.Items.GetItemWithName(targetName)
	if item != nil && err == nil {
		return *item
	}
	return nil
}

// CheckItems is a Checker func which checks the Items in an Area matching the targetName.
func CheckItems(a *Area, targetName string) *things.Thing {
	item, err := a.Items.GetItemWithName(targetName)
	if item != nil && err == nil {
		t := *item
		return t.GetThing()
	}
	return nil
}

// CheckBeings is a Checker func which checks the Beings in an Area matching the targetName.
func CheckBeings(a *Area, targetName string) *things.Thing {
	for _, b := range a.Beings {
		if strings.ToLower(b.GetName()) == strings.ToLower(targetName) {
			t := b.GetThing()
			return &t
		}
	}
	return nil
}

// CheckFeatures is a Checker func which checks the Features in an Area matching the targetName.
func CheckFeatures(a *Area, targetName string) *things.Thing {
	for _, f := range a.Features {
		if strings.ToLower(f.GetName()) == strings.ToLower(targetName) {
			t := f.GetThing()
			return t
		}
	}
	return nil
}
