package xml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	itemsXMLPath    = "/gen/items/data/items.xml"
	featuresXMLPath = "/gen/features/data/features.xml"
	areasXMLPath    = "/gen/areas/data/areas.xml"
	speechXMLPath   = "/gen/speech/data/"
	beingsXMLPath   = "/gen/beings/data/beings.xml"

	rootFolderName = "gostories"
)

// BytesForItems returns the bytes of the items XML data file.
func BytesForItems() []byte {
	return bytes(itemsXMLPath)
}

// BytesForFeatures returns the bytes of the features XML data file.
func BytesForFeatures() []byte {
	return bytes(featuresXMLPath)
}

// BytesForAreas returns the bytes of the areas XML data file.
func BytesForAreas() []byte {
	return bytes(areasXMLPath)
}

// BytesForSpeechTree returns the bytes of an XML represented speech tree.
func BytesForSpeechTree(filename string) []byte {
	return bytes(speechXMLPath + filename + ".xml")
}

// BytesForBeings returns the bytes of the beings XML data file.
func BytesForBeings() []byte {
	return bytes(beingsXMLPath)
}

func bytes(testDirPath string) []byte {
	pathToRead := getRootPath() + testDirPath
	absPath, err := filepath.Abs(pathToRead)
	if err != nil {
		panic(fmt.Sprintf("Error finding absolute path for file [%v]: %v", pathToRead, err))
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading file [%s]: %s", pathToRead, err))
	}
	return bytes
}

func getRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Check if already in root directory (e.g. main.go)
	if getCurrentDirectoryNameFromPath(dir) == rootFolderName {
		return dir
	}

	// Get parent directory
	dir = filepath.Dir(dir)
	// If not at root, continue getting parent
	for getCurrentDirectoryNameFromPath(dir) != rootFolderName {
		dir = filepath.Dir(dir)
	}
	return dir
}

func getCurrentDirectoryNameFromPath(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}
