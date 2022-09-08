package game

import (
	"adventureengine/pkg/inventory"
	"fmt"
)

type Location struct {
	Id      string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []inventory.Item
}

type GoTo map[string]string

var CurrentLocation Location = Location{
	Id:      "start",
	Message: "You arrive at the start of your journey!",
	Events: []Event{
		{},
	},
	GoTo: goTo,
	Items: []inventory.Item{
		{Name: "Test"},
	},
}

var goTo = GoTo{
	"left":  "solarium",
	"right": "sanctum",
}

func Walk(direction string) {
	if CurrentLocation.GoTo[direction] == "" {
		fmt.Println("nothing here.")
		return
	}
	fmt.Println(CurrentLocation.GoTo[direction])
}
