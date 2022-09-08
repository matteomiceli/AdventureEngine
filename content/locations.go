package content

import (
	"adventureengine/pkg/game"
	"adventureengine/pkg/inventory"
)

var Locations = map[string]game.Location{
	"test": {
		Id:      "test",
		Message: "",
		Events:  []game.Event{},
		GoTo:    game.GoTo{},
		Items:   []inventory.Item{},
	},
}
