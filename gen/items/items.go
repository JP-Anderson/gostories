package items

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gostories/engine/io"
	"gostories/things"
)

const rootFolderName = "gostories"

// Item returns any item which has a name matching the provided name
func Item(name string) things.Item {
	return items[name]
}

// Items contains all the items a player can pick up. It is currently indexed by the item name, however ideally
// it should be indexed by an enum/custom int and this should be the only place to access items.
var items = getItems()

func getItems() (i map[string]things.Item) {
	return loadFromXML()
}

func loadFromXML() (items map[string]things.Item) {
	itemsXMLPath := getRootPath() + "/gen/items/data/items.xml"
	print(itemsFromXML)
	absPath, err := filepath.Abs(itemsXMLPath)
	if err != nil {
		io.ActiveInputOutputHandler.NewLinef("Error finding absolute path for file [%v]: %v", itemsXMLPath, err)
		panic("ah")
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		print(fmt.Sprintf("Error loading file [%s]: %s", itemsXMLPath, err))
	}
	return itemsFromXML(bytes)
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

func itemsFromXML(xmlBytes []byte) map[string]things.Item {
	t := &XItems{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.ActiveInputOutputHandler.NewLinef("speechFromXML failed: %v", err)
	}
	m := make(map[string]things.Item, len(t.XItem))
	for _, i := range t.XItem {
		isVisible := i.IsVisible == "y"
		if i.IsEquippable == "y" {
			m[i.Name] = &equippable{
				&things.Thing{
					Name:     i.Name,
					LookText: i.LookText,
					Visible:  isVisible,
				},
			}
		} else {
			m[i.Name] = &item{
				&things.Thing{
					Name:     i.Name,
					LookText: i.LookText,
					Visible:  isVisible,
				},
			}
		}
	}
	return m
}

// XItems specifies the xml schema for a list of Items
type XItems struct {
	XItem []XItem
}

// XItem specifies the xml schema for an Item
type XItem struct {
	Name         string
	LookText     string
	IsVisible    string
	IsEquippable string
}

type item struct {
	t *things.Thing
	// Name         string
	// LookText     string
	// IsVisible    bool
	// IsEquippable bool
}

func (i *item) GetName() string {
	return i.t.Name
}

func (i *item) GetLookText() string {
	return i.t.LookText
}

func (i *item) GetThing() *things.Thing {
	return i.t
}

type equippable struct {
	t *things.Thing
	// Name         string
	// LookText     string
	// IsVisible    bool
	// IsEquippable bool
}

func (i *equippable) GetName() string {
	return i.t.Name
}

func (i *equippable) GetLookText() string {
	return i.t.LookText
}

func (i *equippable) GetThing() *things.Thing {
	return i.t
}

func (i *equippable) Toggle() {}
