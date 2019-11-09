package speech

import (
	"encoding/xml"

	"gostories/engine/io"
	"gostories/engine/speech"
	gxml "gostories/xml"
)

// Tree loads a conversation speech Tree from bytes.
func Tree(name string) speech.Tree {
	return speechFromXML(gxml.BytesForSpeechTree(name))
}

func speechFromXML(xmlBytes []byte) speech.Tree {
	t := &speech.Tree{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		io.Handler.NewLinef("speechFromXML failed: %v", err)
	}
	return *t
}
