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

func (i *ItemStore) GetItemWithName(itemName string) (*things.Item, error) {
	index, err := i.getIndexForItemName(itemName)
        if err != nil {
                return nil, err
        }
	return i.getItemAtIndex(index)
}

func (i *ItemStore) getIndexForItemName(itemName string) (int, error) {
	index := findByName(i.items, itemName)
        if index < 0 {
                return -1, errors.New("item not present")
        }
	return index, nil
}

func (i *ItemStore) getItemAtIndex(index int) (*things.Item, error) {
        if index >= len(i.items) {
                return nil, errors.New("Index exceeds item slice length.")
        }
        if index < 0 {
                return nil, errors.New("Index cannot be less than 0")
        }
        item := i.items[index]
        return &item, nil
}

func (i *ItemStore) RemoveItem(item things.Item) (*things.Item, error) {
        return i.RemoveItemWithName(item.GetName())
}

func (i *ItemStore) RemoveItemWithName(itemName string) (*things.Item, error) {
        index, err := i.getIndexForItemName(itemName)
        if err != nil {
                return nil, err
        }
        return i.removeItemAtIndex(index)
}

func (i *ItemStore) removeItemAtIndex(indexToRemove int) (*things.Item, error) {
	item, err := i.getItemAtIndex(indexToRemove)
	if err != nil {
		return nil, errors.New("Could not find item to remove.")
	}
	copy(i.items[indexToRemove:], i.items[indexToRemove+1:])
	i.items[len(i.items)-1] = nil
	i.items = i.items[:len(i.items)-1]
	return item, nil
}

func findByName(items []things.Item, desiredItemName string) int {
	for i, item := range items {
		if desiredItemName == item.GetName() {
			return i
		}
	}
	return -1
}
