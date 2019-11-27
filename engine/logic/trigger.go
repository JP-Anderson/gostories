package logic

import (
	"fmt"
	"strconv"
	"strings"

	"gostories/engine/io"
	"gostories/engine/state"
	"gostories/gen/areas"
	"gostories/gen/items"
	"gostories/things/area"
)

// EvaluateTrigger when  given a string of format "TRIGGER(TARGET)", attempts to retrieve a trigger
// function mapped to the value of TRIGGER. If it finds a trigger function, it will attempt to
// apply the trigger function on the noun/named object TARGET, which will have some side-effects on
// the provided game State. All trigger funcs can also return an error.
func EvaluateTrigger(gameState *state.State, triggerStr string) error {
	if strings.Contains(triggerStr, ";") {
		evaluateMultiple(gameState, triggerStr)
		// TODO: potentially want to return wrapped errors here or a struct with all errors
		return nil
	}
	triggerFunc := getTrigger(triggerStr)
	targetStr := parseFuncParam(triggerStr)
	return triggerFunc(gameState, targetStr)
}

func evaluateMultiple(gameState *state.State, triggerStr string) {
	subTriggers := strings.Split(triggerStr, ";")
	for _, trigger := range subTriggers {
		EvaluateTrigger(gameState, trigger)
	}
}

func getTrigger(triggerStr string) triggerFn {
	triggerFuncStr := parseFuncName(triggerStr)
	trigger := triggerStringsMap[triggerFuncStr]
	return trigger
}

type triggerFn func(*state.State, string) error

func triggerRemoveItem(gameState *state.State, itemName string) error {
	_, err1 := gameState.Inventory.RemoveItemWithName(itemName)
	_, err2 := gameState.EquippedItems.RemoveItemWithName(itemName)
	if err1 != nil && err2 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func triggerRevealItem(gameState *state.State, itemName string) error {
	item := area.CheckItems(gameState.CurrentArea, itemName)
	if item != nil {
		if item.Visible {
			io.Handler.NewLine(itemName + "is already visible")
		} else {
			item.Show()
			io.Handler.NewLinef("Revealing item %v", itemName)
		}
	}
	return nil
}

func triggerAddItem(gameState *state.State, itemName string) error {
	i := items.Get(itemName)
	if i != nil {
		gameState.Inventory.StoreItem(i)
		io.Handler.NewLinef("You now have the %v", itemName)
	}
	return nil
}

func triggerAddExit(gameState *state.State, input string) error {
	stringSlice := strings.Split(input, ",")
	areaName := strings.TrimSpace(stringSlice[1])
	a := areas.Get(areaName)
	if a != nil {
		dir := area.StringToDirection[strings.TrimSpace(stringSlice[0])]
		exit := area.Exit{
			To:   a,
			From: gameState.CurrentArea,
		}
		gameState.CurrentArea.Exits[dir] = exit
		reverseExit := area.Exit{
			To:   gameState.CurrentArea,
			From: a,
		}
		a.Exits[area.OppositeDirection[strings.TrimSpace(stringSlice[0])]] = reverseExit
		return nil
	}
	return fmt.Errorf("could not find area %s", stringSlice[1])
}

func triggerChangeLookText(gameState *state.State, input string) error {
	strs := strings.SplitN(input, ",", 2)
	if len(strs) != 2 {
		return fmt.Errorf("input must be of format (int, string). was: %s", input)
	}
	str1 := strings.TrimSpace(strs[0])
	var i int
	i, err := strconv.Atoi(str1)
	if err != nil {
		return fmt.Errorf("first param must be an integer. was: %s", str1)
	}
	currentArea := gameState.CurrentArea
	currentArea.ChangeLookText(i)
	io.Handler.NewLine(strings.TrimSpace(strs[1]))
	return nil
}

var triggerStringsMap = map[string]triggerFn{
	"reveal-item": triggerRevealItem,
	"remove-item": triggerRemoveItem,
	"add-item":    triggerAddItem,
	"add-exit":    triggerAddExit,
	"change-look-text": triggerChangeLookText,
}
