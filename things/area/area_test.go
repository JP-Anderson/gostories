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

	theThing := testArea.CheckAreaItemsForThing("collar")
	assert.NotNil(t, theThing)

	theThing.Show()

	theThing2 := testArea.CheckAreaItemsForThing("collar")
	assert.True(t, theThing2.Visible)
}

func getTestItem() things.Item {
	return items.ItemCollar
}
