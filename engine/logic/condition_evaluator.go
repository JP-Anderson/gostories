package logic

import (
	"gostories/engine/state"
	"gostories/things"
)

type conditionFn func(state.State, string) bool

// EvaluateCondition is given a string of format CONDITION(TARGET). CONDITION is matched to the name
// of an available condition function, if found, the function will be ran with the string TARGET as
// the parameter. A bool return signals if condition returns true or false for the target.
func EvaluateCondition(gameState state.State, conditionStr string) bool {
	conditionFunc := getConditional(conditionStr)
	targetStr := parseFuncParam(conditionStr)
	return conditionFunc(gameState, targetStr)
}

func getConditional(conditionStr string) conditionFn {
	conditionFuncStr := parseFuncName(conditionStr)
	condition := conditionStringsMap[conditionFuncStr]
	return condition
}

var conditionStringsMap = map[string]conditionFn{
	"item-equipped":           conditionItemIsEquipped,
	"inventory-contains-item": conditionInventoryContainsItem,
}

func conditionItemIsEquipped(ctx state.State, itemName string) bool {
	return ctx.EquippedItems.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}

func conditionInventoryContainsItem(ctx state.State, itemName string) bool {
	return ctx.Inventory.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}
