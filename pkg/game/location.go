package game

import "fmt"

type Location struct {
	name    string
	message string
	events  []Event
	goTo    GoTo
}

type GoTo map[string]string

var goTo = GoTo{
	"left":  "solarium",
	"right": "sanctum",
}

func Walk(direction string) {
	if CurrentLocation.goTo[direction] == "" {
		fmt.Println("nothing here.")
		return
	}
	fmt.Println(CurrentLocation.goTo[direction])
}

var CurrentLocation Location = Location{
	name:    "start",
	message: "You arrive at the start of your journey!",
	events: []Event{
		{},
	},
	goTo: goTo,
}
