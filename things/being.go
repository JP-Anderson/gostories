package things

import (
	"gostories/speech"
)

type Being struct {
	Name    string
	Species string
	Speech speech.Tree
	AltSpeech *speech.Tree
}
