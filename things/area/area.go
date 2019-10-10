package area

import (
	"strings"

	"gostories/engine/inventory"
	"gostories/things"
)

// Area represents an in-game location. It can be connected to other Areas by exits, and it contains
// game objects (Items, Beings, Features, etc.) which the player can interact with if they are within
// the Area.
type Area struct {
	Look     string
	Exits    map[Direction]Exit
	Items    *inventory.ItemStore
	Beings   []things.Being
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

// CheckAreaBeingsForThing takes a target Being name, it iterates through the Being objects stored in
// the Area, and returns a Thing pointer to the Being if it exists. Note, this method does not take
// into account if the Being is visible to the player.
func (a Area) CheckAreaBeingsForThing(targetName string) *things.Thing {
	for _, b := range a.Beings {
		if strings.ToLower(b.GetName()) == strings.ToLower(targetName) {
			t := b.GetThing()
			return &t
		}
	}
	return nil
}

// CheckAreaFeaturesForThing takes a target Feature name, it iterates through the Feature objects
// stored in the Area, and returns a Thing pointer to the Feature if it exists. Note, this method
// does not take into account if the Feature is visible to the player.
func (a Area) CheckAreaFeaturesForThing(targetName string) *things.Thing {
	for _, f := range a.Features {
		if strings.ToLower(f.GetName()) == strings.ToLower(targetName) {
			t := f.GetThing()
			return t
		}
	}
	return nil
}
