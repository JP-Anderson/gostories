package store

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/gen/items"
	"gostories/things"
)

func TestStoreItemIncrementsSize(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	assert.Equal(t, 0, i.Size())
	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
}

func TestContainsFindsStoredItem(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	i.StoreItem(item)
	assert.True(t, i.Contains(item))
}

func TestCannotAddSameItemTwice(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	assert.Equal(t, 0, i.Size())
	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
	assert.True(t, i.Contains(item))
	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
	assert.True(t, i.Contains(item))
}

func TestRemoveItemRemovesItem(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	assert.Equal(t, 0, i.Size())
	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
	assert.True(t, i.Contains(item))
	removed, err := i.RemoveItem(item)
	assert.NoError(t, err)
	assert.Equal(t, removed, &item)
	assert.Equal(t, 0, i.Size())
	assert.False(t, i.Contains(item))
}

func TestRemoveItemRemovesCorrectItem(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	item2 := getAnotherTestItem()

	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
	assert.True(t, i.Contains(item))

	i.StoreItem(item2)
	assert.Equal(t, 2, i.Size())
	assert.True(t, i.Contains(item))
	assert.True(t, i.Contains(item2))

	removed, err := i.RemoveItem(item)
	assert.NoError(t, err)
	assert.Equal(t, removed, &item)

	assert.Equal(t, 1, i.Size())
	assert.False(t, i.Contains(item))
	assert.True(t, i.Contains(item2))
}

func TestContainsMatchWithTextMatcher(t *testing.T) {
	i := NewInventory()
	shrubbery := getAnotherTestItem()
	shrubberyMatcher := func(item things.Item) bool {
		if item.GetName() == "shrubbery" {
			return true
		}
		return false
	}
	assert.False(t, i.ContainsMatch(shrubberyMatcher))

	i.StoreItem(shrubbery)
	assert.True(t, i.Contains(shrubbery))
	assert.True(t, i.ContainsMatch(shrubberyMatcher))

	_, err := i.RemoveItem(shrubbery)
	assert.NoError(t, err)
	assert.False(t, i.ContainsMatch(shrubberyMatcher))
}

func TestContainsMatchNegative(t *testing.T) {
	i := NewInventory()
	collar := getTestItem()
	shrubberyMatcher := func(item things.Item) bool {
		if item.GetName() == "shrubbery" {
			return true
		}
		return false
	}
	assert.False(t, i.ContainsMatch(shrubberyMatcher))

	i.StoreItem(collar)
	assert.True(t, i.Contains(collar))
	assert.False(t, i.ContainsMatch(shrubberyMatcher))
}

func TestEquippedItemsStoreItemDoesNotStoreNonEquippables(t *testing.T) {
	equippedItems := NewEquippedItems()
	output := equippedItems.StoreItem(items.ItemShrubbery)
	assert.False(t, output)
	assert.Equal(t, 0, equippedItems.Size())
}

func TestEquippedItemsStoreItemCanStoreEquippables(t *testing.T) {
	equippedItems := NewEquippedItems()
	output := equippedItems.StoreItem(items.ItemCollar)
	assert.True(t, output)
	assert.Equal(t, 1, equippedItems.Size())
}

func getTestItem() things.Item {
	return items.ItemCollar
}

func getAnotherTestItem() things.Item {
	return items.ItemShrubbery
}
