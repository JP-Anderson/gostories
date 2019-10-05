package engine

import (
	"gostories/engine/context"
	"gostories/engine/io"
	"gostories/gen/items"
)

// EvaluateTrigger, given a string of format "TRIGGER(TARGET)", attempts to retrieve a trigger
// function mapped to the value of TRIGGER. If it finds a trigger function, it will attempt to 
// apply the trigger function on the noun/named object TARGET, which will have some side-effects on
// the provided GameContext. All trigger funcs can also return an error.
func EvaluateTrigger(gameContext context.Context, triggerStr string) error {
	triggerFunc := getTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameContext, targetStr)
}

type triggerFn = func(context.Context, string) error

func triggerRemoveItem(gameContext context.Context, itemName string) error {
	err := gameContext.Inventory.RemoveItemWithName(itemName)
	if err != nil {
		return err
	}
	err = gameContext.EquippedItems.RemoveItemWithName(itemName)
	return err
}

func triggerRevealItem(gameContext context.Context, itemName string) error {
	io.NewLinef("Revealing item %v", itemName)
	item := gameContext.CurrentArea.CheckAreaItemsForThing(itemName); if item != nil {
		if item.Visible {
			io.NewLine(itemName + "is already visible")
		} else {
			item.Show()
		}
	}
	return nil
}

func triggerAddItem(gameContext context.Context, itemName string) error {
	io.NewLinef("%#v", items.Items)
	io.NewLinef("Looking for %v", itemName)
	i, ok := items.Items[itemName]; if ok {
		io.NewLine("Found!")
		gameContext.Inventory.StoreItem(i)
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
