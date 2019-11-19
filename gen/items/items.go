package items

import (
	"encoding/xml"

	"gostories/engine/io"
	"gostories/things"
	"gostories/utils/strings"
	gxml "gostories/xml"
)

// Get returns any item which has a name matching the provided name.
func Get(name string) things.Item {
	return items[strings.ToIDString(name)]
}

var items = getItems()

func getItems() (i map[string]things.Item) {
	return loadFromXML()
}

func loadFromXML() (items map[string]things.Item) {
	return itemsFromXML(gxml.BytesForItems())
}

func itemsFromXML(xmlBytes []byte) map[string]things.Item {
	t := &Items{}
	err := xml.Unmarshal(xmlBytes, t)
	if err != nil {
		print("err here")
		io.Handler.NewLinef("itemsFromXML failed: %v", err)
	}
	m := make(map[string]things.Item, len(t.Item))
	for _, i := range t.Item {
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

// Items specifies the xml schema for a list of Items.
type Items struct {
	Item []Item
}

// Item specifies the xml schema for an Item.
type Item struct {
	Name         string
	LookText     string
	IsVisible    string
	IsEquippable string
}

type item struct {
	t *things.Thing
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
