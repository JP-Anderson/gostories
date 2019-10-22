package logic

import (
	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/gen/items"
	"gostories/things/area"
)

// EvaluateTrigger when  given a string of format "TRIGGER(TARGET)", attempts to retrieve a trigger
// function mapped to the value of TRIGGER. If it finds a trigger function, it will attempt to
// apply the trigger function on the noun/named object TARGET, which will have some side-effects on
// the provided game State. All trigger funcs can also return an error.
func EvaluateTrigger(gameState state.State, triggerStr string) error {
	triggerFunc := getTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameState, targetStr)
}

func getTrigger(triggerStr string) triggerFn {
	triggerFuncStr := parseFuncName(triggerStr)
	trigger := triggerStringsMap[triggerFuncStr]
	return trigger
}

type triggerFn func(state.State, string) error

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
	io.ActiveInputOutputHandler.NewLinef("Revealing item %v", itemName)
	item := area.CheckItems(gameState.CurrentArea, itemName)
	if item != nil {
		if item.Visible {
			io.ActiveInputOutputHandler.NewLine(itemName + "is already visible")
		} else {
			item.Show()
		}
	}
	return nil
}

func triggerAddItem(gameState state.State, itemName string) error {
	i := items.Get(itemName)
	if i != nil {
		gameState.Inventory.StoreItem(i)
	}
	return nil
}

var triggerStringsMap = map[string]triggerFn{
	"reveal-item": triggerRevealItem,
	"remove-item": triggerRemoveItem,
	"add-item":    triggerAddItem,
}
