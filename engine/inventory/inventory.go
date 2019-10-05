package inventory

import (
	"errors"

	"gostories/engine/io"
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

// PrintContents prints the name of each item in the store
func (i *ItemStore) PrintContents() {
	if i.Size() > 0 {
		for _, item := range i.items {
			io.NewLine(item.GetName())
		}
	} else {
		io.NewLinef("Empty.")
	}
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

func (i *ItemStore) ContainsMatch(matcher func(item things.Item) bool) bool {
	for _, item := range i.items {
		if matcher(item) {
			return true
		}
	}
	return false
}

func (i *ItemStore) RemoveItem(item things.Item) error {
	return i.removeItemAtIndex(find(i.items, item))
}

func (i *ItemStore) RemoveItemWithName(itemName string) error {
	return i.removeItemAtIndex(findByName(i.items, itemName))
}

func (i *ItemStore) removeItemAtIndex(indexToRemove int) error {
	if indexToRemove >= len(i.items) {
		return errors.New("Could not find item to remove.")
	}
	if indexToRemove < 0 {
		return nil
	}
	copy(i.items[indexToRemove:], i.items[indexToRemove+1:])
	i.items[len(i.items)-1] = nil
	i.items = i.items[:len(i.items)-1]
	return nil
}

func find(items []things.Item, desiredItem things.Item) int {
	return findByName(items, desiredItem.GetName())
}

func findByName(items []things.Item, desiredItemName string) int {
	for i, item := range items {
		if desiredItemName == item.GetName() {
			return i
		}
	}
	return -1
}

