package things

import "strings"

type Area struct {
	Look     string
	Exits    map[Direction]Exit
	Items    []Item
	Beings   []Being
	Features []Feature
}

type Exit struct {
	To   *Area
	From *Area
}

type Direction string

const (
	North = "north"
	East  = "east"
	South = "south"
	West  = "west"
)

func (a Area) CheckAreaItemsForThing(targetName string) *Thing {
	for _, i := range a.Items {
		if strings.ToLower(i.GetName()) == strings.ToLower(targetName) {
			t := i.GetThing()
			return &t
		}
	}
	return nil
}

func (a Area) CheckAreaBeingsForThing(targetName string) *Thing {
	for _, b := range a.Beings {
		if strings.ToLower(b.GetName()) == strings.ToLower(targetName) {
			t := b.GetThing()
			return &t
		}
	}
	return nil
}

func (a Area) CheckAreaFeaturesForThing(targetName string) *Thing {
	for _, f := range a.Features {
		if strings.ToLower(f.GetName()) == strings.ToLower(targetName) {
			t := f.GetThing()
			return &t
		}
	}
	return nil
}
