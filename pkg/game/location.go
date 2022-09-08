package game

import (
	"adventureengine/state"
	"fmt"
)

func Walk(direction string) {
	// compare against list of locations in content
	if state.CurrentLocation == "" {
		fmt.Println("nothing here.")
		return
	}
	fmt.Println(state.CurrentLocation)
}
