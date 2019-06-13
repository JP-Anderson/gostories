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

func getTestItem() things.Item {
	return things.CatCollarItem{}
}
