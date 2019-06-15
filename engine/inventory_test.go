package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
	err := i.RemoveItem(item)
	assert.NoError(t, err)
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

	err := i.RemoveItem(item)
	assert.NoError(t, err)

	assert.Equal(t, 1, i.Size())
	assert.False(t, i.Contains(item))
	assert.True(t, i.Contains(item2))
}

func getTestItem() things.Item {
	return things.CatCollarItem{}
}

func getAnotherTestItem() things.Item {
	return things.ShrubberyItem{}
}
