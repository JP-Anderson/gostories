package speechgeneration

import (
	"encoding/xml"
	"io/ioutil"
	path "path/filepath"

	"gostories/engine/io"
	"gostories/speech"
)

func SpeechFromXMLFile(filepath string) speech.Tree {
	absPath, err := path.Abs(filepath)
	if err != nil {
		io.NewLinef("Error finding absolute path for file [%v]: %v", filepath, err)
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		//
	}
	return SpeechFromXml(bytes)
}

func SpeechFromXml(xmlBytes []byte) speech.Tree {
	t := &speech.Tree{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		io.NewLinef("SpeechFromXML failed: %v", err)
	}
	return *t
}
