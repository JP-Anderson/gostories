package areas

import (
	"encoding/xml"

	"gostories/engine/io"
	"gostories/engine/store"
	"gostories/gen/features"
	"gostories/gen/items"
	"gostories/things"
	"gostories/things/area"
	gxml "gostories/xml"
)

// Area returns any area which has a name matching the provided name.
func Area(name string) area.Area {
	return *areas[name]
}

var areas = getAreas()

func getAreas() map[string]*area.Area {
	return loadFromXML()
}

func loadFromXML() (items map[string]*area.Area) {
	return areasFromXML(gxml.BytesForAreas())
}

func areasFromXML(xmlBytes []byte) map[string]*area.Area {
	t := &Areas{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.ActiveInputOutputHandler.NewLinef("areasFromXML failed: %v", err)
	}
	m := make(map[string]*area.Area, len(t.XArea))
	for _, a := range t.XArea {
		newArea := &area.Area{
			Look: a.LookText,
		}
		if len(a.Beings) > 0 {
			//hardcoding the only being for now. TODO generate with XML
			newArea.Beings = []*things.Being{
				things.NewBubbles(),
			}
		}
		newArea.Features = []things.Feature{}
		for _, f := range a.Features.Feature {
			newArea.Features = append(newArea.Features, features.Feature(f.Name))
		}
		newArea.Items = store.NewItemStore()
		for _, i := range a.Items.Item {
			newArea.Items.StoreItem(items.Item(i.Name))
		}
		m[a.Name] = newArea
	}
	// Loop a second time to add the exits between Areas (we need to do this after the first
	// loop to ensure all the Areas are created when we start linking them.)
	for _, a := range t.XArea {
		current := m[a.Name]
		current.Exits = make(map[area.Direction]area.Exit)
		if a.Exits.North != "" {
			current.Exits[area.North] = area.Exit{
				To:   m[a.Exits.North],
				From: m[a.Name],
			}
		}
		if a.Exits.East != "" {
			current.Exits[area.East] = area.Exit{
				To:   m[a.Exits.East],
				From: m[a.Name],
			}
		}
		if a.Exits.South != "" {
			current.Exits[area.South] = area.Exit{
				To:   m[a.Exits.South],
				From: m[a.Name],
			}
		}
		if a.Exits.West != "" {
			current.Exits[area.West] = area.Exit{
				To:   m[a.Exits.West],
				From: m[a.Name],
			}
		}
	}
	return m
}

// Areas specifies the xml schema for a list of areas.
type Areas struct {
	XArea []XArea
}

// XArea specifies the xml schema for an area in-game.
type XArea struct {
	Name     string
	LookText string
	Exits    Exits
	Features Features
	Beings   []Being
	Items    Items
}

// Exits specifies the xml schema for the exits of an Area.
type Exits struct {
	North string
	East  string
	South string
	West  string
}

// Features specifies the xml schema for a list of Feature references.
type Features struct {
	Feature []Feature
}

// Feature specifies the xml schema for a Feature reference in an Area. Note a Feature reference is not
// a full feature object, but contains a string ID of a feature to be linked from the loaded features.
type Feature struct {
	Name string
}

// Items specifies the xml schema for a list of Item references.
type Items struct {
	Item []Item
}

// Item specifies the xml schema for an Item reference in an Area. Note a Item reference is not
// a full item object, but contains a string ID of an item to be linked from the loaded items.
type Item struct {
	Name string
}

// Being specifies the xml schema for a Being reference in an Area. Note a Being reference is not
// a full Being object, but contains a string ID of a Being to be linked from the loaded Beings.
type Being struct {
	Name string
}

// XTriggerStrings specifies the xml schema for a list of action trigger strings that are present on a feature.
type XTriggerStrings struct {
	XTriggerString []XTriggerString
}

// XTriggerString specifies the xml schema for an action trigger on a feature, which maps verb Action to a
// trigger string (Trigger), to be executed.
type XTriggerString struct {
	Action  string
	Trigger string
}

type feature struct {
	t *things.Thing
}

func (i *feature) GetName() string {
	return i.t.Name
}

func (i *feature) GetLookText() string {
	return i.t.LookText
}

func (i *feature) GetThing() *things.Thing {
	return i.t
}
