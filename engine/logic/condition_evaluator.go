package logic

import (
	"strings"

	"gostories/engine/context"
	"gostories/things"
)

type conditionFn = func(context.Context, string) bool

// EvaluateCondition is given a string of format CONDITION(TARGET). CONDITION is matched to the name
// of an available condition function, if found, the function will be ran with the string TARGET as
// the parameter. A bool return signals if condition returns true or false for the target.
func EvaluateCondition(gameContext context.Context, conditionStr string) bool {
	conditionFunc := getConditional(conditionStr)
	targetStr := parseFuncParam(conditionStr)
	return conditionFunc(gameContext, targetStr)
}

func getConditional(conditionStr string) conditionFn {
	conditionFuncStr := parseSingleValueFuncName(conditionStr)
	condition := conditionStringsMap[conditionFuncStr]
	return condition
}

func parseFuncParam(input string) string {
	return input[strings.Index(input, "(")+1 : strings.Index(input, ")")]
}

func parseSingleValueFuncName(input string) string {
	return input[:strings.Index(input, "(")]
}

var conditionStringsMap = map[string]conditionFn{
	"item-equipped":           conditionItemIsEquipped,
	"inventory-contains-item": conditionInventoryContainsItem,
}

func conditionItemIsEquipped(ctx context.Context, itemName string) bool {
	return ctx.EquippedItems.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}

func conditionInventoryContainsItem(ctx context.Context, itemName string) bool {
	return ctx.Inventory.ContainsMatch(func(item things.Item) bool {
		if item.GetName() == itemName {
			return true
		}
		return false
	})
}
