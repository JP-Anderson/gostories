package engine

import (
	"strings"

	"gostories/engine/context"
	"gostories/engine/inventory"
	"gostories/engine/logic"
	"gostories/engine/io"
	"gostories/things"
)

// Stage holds the main game context and the game control loop.
type Stage struct {
	context context.Context
}

// Start initialise a Stage, it is passed an Area, which is used to initialise the game context.
func (s Stage) Start(area things.Area) {
	s.context = context.Context{
		CurrentArea:   area,
		Inventory:     inventory.NewInventory(),
		EquippedItems: inventory.NewEquippedItems(),
	}
	s.loopUntilExit()
}

func (s Stage) loopUntilExit() {
	newArea := true
	for {
		if newArea {
			io.NewLine(s.context.CurrentArea.Look)
			newArea = false
		}
		// TODO: move the action parsing to another file/function
		action, noun := io.SimpleParse()
		// TODO: set targetedThing to every noun item. Refactor in the process!
		var targetedThing *things.Thing
		if action.Name == "look" {
			targetedThing = executeLookCommand(noun, s.context)
		} else if action.Name == "travel" {
			newArea = executeTravelCommand(noun, &s.context)
		} else if action.Name == "talk" {
			executeTalkCommand(noun, s.context)
		} else if action.Name == "take" {
			executeTakeCommand(noun, s.context)
		} else if action.Name == "equip" {
			executeEquipCommand(noun, s.context)
		} else if action.Name == "inventory" {
			io.NewLine("You take stock of your inventory.")
			s.context.Inventory.PrintContents()
			io.NewLine("You have the following equipped:")
			s.context.EquippedItems.PrintContents()
		} else if action.Name == "exit" {
			break
		} else {
			io.NewLine("Unknown action")
		}
		if targetedThing == nil {
			continue
		}
		trigger, ok := targetedThing.Triggers[action.Name]; if ok {
			err := logic.EvaluateTrigger(s.context, trigger)
                        if err != nil {
                                io.NewLinef("Error evaluating trigger: %v", trigger)
                        }
		}
	}
}

func executeLookCommand(lookTarget string, context context.Context) (target *things.Thing) {
	defer func() {
		if target != nil {
			io.NewLine(target.LookText)
		}
	}()

	if lookTarget == "" {
		io.NewLine(context.CurrentArea.Look)
	}

	target = context.CurrentArea.CheckAreaItemsForThing(lookTarget)
	if target != nil {
		return
	}

	target = context.CurrentArea.CheckAreaFeaturesForThing(lookTarget)
	if target != nil {
		return
	}

	target = context.CurrentArea.CheckAreaBeingsForThing(lookTarget)
	if target != nil {
		return
	}

	io.NewLinef("Couldn't find a %v to look at!", lookTarget)
	return
}

func executeTravelCommand(travelTarget string, context *context.Context) bool {
	trimmed := io.Trim(strings.ToLower(travelTarget))
	exit, exists := context.CurrentArea.Exits[things.Direction(trimmed)]
	if exists {
		context.CurrentArea = *exit.To
		return true
	} else {
		io.NewLinef("Could not find an exit to the %v", trimmed)
	}
	return false
}

func executeTalkCommand(talkTarget string, context context.Context) {
	for _, being := range context.CurrentArea.Beings {
		io.NewLine(being.Name)
		if strings.ToLower(being.Name) == strings.ToLower(talkTarget) {
			io.NewLinef("You speak to %v.", being.Name)
			RunWithAlt(being.Speech, being.AltSpeech, context)
			return
		}
	}
	io.NewLinef("Could not find a %v to talk to!", talkTarget)
}

func executeTakeCommand(takeTarget string, context context.Context) {
	//TODO refactor. item store already has a method to iterate its store by name
	for _, item := range context.CurrentArea.Items {
		if strings.ToLower(item.GetName()) == strings.ToLower(takeTarget) {
			io.NewLinef("You take the %v", item.GetName())
			context.Inventory.StoreItem(item)
			return
		}
	}
	//TODO refactor. item store already has a method to iterate its store by name
	for _, feature := range context.CurrentArea.Features {
		if strings.ToLower(feature.GetName()) == strings.ToLower(takeTarget) {
			io.NewLinef("You can't really take the %v...", feature.GetName())
			return
		}
	}
	io.NewLinef("Couldn't find a %v to pick up.", takeTarget)
}

func executeEquipCommand(equipTarget string, context context.Context) {
	defer func() {
		recover()
	}()
	item, err := context.Inventory.GetItemWithName(equipTarget)
	if err == nil {
		var itemInterface interface{}
		itemInterface = *item
		if ok := itemInterface.(things.Equippable); ok != nil {
			item, err := context.Inventory.RemoveItemWithName(equipTarget)
			if item != nil && err == nil {
				context.EquippedItems.StoreItem(*item)
			} else {
				io.NewLinef("Failed to equip item...")
			}
		} else {
			io.NewLinef("How do you expect to equip the %v?", equipTarget)
		}
	} else {
		io.NewLinef("%v", err)
		io.NewLinef("Do not have a %v to equip.", equipTarget)
	}
}
