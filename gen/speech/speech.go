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

// TreePtr is a helper method which returns a pointer to a Tree it is given.
func TreePtr(tree speech.Tree) *speech.Tree { return &tree }

func speechFromXML(xmlBytes []byte) speech.Tree {
	t := &speech.Tree{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		io.ActiveInputOutputHandler.NewLinef("speechFromXML failed: %v", err)
	}
	return *t
}
