package inventory

import (
	"adventureengine/pkg/color"
	"adventureengine/state"
	"fmt"
	"strings"
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
	isFull := true
	for i, element := range state.Store {
		if element == object {
			fmt.Printf("You're already carrying a %s . \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			return
		}
		if element == "" {
			state.Store[i] = object
			isFull = false
			fmt.Printf("%s added to inventory! \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			return
		}
	}
	if isFull {
		fmt.Printf("Inventory full, please %s an item before trying to pick up something else. \n", color.PaintText(color.Yellow, "DROP"))
	}
}

func Drop(object string) {
	found := true
	for i, element := range state.Store {
		if element == object {
			state.Store[i] = ""
			fmt.Printf("%s dropped. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			return
		}
	}
	if !found {
		fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
	}
}
