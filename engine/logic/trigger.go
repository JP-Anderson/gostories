package logic

import (
	"gostories/engine/state"
	"gostories/engine/io"
	"gostories/gen/items"
)

// EvaluateTrigger, given a string of format "TRIGGER(TARGET)", attempts to retrieve a trigger
// function mapped to the value of TRIGGER. If it finds a trigger function, it will attempt to 
// apply the trigger function on the noun/named object TARGET, which will have some side-effects on
// the provided game State. All trigger funcs can also return an error.
func EvaluateTrigger(gameState state.State, triggerStr string) error {
	triggerFunc := getTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameState, targetStr)
}

type triggerFn = func(state.State, string) error

func triggerRemoveItem(gameState state.State, itemName string) error {
	_, err := gameState.Inventory.RemoveItemWithName(itemName)
	if err == nil {
		// this should be fine as long as item is always removed from inventory on equip
		return nil
	}
	_, err = gameState.EquippedItems.RemoveItemWithName(itemName)
	if err != nil {
		return err
	}
	return nil
}

func triggerRevealItem(gameState state.State, itemName string) error {
	io.NewLinef("Revealing item %v", itemName)
	item := gameState.CurrentArea.CheckAreaItemsForThing(itemName); if item != nil {
		if item.Visible {
			io.NewLine(itemName + "is already visible")
		} else {
			item.Show()
		}
	}
	return nil
}

func triggerAddItem(gameState state.State, itemName string) error {
	io.NewLinef("%#v", items.Items)
	io.NewLinef("Looking for %v", itemName)
	i, ok := items.Items[itemName]; if ok {
		io.NewLine("Found!")
		gameState.Inventory.StoreItem(i)
	}
	return nil
}

var triggerStringsMap = map[string]triggerFn{
	"reveal-item": triggerRevealItem,
	"remove-item": triggerRemoveItem,
	"add-item":    triggerAddItem,
}

func getTrigger(triggerStr string) triggerFn {
	triggerFuncStr := parseSingleValueFuncName(triggerStr)
	trigger := triggerStringsMap[triggerFuncStr]
	return trigger
}