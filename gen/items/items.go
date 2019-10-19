package items

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gostories/engine/io"
	"gostories/things"
	"io/ioutil"
)

func Item(name string) things.Item {
	return Items()[name]
}

func Items() map[string]things.Item {
	return items
}

func getItems() (i map[string]things.Item) {
	return loadFromXML()
}

// Items contains all the items a player can pick up. It is currently indexed by the item name, however ideally
// it should be indexed by an enum/custom int and this should be the only place to access items.
var items = getItems()

type itemToAdd struct {
	Name         string
	LookText     string
	IsVisible    bool
	IsEquippable bool
}

type XItems struct {
	XItem []XItem
}

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

func buildItemsMap() (items map[string]things.Item) {
	itemsToAdd := []itemToAdd{
		itemToAdd{"collar", "A small red cat collar with a bell.", false, true},
		itemToAdd{"shrubbery", "A small but rather well cared for shrubbery.", true, false},
		itemToAdd{"sardines", "A tin of tasty sardines preserved in olive oil.", true, false},
	}
	items = make(map[string]things.Item, len(itemsToAdd))
	for _, itemToAdd := range itemsToAdd {
		if itemToAdd.IsEquippable {
			items[itemToAdd.Name] = &equippable{
				&things.Thing{
					Name:     itemToAdd.Name,
					LookText: itemToAdd.LookText,
					Visible:  itemToAdd.IsVisible,
				},
			}
		} else {
			items[itemToAdd.Name] = &item{
				&things.Thing{
					Name:     itemToAdd.Name,
					LookText: itemToAdd.LookText,
					Visible:  itemToAdd.IsVisible,
				},
			}
		}

	}
	return
}

func getRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Check if already in root directory (e.g. main.go)
	lastSlashIndex := strings.LastIndex(dir, "/")
	cName := dir[lastSlashIndex+1:]
	if cName == "gostories" {
		return dir
	}

	// Get parent directory
	parent := filepath.Dir(dir)
	lastSlashIndex = strings.LastIndex(parent, "/")
	pName := parent[lastSlashIndex+1:]

	// If not at root, continue getting parent
	for pName != "gostories" {
		parent = filepath.Dir(parent)
		lastSlashIndex = strings.LastIndex(parent, "/")
		pName = parent[lastSlashIndex+1:]
	}
	return parent
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
