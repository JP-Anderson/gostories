package engine

import (
	"errors"

	"gostories/things"
)

type Inventory struct {
	ItemStore
}

func NewInventory() *Inventory {
	return &Inventory{}
}

// TODO: Introduce limited equipment slots
type EquippedItems struct {
	ItemStore
}

func NewEquippedItems() *EquippedItems {
	return &EquippedItems{}
}

type ItemStore struct {
	items []things.Item
}

func (i *ItemStore) Size() int {
	return len(i.items)
}

func (i *ItemStore) StoreItem(newItem things.Item) {
	if i.Contains(newItem) {
		return
	}
	i.items = append(i.items, newItem)
}

func (i *ItemStore) Contains(desiredItem things.Item) bool {
	for _, item := range i.items {
		if item.GetName() == desiredItem.GetName() {
			return true
		}
	}
	return false
}

func (i *ItemStore) RemoveItem(item things.Item) error {
	indexToRemove := find(i.items, item)
	if indexToRemove >= len(i.items) {
		return errors.New("Could not find item to remove.")
	}
	copy(i.items[indexToRemove:], i.items[indexToRemove+1:])
	i.items[len(i.items)-1] = nil
	i.items = i.items[:len(i.items)-1]
	return nil
}

// find returns the smallest index i at which the desiredItem == a[i], or -1 if there is no such index.
func find(items []things.Item, desiredItem things.Item) int {
	for i, item := range items {
		if desiredItem.GetName() == item.GetName() {
			return i
		}
	}
	return -1
}
