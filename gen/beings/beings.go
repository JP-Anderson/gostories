package beings

import (
	"encoding/xml"
	"strings"

	"gostories/engine/io"
	"gostories/engine/speech"
	generator "gostories/gen/speech"
	"gostories/things"
	gxml "gostories/xml"
)

// Get returns a *Being matching the provided name.
func Get(name string) *things.Being {
	return beings[name]
}

var beings = getBeings()

func getBeings() map[string]*things.Being {
	return loadFromXML()
}

func loadFromXML() map[string]*things.Being {
	return beingsFromXML(gxml.BytesForBeings())
}

func beingsFromXML(xmlBytes []byte) map[string]*things.Being {
	t := &Beings{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.ActiveInputOutputHandler.NewLinef("beingsFromXML failed: %v", err)
	}
	m := make(map[string]*things.Being, len(t.Being))
	for _, b := range t.Being {
		being := newBeing(b)
		m[strings.ToLower(being.Name)] = being
	}
	return m
}

// Beings specifies the xml schema for a list of Beings.
type Beings struct {
	Being []Being
}

// Being specifies the xml schema for a Being in-game.
type Being struct {
	Name      string
	LookText  string
	Species   string
	IsVisible string
	Speech    string
	AltSpeech string
}

func newBeing(in Being) *things.Being {
	mainSpeech := generator.Tree(in.Speech)
	var altSpeech *speech.Tree
	if in.AltSpeech != "" {
		altSpeech = generator.TreePtr(generator.Tree(in.AltSpeech))
	}
	being := &things.Being{
		Thing: things.Thing{
			Name:     in.Name,
			LookText: in.LookText,
		},
		Species:   in.Species,
		Speech:    mainSpeech,
		AltSpeech: altSpeech,
	}
	if in.IsVisible == "y" {
		being.Show()
	}
	return being
}
