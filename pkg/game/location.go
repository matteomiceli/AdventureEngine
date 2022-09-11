package game

import (
	"adventureengine/content"
	"adventureengine/state"
	"fmt"
)

func Walk(direction string) {
	if !isValidDirection(direction) {
		fmt.Println("Enter a valid direction.")
		return
	}

	current := content.Locations[state.CurrentLocationId]

	if _, ok := current.GoTo[direction]; !ok {
		fmt.Println("There doesn't appear to be anything here.")
		return
	}
	state.CurrentLocationId = current.GoTo[direction]
	fmt.Println(content.Locations[state.CurrentLocationId].Message)
}

func isValidDirection(direction string) bool {
	if direction == "left" || direction == "right" || direction == "back" || direction == "forward" {
		return true
	}
	return false
}
