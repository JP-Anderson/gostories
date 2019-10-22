package things

import (
	"gostories/engine/speech"
)

const speechDataRoot = "./gen/speech/speech_data/"

// Being is a Non-Playable Character (NPC) which the player can interact with.
type Being struct {
	Thing
	Species   string
	Speech    speech.Tree
	AltSpeech *speech.Tree
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
