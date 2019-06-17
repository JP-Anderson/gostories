package engine

import (
	"strings"

	"gostories/engine/io"
	"gostories/things"
)

type Stage struct {
	context Context
}

func (s Stage) Start(area things.Area) {
	s.context = Context{
		CurrentArea:   area,
		Inventory:     NewInventory(),
		EquippedItems: NewEquippedItems(),
	}
	s.LoopUntilExit()
}

func (s Stage) LoopUntilExit() {
	newArea := true
	for {
		if newArea {
			io.NewLine(s.context.CurrentArea.Look)
			newArea = false
		}
		// TODO: move the action parsing to another file/function
		action, noun := io.SimpleParse()
		if action.Name == "look" {
			// TODO also Look at NPCs
			if noun == "" {
				io.NewLine(s.context.CurrentArea.Look)
			} else {
				found := false
				for _, item := range s.context.CurrentArea.Items {
					if strings.ToLower(item.GetName()) == strings.ToLower(noun) {
						found = true
						io.NewLine(item.GetLookText())
					}
				}
				if !found {
					io.NewLinef("Couldn't find a %v to look at!", noun)
				}
			}
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
			if !found {
				io.NewLinef("Couldn't find a %v to pick up.", noun)
			}
		} else if action.Name == "equip" {
			for _, item := range s.context.Inventory.items {
				if ok := item.(things.Equippable); ok != nil {
					s.context.Equip(item)
				}
			}
		} else if action.Name == "inventory" {
			if s.context.Inventory.Size() > 0 {
				io.NewLine("You take stock of your items.")
				for _, item := range s.context.Inventory.items {
					io.NewLine(item.GetName())
				}
			} else {
				io.NewLinef("You aren't carrying anything.")
			}
			if s.context.EquippedItems.Size() > 0 {
				io.NewLine("You have the following equipped:")
				for _, item := range s.context.EquippedItems.items {
					io.NewLine(item.GetName())
				}
			}
		} else if action.Name == "exit" {
			break
		} else {
			io.NewLine("Unknown action")
		}
	}
}
