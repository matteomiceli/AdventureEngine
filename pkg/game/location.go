package game

import (
	"adventureengine/content"
	"adventureengine/model"
	"adventureengine/state"
	"fmt"
)

func CurrentLocation() model.Location {
	return content.Locations[state.CurrentLocationId]
}

func Walk(direction string) {
	if !isValidDirection(direction) {
		fmt.Println("Enter a valid direction.")
		return
	}

	current := CurrentLocation()

	if _, ok := current.GoTo[direction]; !ok {
		fmt.Println("There doesn't appear to be anything here.")
		return
	}
	state.CurrentLocationId = current.GoTo[direction]
	fmt.Println(CurrentLocation().Message)
}

func isValidDirection(direction string) bool {
	if direction == "left" || direction == "right" || direction == "back" || direction == "forward" {
		return true
	}
	return false
}
