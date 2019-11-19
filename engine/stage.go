package engine

import (
	"gostories/engine/action"
	"gostories/engine/io"
	"gostories/engine/logic"
	"gostories/engine/state"
	"gostories/engine/store"
	"gostories/things"
	"gostories/things/area"
)

// Stage holds the main game state and the game control loop.
type Stage struct {
	state *state.State
}

// Start initialise a Stage, it is passed an Area, which is used to initialise the game state.
func (s Stage) Start(area area.Area) {
	s.state = &state.State{
		CurrentArea:   &area,
		Inventory:     store.NewInventory(),
		EquippedItems: store.NewEquippedItems(),
	}
	s.loopUntilExit()
}

func (s Stage) loopUntilExit() {
	isNewArea := true
	for {
		if isNewArea {
			io.Handler.NewLine(s.state.CurrentArea.Look)
			isNewArea = false
		}
		inputAction, targets := io.Handler.SimpleParse()
		var targetedThing *things.Thing
		if inputAction.Name == "look" {
			targetedThing = action.ExecuteLookCommand(targets[0], s.state)
		} else if inputAction.Name == "travel" {
			isNewArea = action.ExecuteTravelCommand(targets[0], s.state)
		} else if inputAction.Name == "talk" {
			action.ExecuteTalkCommand(targets[0], s.state)
		} else if inputAction.Name == "take" {
			action.ExecuteTakeCommand(targets[0], s.state)
		} else if inputAction.Name == "equip" {
			action.ExecuteEquipCommand(targets[0], s.state)
		} else if inputAction.Name == "place" {
			if len(targets) == 1 {
				// TODO: can actually change the trigger here to "drop" if we want to have a different trigger for put and drop.
				action.ExecutePlaceCommand(targets[0], nil, s.state)
			} else if len(targets) == 2 {
				targetedThing = action.ExecutePlaceCommand(targets[0], &targets[1], s.state)
			}
		} else if inputAction.Name == "inventory" {
			io.Handler.NewLine("You take stock of your store.")
			s.state.Inventory.PrintContents()
			io.Handler.NewLine("You have the following equipped:")
			s.state.EquippedItems.PrintContents()
		} else if inputAction.Name == "help" {
			action.ExecuteHelpCommand(s.state)
			continue
		} else if inputAction.Name == "quit" {
			break
		} else {
			io.Handler.NewLine("Unknown action")
		}
		if targetedThing == nil {
			continue
		}
		trigger, ok := targetedThing.Triggers[inputAction.Name]
		if ok {
			err := logic.EvaluateTrigger(s.state, trigger.Action)
			if err != nil {
				io.Handler.NewLinef("Error evaluating trigger: %v", trigger)
			}
		}
	}
}
