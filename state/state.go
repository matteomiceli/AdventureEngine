package state

import (
	"adventureengine/pkg/game"
	"adventureengine/pkg/inventory"
)

var CurrentLocation game.Location = game.Location{
	Name:    "start",
	Message: "You arrive at the start of your journey!",
	Events: []game.Event{
		{},
	},
	GoTo: game.GoTo{},
	Items: []inventory.Item{
		{Name: "Test"},
	},
}
