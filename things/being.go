package things

import (
	"gostories/speech"
)

type Being struct {
	Thing
	Name      string
	Species   string
	Speech    speech.Tree
	AltSpeech *speech.Tree
}
