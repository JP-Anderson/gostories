package context

import (
	"gostories/engine/io"
	"gostories/engine/inventory"
	"gostories/things"
)

type Context struct {
	CurrentArea   things.Area
	Inventory     *inventory.Inventory
	EquippedItems *inventory.EquippedItems
}

func (c *Context) Equip(item things.Item) {
	if !c.Inventory.Contains(item) {
		io.NewLine("Cannot equip item not in inventory.")
		return
	}

	if c.EquippedItems.Contains(item) {
		io.NewLinef("Cannot equip %v as this item is already equipped!", item.GetName())
		return
	}

	err := c.Inventory.RemoveItem(item)
	if err != nil {
		io.NewLinef("Error removing item %v: %v", item.GetName(), err)
		return
	}
	c.EquippedItems.StoreItem(item)
}
