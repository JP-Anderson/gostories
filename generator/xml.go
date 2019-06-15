package generator

import (
	"encoding/xml"

	"gostories/engine"
	"gostories/speech"
)

func SpeechFromXml(xmlBytes []byte) speech.Tree {
	t := &speech.Tree{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		engine.NewLinef("SpeechFromXML failed: %v", err)
	}
	return *t
}
