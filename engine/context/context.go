package context

import (
	"gostories/engine/inventory"
	"gostories/things"
)

type Context struct {
	CurrentArea   things.Area
	Inventory     *inventory.Inventory
	EquippedItems *inventory.EquippedItems
}

