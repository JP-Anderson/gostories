package features

import (
	"encoding/xml"

	"gostories/engine/io"
	"gostories/things"
	gxml "gostories/xml"
)

// Feature returns any feature which has a name matching the provided name.
func Feature(name string) things.Feature {
	return features[name]
}

var features = getFeatures()

func getFeatures() map[string]things.Feature {
	return loadFromXML()
}

func loadFromXML() (items map[string]things.Feature) {
	return featuresFromXML(gxml.BytesForFeatures())
}

func featuresFromXML(xmlBytes []byte) map[string]things.Feature {
	t := &XFeatures{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.ActiveInputOutputHandler.NewLinef("featuresFromXML failed: %v", err)
	}
	m := make(map[string]things.Feature, len(t.XFeature))
	for _, f := range t.XFeature {
		triggers := make(map[string]string)
		for _, trigger := range f.XTriggerStrings.XTriggerString {
			triggers[trigger.Action] = trigger.Trigger
		}
		newFeature := &feature{
			&things.Thing{
				Name:     f.Name,
				LookText: f.LookText,
				Triggers: triggers,
			},
		}
		m[f.Name] = newFeature
	}
	return m
}

// XFeatures specifies the xml schema for a list of features.
type XFeatures struct {
	XFeature []XFeature
}

// XFeature specifies the xml schema for a feature (area/item/feature of interest in-game).
type XFeature struct {
	Name            string
	LookText        string
	XTriggerStrings XTriggerStrings
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
