package content

import "adventureengine/model"

var Locations = map[string]*model.Location{
	"grandEntrance": {
		Id:      "grandEntrance",
		Display: "Grand Entrance",
		Message: "Welcome to the game",
		Events:  []model.Event{},
		GoTo: model.GoTo{
			"left":    "closet",
			"right":   "sittingRoom",
			"forward": "diningRoom",
		},
		Items: []string{},
	},
	"closet": {
		Id:      "closet",
		Display: "Closet",
		Message: "You enter the closet...",
		Events:  []model.Event{},
		GoTo: model.GoTo{
			"back": "grandEntrance",
		},
		Items: []string{},
		Door: model.Door{
			Locked: true,
			Key:    "key",
		},
	},
	"sittingRoom": {
		Id:      "sittingRoom",
		Display: "Sitting Room",
		Message: "You enter the sitting room...",
		Events:  []model.Event{},
		GoTo: model.GoTo{
			"back": "grandEntrance",
		},
		Items: []string{},
	},
	"diningRoom": {
		Id:      "diningRoom",
		Display: "Dining Room",
		Message: "You enter the dining room...",
		Events:  []model.Event{},
		GoTo: model.GoTo{
			"back": "grandEntrance",
		},
		Items: []string{
			"key",
		},
	},
}
