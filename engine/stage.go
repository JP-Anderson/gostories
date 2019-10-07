package engine

import (
	"strings"

	"gostories/engine/action"
	"gostories/engine/state"
	"gostories/engine/inventory"
	"gostories/engine/logic"
	"gostories/engine/io"
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
			io.NewLine(s.state.CurrentArea.Look)
			newArea = false
		}
		// TODO: move the action parsing to another file/function
		inputAction, noun := io.SimpleParse()
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
			io.NewLine("You take stock of your inventory.")
			s.state.Inventory.PrintContents()
			io.NewLine("You have the following equipped:")
			s.state.EquippedItems.PrintContents()
		} else if inputAction.Name == "exit" {
			break
		} else {
			io.NewLine("Unknown action")
		}
		if targetedThing == nil {
			continue
		}
		trigger, ok := targetedThing.Triggers[inputAction.Name]; if ok {
			err := logic.EvaluateTrigger(s.state, trigger)
                        if err != nil {
                                io.NewLinef("Error evaluating trigger: %v", trigger)
                        }
		}
	}
}

func executeTravelCommand(travelTarget string, state *state.State) bool {
	trimmed := io.Trim(strings.ToLower(travelTarget))
	exit, exists := state.CurrentArea.Exits[area.Direction(trimmed)]
	if exists {
		state.CurrentArea = exit.To
		return true
	}
	io.NewLinef("Could not find an exit to the %v", trimmed)
	return false
}

func executeTalkCommand(talkTarget string, state state.State) {
	for _, being := range state.CurrentArea.Beings {
		io.NewLine(being.Name)
		if strings.ToLower(being.Name) == strings.ToLower(talkTarget) {
			io.NewLinef("You speak to %v.", being.Name)
			RunWithAlt(being.Speech, being.AltSpeech, state)
			return
		}
	}
	io.NewLinef("Could not find a %v to talk to!", talkTarget)
}

func executeTakeCommand(takeTarget string, state state.State) {
	//TODO refactor. item store already has a method to iterate its store by name
	for _, item := range state.CurrentArea.Items {
		if strings.ToLower(item.GetName()) == strings.ToLower(takeTarget) {
			if item.GetThing().Visible {
				io.NewLinef("You take the %v", item.GetName())
				state.Inventory.StoreItem(item)
				return
			}
		}
	}
	//TODO refactor. item store already has a method to iterate its store by name
	for _, feature := range state.CurrentArea.Features {
		if strings.ToLower(feature.GetName()) == strings.ToLower(takeTarget) {
			io.NewLinef("You can't really take the %v...", feature.GetName())
			return
		}
	}
	io.NewLinef("Couldn't find a %v to pick up.", takeTarget)
}

func executeEquipCommand(equipTarget string, state state.State) {
	defer func() {
		recover()
	}()
	item, err := state.Inventory.GetItemWithName(equipTarget)
	if err == nil {
		var itemInterface interface{}
		itemInterface = *item
		if ok := itemInterface.(things.Equippable); ok != nil {
			item, err := state.Inventory.RemoveItemWithName(equipTarget)
			if item != nil && err == nil {
				state.EquippedItems.StoreItem(*item)
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
