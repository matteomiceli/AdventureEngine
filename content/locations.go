package content

import "adventureengine/models"

var Locations = map[string]models.Location{
	"test": {
		Id:      "test",
		Message: "",
		Events:  []models.Event{},
		GoTo:    models.GoTo{},
		Items:   []models.Item{},
	},
}
