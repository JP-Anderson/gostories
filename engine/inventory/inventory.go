package inventory

import (
	"errors"

	"gostories/engine/io"
	"gostories/things"
)

// Inventory is an ItemStore for storing items a player is currently carrying. If the item is equipped it
// will be removed from the Inventory and stored in the EquippedItems ItemStore.
type Inventory struct {
	ItemStore
}

// NewInventory returns a new Inventory ItemStore
func NewInventory() *Inventory {
	return &Inventory{}
}

// EquippedItems is an ItemStore for storing equippable Items the player has equipped. Currently it just contains
// a regular ItemStore.
// TODO: Introduce limited equipment slots
// TODO: Enforce this to store only Equippables rather than making users of this struct handle this.
type EquippedItems struct {
	ItemStore
}

// NewEquippedItems returns a new ItemStore for storing items the player has equipped.
func NewEquippedItems() *EquippedItems {
	return &EquippedItems{}
}

// ItemStore is a wrapper around a slice of things.Item, which provides helper methods for adding and removing Items
// from the slice.
type ItemStore struct {
	items []things.Item
}

// PrintContents prints the name of each Item in the ItemStore.
func (i *ItemStore) PrintContents() {
	if i.Size() > 0 {
		for _, item := range i.items {
			io.ActiveInputOutputHandler.NewLine(item.GetName())
		}
	} else {
		io.ActiveInputOutputHandler.NewLinef("Empty.")
	}
}

// Size returns the number of Items currently stored in the ItemStore.
func (i *ItemStore) Size() int {
	return len(i.items)
}

// StoreItem takes an Item struct and places it in the ItemStore. Currently, only one of each Item can be added
// to the ItemStore.
func (i *ItemStore) StoreItem(newItem things.Item) {
	if i.Contains(newItem) {
		return
	}
	i.items = append(i.items, newItem)
}

// Contains returns true if any item in the ItemStore has a name which matches the given Item parameters name.
// It basically wraps GetItemWithName.
func (i *ItemStore) Contains(desiredItem things.Item) bool {
	item, err := i.GetItemWithName(desiredItem.GetName())
	if item != nil && err == nil {
		return true
	}
	return false
}

// ContainsMatch takes a func which takes an item and returns a bool. It will iterate through the items in the
// ItemStore and return true if any of the Items return true for the given "matcher".
func (i *ItemStore) ContainsMatch(matcher func(item things.Item) bool) bool {
	for _, item := range i.items {
		if matcher(item) {
			return true
		}
	}
	return false
}

// GetItemWithName returns a pointer to the first item in the ItemStore which name matches the provided name, or it
// returns an error.
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

// RemoveItem takes an Item and removes the first Item with that name from the ItemStore, or it returns an error.
func (i *ItemStore) RemoveItem(item things.Item) (*things.Item, error) {
        return i.RemoveItemWithName(item.GetName())
}

// RemoveItemWithName removes the first Item matching the provided name from the ItemStore, or it returns an error.
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

