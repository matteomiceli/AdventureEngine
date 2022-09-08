package inventory

import (
	"adventureengine/pkg/color"
	"adventureengine/state"
	"fmt"
)

func Draw() {
	for _, element := range state.Store {
		if element == "" {
			break
		}
		fmt.Println(element)
	}
}

func Add(object string) {
	if len(state.Store) >= 3 {
		fmt.Printf("Inventory full, please %s an item before trying to pick up something else. \n", color.PaintText(color.Yellow, "DROP"))
	}
	// if object does not exist in list of items in room, return

	for _, element := range state.Store {
		if element != "" {

		}

		fmt.Println(element)
	}
}
