package area

import (
	"gostories/engine/store"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/gen/features"
	"gostories/gen/items"
	"gostories/things"
)

func TestItemMadeVisibleInAreaStaysVisible(t *testing.T) {
	item := getTestItem()
	assert.NotNil(t, item)
	assert.False(t, item.GetThing().Visible)
	testArea := Area{
		Items: store.NewItemStore(),
	}
	testArea.Items.StoreItem(item)

	ref1 := testArea.FindItemByName("collar")
	assert.NotNil(t, ref1)

	ref1.GetThing().Show()

	ref2 := testArea.FindItemByName("collar")

	assert.True(t, ref1.GetThing().Visible)
	assert.True(t, ref2.GetThing().Visible)
}

func TestFeaturesChecker(t *testing.T) {
	area := NewArea()
	stand := features.Get("stand")
	area.Features = append(area.Features, stand)

	t.Run("matches name", func (t *testing.T) {
		output := area.CheckAreaForThing("stand", CheckFeatures)
		assert.NotNil(t, output)
	})

	t.Run("case insensitive", func (t *testing.T) {
		output := area.CheckAreaForThing("sTAnd", CheckFeatures)
		assert.NotNil(t, output)
	})

	t.Run("no match for missing feature", func (t *testing.T) {
		output := area.CheckAreaForThing("shelf", CheckFeatures)
		assert.Nil(t, output)
	})
}

func getTestItem() things.Item {
	return items.Get("collar")
}
