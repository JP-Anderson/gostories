package area

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/gen/items"
	"gostories/things"
)

func TestItemMadeVisibleInAreaStaysVisible(t *testing.T) {
	item := getTestItem()
	assert.False(t, item.GetThing().Visible)
	testArea := Area{
		Items: []things.Item{
			item,
		},
	}

	ref1 := testArea.CheckAreaItemsForThing("collar")
	assert.NotNil(t, ref1)

	ref1.Show()

	ref2 := testArea.CheckAreaItemsForThing("collar")

	assert.True(t, ref1.Visible)
	assert.True(t, ref2.Visible)
}

func getTestItem() things.Item {
	return items.ItemCollar
}
