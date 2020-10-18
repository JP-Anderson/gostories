package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gostories/engine/io"
	mockio "gostories/engine/io/mock"
	"gostories/gen/items"
	"gostories/things"
)

type StoreTestSuite struct {
	suite.Suite
}

func (s *StoreTestSuite) SetupSuite() {
	io.Handler = mockio.NewMockHandler()
	println("HERE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func (s *StoreTestSuite) TestStoreItemIncrementsSize(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	assert.Equal(t, 0, i.Size())
	i.StoreItem(item)
	assert.Equal(t, 1, i.Size())
}

func (s *StoreTestSuite) TestContainsFindsStoredItem(t *testing.T) {
	i := NewInventory()
	item := getTestItem()
	i.StoreItem(item)
	assert.True(t, i.Contains(item))
}

func (s *StoreTestSuite) TestCannotAddSameItemTwice(t *testing.T) {
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

func (s *StoreTestSuite) TestRemoveItemRemovesItem(t *testing.T) {
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

func (s *StoreTestSuite) TestRemoveItemRemovesCorrectItem(t *testing.T) {
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

func (s *StoreTestSuite) TestContainsMatchWithTextMatcher(t *testing.T) {
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

func (s *StoreTestSuite) TestContainsMatchNegative(t *testing.T) {
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

func (s *StoreTestSuite) TestEquippedItemsStoreItemDoesNotStoreNonEquippables(t *testing.T) {
	equippedItems := NewEquippedItems()
	output := equippedItems.StoreItem(getAnotherTestItem())
	assert.False(t, output)
	assert.Equal(t, 0, equippedItems.Size())
}

func (s *StoreTestSuite) TestEquippedItemsStoreItemCanStoreEquippables(t *testing.T) {
	equippedItems := NewEquippedItems()
	output := equippedItems.StoreItem(getTestItem())
	assert.True(t, output)
	assert.Equal(t, 1, equippedItems.Size())
}

func getTestItem() things.Item {
	return items.Get("collar")
}

func getAnotherTestItem() things.Item {
	return items.Get("shrubbery")
}
