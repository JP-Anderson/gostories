package speechgeneration

import (
	"encoding/xml"
	"io/ioutil"
	path "path/filepath"

	"gostories/engine/io"
	"gostories/speech"
)


// SpeechFromXMLFile takes a filepath string and attempts to load an XML file from the path, and parse it into a
// speech.Tree object.
// TODO: error handling
func SpeechFromXMLFile(filepath string) speech.Tree {
	absPath, err := path.Abs(filepath)
	if err != nil {
		io.ActiveInputOutputHandler.NewLinef("Error finding absolute path for file [%v]: %v", filepath, err)
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		//
	}
	return speechFromXML(bytes)
}

func speechFromXML(xmlBytes []byte) speech.Tree {
	t := &speech.Tree{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		io.ActiveInputOutputHandler.NewLinef("speechFromXML failed: %v", err)
	}
	return *t
}
