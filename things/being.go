package things

import (
	"gostories/generator"
	"gostories/speech"
)

type Being struct {
	Thing
	Species   string
	Speech    speech.Tree
	AltSpeech *speech.Tree
}

// Simple examples constructed from code for now
// Will either construct these from XML in future or find some way to autogenerate
func NewBubbles() Being {
	translatedSpeech := generator.SpeechFromXMLFile("./generator/speech_data/bubbles_human.xml")
	catSpeech := generator.SpeechFromXMLFile("./generator/speech_data/bubbles.xml")
	being := Being{
		Thing: Thing{
			Name:     "Bubbles",
			LookText: "The cat is reasonably small, ginger, and chunky.",
		},
		Species:   "Cat",
		Speech:    translatedSpeech,
		AltSpeech: &catSpeech,
	}
	being.Show()
	return being
}

func (b Being) GetName() string {
	return b.Name
}

func (b Being) GetLookText() string {
	return b.LookText
}

func (b Being) GetThing() Thing { return b.Thing }
