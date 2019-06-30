package engine

import (
	"strings"

	"gostories/things"
)

type ConditionFn = func(Context, string) bool

func EvaluateCondition(gameContext Context, conditionStr string) bool {
	conditionFunc := GetConditional(conditionStr)
	targetStr := parseFuncParam(conditionStr)
	return conditionFunc(gameContext, targetStr)
}

func GetConditional(conditionStr string) ConditionFn {
	conditionFuncStr := parseSingleValueFuncName(conditionStr)
	condition := ConditionStringsMap[conditionFuncStr]
	return condition
}

func parseFuncParam(input string) string {
	return input[strings.Index(input, "(")+1 : strings.Index(input, ")")]
}

func parseSingleValueFuncName(input string) string {
	return input[:strings.Index(input, "(")]
}

var ConditionStringsMap = map[string]ConditionFn{
	"item-equipped":           ConditionItemIsEquipped,
	"inventory-contains-item": ConditionInventoryContainsItem,
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
