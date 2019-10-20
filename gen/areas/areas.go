package areas

import (
	"encoding/xml"

	"gostories/engine/io"
	"gostories/things"
	"gostories/things/area"
	gxml "gostories/xml"
)

// Area returns any area which has a name matching the provided name.
func Area(name string) area.Area {
	return areas[name]
}

var areas = getAreas()

func getAreas() map[string]area.Area {
	return loadFromXML()
}

func loadFromXML() (items map[string]area.Area) {
	return areasFromXML(gxml.BytesForAreas())
}

func areasFromXML(xmlBytes []byte) map[string]area.Area {
	t := &Areas{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.ActiveInputOutputHandler.NewLinef("areasFromXML failed: %v", err)
	}
	m := make(map[string]area.Area, len(t.XArea))
	for _, a := range t.XArea {
		newArea := area.Area{
			Look: a.LookText,
		}
		if len(a.Beings) > 0 {

		}
		m[a.Name] = newArea
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
	// TODO Exits    map[string]
	Features []Feature
	Beings   []Being
}

type Feature struct {
	Name string
}

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
