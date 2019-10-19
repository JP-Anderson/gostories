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
	itemsXMLPath   = "/gen/items/data/items.xml"
	rootFolderName = "gostories"
)

func BytesForItems() []byte {
	return bytes(itemsXMLPath)
}

func bytes(path string) []byte {
	pathToRead := getRootPath() + itemsXMLPath
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
