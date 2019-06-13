package engine

import (
	"gostories/things"
)

type Inventory struct {
	items []things.Item
}

func NewInventory() *Inventory {
	return &Inventory{}
}

func (i *Inventory) Size() int {
	return len(i.items)
}

func (i *Inventory) StoreItem(newItem things.Item) {
	if i.Contains(newItem) {
		return
	}
	i.items = append(i.items, newItem)
}

func (i *Inventory) Contains(desiredItem things.Item) bool {
	for _, item := range i.items {
		if item.GetName() == desiredItem.GetName() {
			return true
		}
	}
	return false
}