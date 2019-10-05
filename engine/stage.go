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
			trimmed := io.Trim(strings.ToLower(noun))
			exit, exists := s.context.CurrentArea.Exits[things.Direction(trimmed)]
			if exists {
				s.context.CurrentArea = *exit.To
				newArea = true
			} else {
				io.NewLinef("Could not find an exit to the %v", noun)
			}
		} else if action.Name == "talk" {
			found := false
			for _, being := range s.context.CurrentArea.Beings {
				io.NewLine(being.Name)
				if strings.ToLower(being.Name) == strings.ToLower(noun) {
					found = true
					io.NewLinef("You speak to %v.", being.Name)
					RunWithAlt(being.Speech, being.AltSpeech, s.context)
				}
			}
			if !found {
				io.NewLinef("Could not find a %v to talk to!", noun)
			}
		} else if action.Name == "take" {
			found := false
			for _, item := range s.context.CurrentArea.Items {
				if strings.ToLower(item.GetName()) == strings.ToLower(noun) {
					found = true
					io.NewLinef("You take the %v", item.GetName())
					s.context.Inventory.StoreItem(item)
				}
			}
			for _, feature := range s.context.CurrentArea.Features {
				if strings.ToLower(feature.GetName()) == strings.ToLower(noun) {
					found = true
					io.NewLinef("You can't really take the %v...", feature.GetName())
				}
			}
			if !found {
				io.NewLinef("Couldn't find a %v to pick up.", noun)
			}
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
