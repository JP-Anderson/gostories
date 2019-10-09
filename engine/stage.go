package engine

import (
	"strings"

	"gostories/engine/action"
	"gostories/engine/state"
	"gostories/engine/inventory"
	"gostories/engine/logic"
	"gostories/engine/io"
	"gostories/engine/io/console"
	"gostories/things"
	"gostories/things/area"
)

// Stage holds the main game state and the game control loop.
type Stage struct {
	state state.State
}

// Start initialise a Stage, it is passed an Area, which is used to initialise the game state.
func (s Stage) Start(area area.Area) {
	s.state = state.State{
		CurrentArea:   &area,
		Inventory:     inventory.NewInventory(),
		EquippedItems: inventory.NewEquippedItems(),
	}
	s.loopUntilExit()
}

func (s Stage) loopUntilExit() {
	newArea := true
	for {
		if newArea {
			io.ActiveInputOutputHandler.NewLine(s.state.CurrentArea.Look)
			newArea = false
		}
		// TODO: move the action parsing to another file/function
		inputAction, noun := io.ActiveInputOutputHandler.SimpleParse()
		// TODO: set targetedThing to every noun item. Refactor in the process!
		var targetedThing *things.Thing
		if inputAction.Name == "look" {
			targetedThing = action.ExecuteLookCommand(noun, s.state)
		} else if inputAction.Name == "travel" {
			newArea = executeTravelCommand(noun, &s.state)
		} else if inputAction.Name == "talk" {
			executeTalkCommand(noun, s.state)
		} else if inputAction.Name == "take" {
			executeTakeCommand(noun, s.state)
		} else if inputAction.Name == "equip" {
			executeEquipCommand(noun, s.state)
		} else if inputAction.Name == "inventory" {
			io.ActiveInputOutputHandler.NewLine("You take stock of your inventory.")
			s.state.Inventory.PrintContents()
			io.ActiveInputOutputHandler.NewLine("You have the following equipped:")
			s.state.EquippedItems.PrintContents()
		} else if inputAction.Name == "exit" {
			break
		} else {
			io.ActiveInputOutputHandler.NewLine("Unknown action")
		}
		if targetedThing == nil {
			continue
		}
		trigger, ok := targetedThing.Triggers[inputAction.Name]; if ok {
			err := logic.EvaluateTrigger(s.state, trigger)
                        if err != nil {
                                io.ActiveInputOutputHandler.NewLinef("Error evaluating trigger: %v", trigger)
                        }
		}
	}
}

func executeTravelCommand(travelTarget string, state *state.State) bool {
	trimmed := consoleio.Trim(strings.ToLower(travelTarget))
	exit, exists := state.CurrentArea.Exits[area.Direction(trimmed)]
	if exists {
		state.CurrentArea = exit.To
		return true
	}
	io.ActiveInputOutputHandler.NewLinef("Could not find an exit to the %v", trimmed)
	return false
}

func executeTalkCommand(talkTarget string, state state.State) {
	for _, being := range state.CurrentArea.Beings {
		io.ActiveInputOutputHandler.NewLine(being.Name)
		if strings.ToLower(being.Name) == strings.ToLower(talkTarget) {
			io.ActiveInputOutputHandler.NewLinef("You speak to %v.", being.Name)
			RunWithAlt(being.Speech, being.AltSpeech, state)
			return
		}
	}
	io.ActiveInputOutputHandler.NewLinef("Could not find a %v to talk to!", talkTarget)
}

func executeTakeCommand(takeTarget string, state state.State) {
	//TODO refactor. item store already has a method to iterate its store by name
	for _, item := range state.CurrentArea.Items {
		if strings.ToLower(item.GetName()) == strings.ToLower(takeTarget) {
			if item.GetThing().Visible {
				io.ActiveInputOutputHandler.NewLinef("You take the %v", item.GetName())
				state.Inventory.StoreItem(item)
				return
			}
		}
	}
	//TODO refactor. item store already has a method to iterate its store by name
	for _, feature := range state.CurrentArea.Features {
		if strings.ToLower(feature.GetName()) == strings.ToLower(takeTarget) {
			io.ActiveInputOutputHandler.NewLinef("You can't really take the %v...", feature.GetName())
			return
		}
	}
	io.ActiveInputOutputHandler.NewLinef("Couldn't find a %v to pick up.", takeTarget)
}

func executeEquipCommand(equipTarget string, state state.State) {
	item, err := state.Inventory.GetItemWithName(equipTarget)
	if err == nil {
		var itemInterface interface{}
		itemInterface = *item
		_, ok := itemInterface.(things.Equippable); if ok {
			item, err := state.Inventory.RemoveItemWithName(equipTarget)
			if item != nil && err == nil {
				state.EquippedItems.StoreItem(*item)
			} else {
				io.ActiveInputOutputHandler.NewLinef("Failed to equip item...")
			}
		} else {
			io.ActiveInputOutputHandler.NewLinef("How do you expect to equip the %v?", equipTarget)
		}
	} else {
		io.ActiveInputOutputHandler.NewLinef("%v", err)
		io.ActiveInputOutputHandler.NewLinef("Do not have a %v to equip.", equipTarget)
	}
}
