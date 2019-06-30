package engine

func EvaluateTrigger(gameContext Context, triggerStr string) error {
	triggerFunc := GetTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameContext, targetStr)
}

type TriggerFn = func(Context, string) error

func TriggerRemoveItem(gameContext Context, itemName string) error {
	err := gameContext.Inventory.RemoveItemWithName(itemName)
	if err != nil {
		return err
	}
	err = gameContext.EquippedItems.RemoveItemWithName(itemName)
	return err
}

var TriggerStringsMap = map[string]TriggerFn{
	"remove-item": TriggerRemoveItem,
}

func GetTrigger(triggerStr string) TriggerFn {
	triggerFuncStr := parseSingleValueFuncName(triggerStr)
	trigger := TriggerStringsMap[triggerFuncStr]
	return trigger
}
