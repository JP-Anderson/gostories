package area

import (
	"gostories/engine/store"
	"testing"

	"github.com/stretchr/testify/assert"

	items "gostories/gen/items/v2"
	"gostories/things"
)

func TestItemMadeVisibleInAreaStaysVisible(t *testing.T) {
	item := getTestItem()
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

func getTestItem() things.Item {
	return items.Items["collar"]
}
