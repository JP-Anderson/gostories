package engine

import (
	"strings"

	"gostories/engine/action"
	"gostories/engine/inventory"
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/engine/state"
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
	isNewArea := true
	for {
		if isNewArea {
			io.ActiveInputOutputHandler.NewLine(s.state.CurrentArea.Look)
			isNewArea = false
		}
		// TODO: move the action parsing to another file/function
		inputAction, noun := io.ActiveInputOutputHandler.SimpleParse()
		// TODO: set targetedThing to every noun item. Refactor in the process!
		var targetedThing *things.Thing
		if inputAction.Name == "look" {
			targetedThing = action.ExecuteLookCommand(noun, s.state)
		} else if inputAction.Name == "travel" {
			isNewArea = action.ExecuteTravelCommand(noun, &s.state)
		} else if inputAction.Name == "talk" {
			executeTalkCommand(noun, s.state)
		} else if inputAction.Name == "take" {
			executeTakeCommand(noun, s.state)
		} else if inputAction.Name == "equip" {
			action.ExecuteEquipCommand(noun, &s.state)
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
		trigger, ok := targetedThing.Triggers[inputAction.Name]
		if ok {
			err := logic.EvaluateTrigger(s.state, trigger)
			if err != nil {
				io.ActiveInputOutputHandler.NewLinef("Error evaluating trigger: %v", trigger)
			}
		}
	}
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
