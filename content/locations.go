package content

import "adventureengine/models"

var Locations = map[string]models.Location{
	"grandEntrance": {
		Id:      "grandEntrance",
		Display: "Grand Entrance",
		Message: "Welcome to the game",
		Events:  []models.Event{},
		GoTo: models.GoTo{
			"left":    "closet",
			"right":   "sittingRoom",
			"forward": "diningRoom",
		},
		Items: []models.Item{},
	},
	"closet": {
		Id:      "closet",
		Display: "Closet",
		Message: "You enter the closet...",
		Events:  []models.Event{},
		GoTo: models.GoTo{
			"back": "grandEntrance",
		},
		Items: []models.Item{},
	},
	"sittingRoom": {
		Id:      "sittingRoom",
		Display: "Sitting Room",
		Message: "You enter the sitting room...",
		Events:  []models.Event{},
		GoTo: models.GoTo{
			"back": "grandEntrance",
		},
		Items: []models.Item{},
	},
	"diningRoom": {
		Id:      "diningRoom",
		Display: "Dining Room",
		Message: "You enter the dining room...",
		Events:  []models.Event{},
		GoTo: models.GoTo{
			"back": "grandEntrance",
		},
		Items: []models.Item{},
	},
}
