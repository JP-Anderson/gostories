package engine

// EvaluateTrigger, given a string of format "TRIGGER(TARGET)", attempts to retrieve a trigger
// function mapped to the value of TRIGGER. If it finds a trigger function, it will attempt to 
// apply the trigger function on the noun/named object TARGET, which will have some side-effects on
// the provided GameContext. All trigger funcs can also return an error.
func EvaluateTrigger(gameContext Context, triggerStr string) error {
	triggerFunc := getTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameContext, targetStr)
}

type triggerFn = func(Context, string) error

func triggerRemoveItem(gameContext Context, itemName string) error {
	err := gameContext.Inventory.RemoveItemWithName(itemName)
	if err != nil {
		return err
	}
	err = gameContext.EquippedItems.RemoveItemWithName(itemName)
	return err
}

var triggerStringsMap = map[string]triggerFn{
	"remove-item": triggerRemoveItem,
}

func getTrigger(triggerStr string) triggerFn {
	triggerFuncStr := parseSingleValueFuncName(triggerStr)
	trigger := triggerStringsMap[triggerFuncStr]
	return trigger
}
