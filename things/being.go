package things

import (
	generator "gostories/gen/speech"
	"gostories/speech"
)

const speechDataRoot = "./gen/speech/speech_data/"

// Being is a Non-Playable Character (NPC) which the player can interact with.
type Being struct {
	Thing
	Species   string
	Speech    speech.Tree
	AltSpeech *speech.Tree
}

// NewBubbles creates a new Bubbles.
// Will either construct these from XML in future or find some way to autogenerate.
func NewBubbles() *Being {
	translatedSpeech := generator.SpeechFromXMLFile(
		speechDataRoot + "bubbles_human.xml",
	)
	catSpeech := generator.SpeechFromXMLFile(
		speechDataRoot + "bubbles.xml",
	)
	being := &Being{
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

// GetName returns the name of the Being.
func (b Being) GetName() string {
	return b.Name
}

// GetLookText returns the description when the player looks at the Being.
func (b Being) GetLookText() string {
	return b.LookText
}

// GetThing returns the underlying Thing struct.
func (b Being) GetThing() Thing { return b.Thing }
