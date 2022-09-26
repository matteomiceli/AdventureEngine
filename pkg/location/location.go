package location

import (
	"adventureengine/helpers"
	"adventureengine/state"
	"fmt"
)

func Walk(direction string) {
	if !isValidDirection(direction) {
		fmt.Println("Enter a valid direction.")
		return
	}

	if _, ok := state.CurrentLocation().GoTo[direction]; !ok {
		fmt.Println("There doesn't appear to be anything here.")
		return
	}
	state.CurrentLocationId = state.CurrentLocation().GoTo[direction]
	fmt.Println(state.CurrentLocation().Message)
}

func isValidDirection(direction string) bool {
	if direction == "left" || direction == "right" || direction == "back" || direction == "forward" || direction == "l" || direction == "r" || direction == "b" || direction == "f" {
		return true
	}
	return false
}

func Search() {
	if len(state.CurrentLocation().Items) == 0 {
		fmt.Println("No items found nearby.")
		return
	}
	fmt.Println("You find the following nearby:")
	helpers.DrawItems(state.CurrentLocation().Items)
}
