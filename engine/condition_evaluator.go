package engine

import (
	"strings"

	"gostories/things"
)

func EvaluateCondition(gameContext Context, conditionStr string) bool {
	conditionFunc := GetConditional(conditionStr)
	targetStr := conditionStr[strings.Index(conditionStr, "(")+1 : strings.Index(conditionStr, ")")]
	return conditionFunc(gameContext, targetStr)
}

var ConditionStringsMap = map[string]ConditionFn{
	"item-equipped": ConditionItemIsEquipped,
	"inventory-contains-item": ConditionInventoryContainsItem,
	//"item-in-inventory" : nil,
}

type ConditionFn = func(Context, string) bool

func GetConditional(conditionStr string) ConditionFn {
	conditionFuncStr := conditionStr[:strings.Index(conditionStr, "(")]
	condition := ConditionStringsMap[conditionFuncStr]
	return condition
}

func ConditionItemIsEquipped(ctx Context, itemName string) bool {
	return ctx.EquippedItems.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}

func ConditionInventoryContainsItem(ctx Context, itemName string) bool {
	return ctx.Inventory.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}
