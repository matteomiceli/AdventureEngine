package game

import (
	"adventureengine/pkg/inventory"
	"adventureengine/state"
	"fmt"
)

type Location struct {
	Name    string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []inventory.Item
}

type GoTo map[string]string

var goTo = GoTo{
	"left":  "solarium",
	"right": "sanctum",
}

func Walk(direction string) {
	if state.CurrentLocation.GoTo[direction] == "" {
		fmt.Println("nothing here.")
		return
	}
	fmt.Println(state.CurrentLocation.GoTo[direction])
}
