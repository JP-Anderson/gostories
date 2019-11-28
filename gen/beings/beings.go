package beings

import (
	"encoding/xml"
	"strings"

	"gostories/engine/io"
	"gostories/engine/speech"
	generator "gostories/gen/speech"
	"gostories/things"
	gstring "gostories/utils/strings"
	gxml "gostories/utils/xml"
)

// Get returns a *Being matching the provided name.
func Get(name string) *things.Being {
	return beings[gstring.ToIDString(name)]
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
		io.Handler.NewLinef("beingsFromXML failed: %v", err)
	}
	m := make(map[string]*things.Being, len(t.Being))
	for _, b := range t.Being {
		being := newBeing(b)
		m[gstring.ToIDString(being.Name)] = being
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
	Names     string
	LookText  string
	Species   string
	IsVisible string
	Speech    string
	AltSpeech string
}

func newBeing(in Being) *things.Being {
	mainSpeech := generator.Tree(in.Speech)
	var altSpeech speech.Tree
	if in.AltSpeech != "" {
		altSpeech = generator.Tree(in.AltSpeech)
	}
	names := strings.Split(in.Names, "|")
	being := &things.Being{
		Thing: things.Thing{
			Name:     in.Name,
			Names:    names,
			LookText: in.LookText,
		},
		Species:   in.Species,
		Speech:    mainSpeech,
		AltSpeech: &altSpeech,
	}
	if in.IsVisible == "y" {
		being.Show()
	}
	return being
}
