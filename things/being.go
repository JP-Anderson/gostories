package things

import "gostories/speech"

type Being struct {
	Name    string
	Species string
	Speech speech.Tree
}

func (b *Being) SpeakTo() {
	speech.Run(b.Speech)
}